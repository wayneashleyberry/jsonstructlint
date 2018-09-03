package main

import (
	"flag"
	"fmt"
	"go/ast"
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
							fmt.Sprintf(`%s: "%s" contains whitespace`, pos, val),
						)
					} else if !isCamelCase(val) {
						lines = append(
							lines,
							fmt.Sprintf(`%s: "%s" is not camelcase`, pos, val),
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
