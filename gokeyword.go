package gokeyword

import (
	"go/ast"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	goKeywordName        = "gokeyword"
	goKeywordErrorMsg    = "detected use of go keyword: %s"
	goKeywordDescription = "detects presence of go keyword"
	defaultDetails       = "no details provided"
	FlagDetails          = "details"
)

func New() *analysis.Analyzer {
	a := &analysis.Analyzer{
		Name:     goKeywordName,
		Doc:      goKeywordDescription,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
	a.Flags.String(FlagDetails, defaultDetails, "Details to include in the error message")
	return a
}

func run(pass *analysis.Pass) (interface{}, error) {
	details := pass.Analyzer.Flags.Lookup(FlagDetails).Value.String()
	i, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errors.New("analyzer is not type *inspector.Inspector")
	}

	i.Preorder([]ast.Node{(*ast.GoStmt)(nil)}, func(node ast.Node) {
		if _, ok := node.(*ast.GoStmt); ok {
			pass.Reportf(node.Pos(), goKeywordErrorMsg, details)
		}
	})
	return nil, nil
}

func gonogo() {
	go func() {
		// do something
	}()
}
