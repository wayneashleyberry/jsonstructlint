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
	for _, pkg := range pkgs {
		for _, filename := range pkg.GoFiles {
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

	if trim(val) != val {
		return false
	}
	return true
}

func lint(filename string) []string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)

	if err != nil {
		panic(err)
	}

	messages := []string{}

	for _, decl := range f.Decls {
		typeDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range typeDecl.Specs {
			typespec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structDecl, ok := typespec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			fields := structDecl.Fields.List

			for _, field := range fields {
				if field.Tag == nil {
					continue
				}
				tag := reflect.StructTag(strings.Replace(field.Tag.Value, "`", "", -1))
				val, ok := tag.Lookup("json")

				if ok {
					if strings.Contains(val, ",") {
						parts := strings.Split(val, ",")
						val = parts[0]
					}
					if !isCamelCase(val) {
						messages = append(
							messages,
							fmt.Sprintf(`%s: "%s" is not camelcase`, filename, val),
						)
					}
				}
			}

		}
	}

	return messages
}
