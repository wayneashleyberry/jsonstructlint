package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"reflect"
	"strings"
	"unicode"

	"golang.org/x/tools/go/packages"
)

func main() {
	flag.Parse()

	importPaths := flag.Args()
	if len(importPaths) == 0 {
		importPaths = []string{"."}
	}

	var flags []string

	cfg := &packages.Config{
		Mode:       packages.LoadSyntax,
		BuildFlags: flags,
		Error: func(error) {
			// don't print type check errors
		},
	}

	pkgs, err := packages.Load(cfg, importPaths...)
	if err != nil {
		log.Fatalf("could not load packages: %s", err)
	}

	var lines []string

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range pkgs {
		for _, obj := range pkg.TypesInfo.Defs {
			if obj == nil {
				continue
			}

			if _, ok := obj.(*types.TypeName); !ok {
				continue
			}

			typ, ok := obj.Type().(*types.Named)
			if !ok {
				continue
			}

			strukt, ok := typ.Underlying().(*types.Struct)
			if !ok {
				continue
			}

			for i := 0; i < strukt.NumFields(); i++ {
				field := strukt.Field(i)
				pos := pkg.Fset.Position(field.Pos())
				relname := strings.Replace(pos.Filename, dir, ".", 1)
				tag := reflect.StructTag(strukt.Tag(i))
				val, ok := tag.Lookup("json")

				if ok {
					if strings.Contains(val, ",") {
						parts := strings.Split(val, ",")
						val = parts[0]
					}

					if trim(val) != val {
						lines = append(
							lines,
							fmt.Sprintf(`%s:%d: "%s" contains whitespace`, relname, pos.Line, val),
						)
					} else if !isCamelCase(val) {
						lines = append(
							lines,
							fmt.Sprintf(`%s:%d: "%s" is not camelcase`, relname, pos.Line, val),
						)
					}
				}
			}
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	if len(lines) > 0 {
		os.Exit(1)
	}
}

func isCamelCase(val string) bool {
	if strings.Contains(val, "_") {
		return false
	}

	if strings.ToLower(string(val[0])) != string(val[0]) {
		return false
	}

	return true
}

func lint(filename string) []string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	messages := []string{}

	for _, decl := range f.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok {
			for _, s := range funcDecl.Body.List {
				decl, ok := s.(*ast.DeclStmt)
				if ok {
					typeDecl, ok := decl.Decl.(*ast.GenDecl)
					if ok {
						for _, s := range typeDecl.Specs {
							m := lintSpec(fset, s)
							messages = append(messages, m...)
						}
					}
				}
			}
		}

		typeDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range typeDecl.Specs {
			m := lintSpec(fset, spec)
			messages = append(messages, m...)
		}
	}

	return messages
}

func shouldIgnore(field *ast.Field) bool {
	if field.Comment == nil {
		return false
	}

	if len(field.Comment.List) == 0 {
		return false
	}

	for _, comment := range field.Comment.List {
		if containsIgnoreString(comment.Text) {
			return true
		}
	}

	return false
}

func containsIgnoreString(in string) bool {
	if !strings.Contains(in, "nolint:") {
		return false
	}

	parts := strings.Split(in, ":")

	for _, part := range parts[1:] {
		if strings.Contains(part, "jsonstructlint") {
			return true
		}
	}

	return false
}

func trim(in string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, in)
}

func lintSpec(fset *token.FileSet, spec ast.Spec) []string {
	messages := []string{}

	typespec, ok := spec.(*ast.TypeSpec)
	if !ok {
		return messages
	}

	structDecl, ok := typespec.Type.(*ast.StructType)
	if !ok {
		return messages
	}

	fields := structDecl.Fields.List

	for _, field := range fields {
		if field.Tag == nil {
			continue
		}
		if shouldIgnore(field) {
			continue
		}
		pos := fset.Position(field.Tag.ValuePos)
		tag := reflect.StructTag(strings.Replace(field.Tag.Value, "`", "", -1))
		val, ok := tag.Lookup("json")

		if ok {
			if strings.Contains(val, ",") {
				parts := strings.Split(val, ",")
				val = parts[0]
			}

			if trim(val) != val {
				messages = append(
					messages,
					fmt.Sprintf(`%s:%d: "%s" contains whitespace`, pos.Filename, pos.Line, val),
				)
			} else if !isCamelCase(val) {
				messages = append(
					messages,
					fmt.Sprintf(`%s:%d: "%s" is not camelcase`, pos.Filename, pos.Line, val),
				)
			}
		}
	}

	return messages
}
