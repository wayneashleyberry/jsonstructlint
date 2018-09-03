package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
	"unicode"
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
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./testdata/foo.go", nil, 0)

	if err != nil {
		panic(err)
	}

	typeDecl := f.Decls[0].(*ast.GenDecl)
	structDecl := typeDecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
	fields := structDecl.Fields.List

	messages := []string{}

	for _, field := range fields {
		if field.Tag == nil {
			continue
		}
		tag := reflect.StructTag(strings.Replace(field.Tag.Value, "`", "", -1))
		val, ok := tag.Lookup("json")
		if ok {
			if strings.Contains(val, "_") {
				messages = append(
					messages,
					fmt.Sprintf("`%s` is not camelcase", val),
				)
			}
			if trim(val) != val {
				messages = append(
					messages,
					fmt.Sprintf("`%s` contains whitespaces", val),
				)
			}
			if strings.ToLower(string(val[0])) != string(val[0]) {
				messages = append(
					messages,
					fmt.Sprintf("`%s` first character is not lowercase", val),
				)
			}
		}
	}

	if len(messages) != 0 {
		for _, msg := range messages {
			fmt.Println(msg)
		}
		os.Exit(1)
	}
}
