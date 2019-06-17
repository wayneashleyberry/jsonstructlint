// Package jsoncheck inspects field types to ensure naming conventions
package jsoncheck

import (
	"go/ast"
	"reflect"
	"strings"

	"github.com/wayneashleyberry/jsonstructlint/v4/pkg/stringutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer will create a new analyzer
func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "jsoncheck",
		Doc:      "reports json convention violations",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	// get the inspector. This will not panic because inspect.Analyzer is part
	// of `Requires`. go/analysis will populate the `pass.ResultOf` map with
	// the prerequisite analyzers.
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// the inspector has a `filter` feature that enables type-based filtering
	// The anonymous function will be only called for the ast nodes whose type
	// matches an element in the filter
	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
	}

	// this is basically the same as ast.Inspect(), only we don't return a
	// boolean anymore as it'll visit all the nodes based on the filter.
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		f := n.(*ast.Field)
		if f.Tag == nil {
			return
		}

		val, ok := lookupJSON(f)
		if !ok {
			return
		}

		if f.Comment != nil && stringutil.ContainsIgnoreString(f.Comment.Text()) {
			return
		}

		if !stringutil.IsTrimmed(val) {
			pass.Reportf(f.Pos(), "`%s` contains whitespace", val)
		}

		if !stringutil.IsCamelCase(val) {
			pass.Reportf(f.Pos(), "`%s` is not camelcase", val)
		}
	})

	return nil, nil
}

func lookupJSON(f *ast.Field) (string, bool) {
	tag := reflect.StructTag(f.Tag.Value[1 : len(f.Tag.Value)-1])
	val, ok := tag.Lookup("json")
	if !ok {
		return "", false
	}

	if strings.Contains(val, ",") {
		parts := strings.Split(val, ",")
		val = parts[0]
	}

	return val, true
}
