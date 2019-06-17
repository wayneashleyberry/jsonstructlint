// Package structcheck inspects struct types to ensure conventions across
// the entire struct, not just individual fields.
package structcheck

import (
	"go/ast"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer will create a new analyzer
func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "structcheck",
		Doc:      "reports json struct convention violations",
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
		(*ast.StructType)(nil),
	}

	// this is basically the same as ast.Inspect(), only we don't return a
	// boolean anymore as it'll visit all the nodes based on the filter.
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		s := n.(*ast.StructType)
		if s.Fields == nil {
			return
		}

		if s.Fields.NumFields() == 0 {
			return
		}

		hasJSON := 0
		hasNoJSON := 0

		missing := []*ast.Field{}

		for _, f := range s.Fields.List {
			if f.Tag == nil {
				hasNoJSON++
				missing = append(missing, f)
				continue
			}

			tag := reflect.StructTag(f.Tag.Value[1 : len(f.Tag.Value)-1])
			_, ok := tag.Lookup("json")
			if ok {
				hasJSON++
			}
			if !ok {
				hasNoJSON++
				missing = append(missing, f)
			}
		}

		if hasJSON > 0 && hasNoJSON != s.Fields.NumFields() {
			for _, ff := range missing {
				pass.Reportf(ff.Pos(), "`%s` is missing a json tag", ff.Names[0].Name)
			}
		}
	})

	return nil, nil
}
