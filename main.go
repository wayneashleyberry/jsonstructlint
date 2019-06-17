package main

import (
	"github.com/wayneashleyberry/jsonstructlint/v4/pkg/jsoncheck"
	"github.com/wayneashleyberry/jsonstructlint/v4/pkg/structcheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		jsoncheck.Analyzer(),
		structcheck.Analyzer(),
	)
}
