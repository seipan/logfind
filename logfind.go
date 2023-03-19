package logfind

import (
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
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
	objmp := getImportObj(pass)
	types := pass.TypesInfo

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CallExpr:
			value, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}
			var nocheck bool
			pos := pass.Fset.Position(n.Pos())
			c, ok := comments[pos.Filename+"_"+strconv.Itoa(pos.Line)]
			if ok {
				if strings.Contains(c, "nocheck:thislog") {
					nocheck = true
				}
			}
			_, ok = objmp[types.ObjectOf(value.Sel)]

			if ok && !nocheck {
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

func getImportObj(pass *analysis.Pass) map[types.Object]bool {
	var mp = make(map[types.Object]bool)
	pkgs := pass.Pkg.Imports()
	obj := analysisutil.LookupFromImports(pkgs, "log", "Print")
	mp[obj] = true
	obj = analysisutil.LookupFromImports(pkgs, "log", "Println")
	mp[obj] = true
	obj = analysisutil.LookupFromImports(pkgs, "log", "Printf")
	mp[obj] = true
	obj = analysisutil.LookupFromImports(pkgs, "log", "Fatal")
	mp[obj] = true
	obj = analysisutil.LookupFromImports(pkgs, "log", "Fatalln")
	mp[obj] = true
	obj = analysisutil.LookupFromImports(pkgs, "log", "Fatalf")
	mp[obj] = true
	obj = analysisutil.LookupFromImports(pkgs, "log", "Panicf")
	mp[obj] = true
	obj = analysisutil.LookupFromImports(pkgs, "log", "Panic")
	mp[obj] = true
	obj = analysisutil.LookupFromImports(pkgs, "log", "Panicln")
	mp[obj] = true

	return mp
}
