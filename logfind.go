package logfind

import (
	"go/ast"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "logfind is ...."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "logfind",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	comments := getCommentMap(pass)
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			var nocheck bool
			pos := pass.Fset.Position(n.Pos())
			c, ok := comments[pos.Filename+"_"+strconv.Itoa(pos.Line)]
			if ok {
				if strings.Contains(c, "nocheck:thislog") {
					nocheck = true
				}
			}
			if n.Name == "log" && !nocheck {
				pass.Reportf(n.Pos(), "here log")
			}
		}
	})

	return nil, nil
}

func getCommentMap(pass *analysis.Pass) map[string]string {
	var mp = make(map[string]string)

	for _, file := range pass.Files {
		for _, cg := range file.Comments {
			for _, c := range cg.List {
				pos := pass.Fset.Position(c.Pos())
				mp[pos.Filename+"_"+strconv.Itoa(pos.Line)] = c.Text
			}
		}
	}

	return mp
}
