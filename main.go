package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"
	"unicode"

	"golang.org/x/tools/go/packages"
)

func trim(in string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, in)
}

func main() {
	flag.Parse()
	pkgs, err := packages.Load(nil, flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}

	messages := []string{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range pkgs {
		for _, filename := range pkg.GoFiles {
			filename = strings.Replace(filename, dir, ".", 1)
			messages = append(messages, lint(filename)...)
		}
	}

	if len(messages) != 0 {
		for _, message := range messages {
			fmt.Println(message)
		}
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
