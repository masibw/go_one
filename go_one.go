package go_one

import (
	_ "database/sql"
	"fmt"
	"github.com/gostaticanalysis/analysisutil"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"log"
)

const doc = "go_one is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "go_one",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	forFilter := []ast.Node{
		(*ast.ForStmt)(nil),
	}

	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	inspect.Preorder(forFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ForStmt:
			ast.Inspect(n, func(n ast.Node) bool {
				switch node := n.(type) {
				case *ast.Ident:
					if tv, ok := pass.TypesInfo.Types[node]; ok {
						obj := analysisutil.TypeOf(pass, "database/sql", "*DB")
						if node.Name == "cnn"{
							ast.Print(nil, node)
						}
						if types.Identical(tv.Type, obj) {
							pass.Reportf(node.Pos(), "this query might be causes bad performance")
							log.Printf("%s is detected %d ", node.Name, node.NamePos)
						}
					}
				case *ast.CallExpr:
					var funcName *ast.Ident
					//ast.Print(nil,node.Fun)
					switch funcExpr := node.Fun.(type){
					case *ast.Ident:
						ast.Print(nil,funcExpr)
						funcName = funcExpr
						fmt.Println(funcName)
						obj := info.ObjectOf(funcName)
						ast.Print(nil,obj)
						if obj == nil {
							return false
						}
					}


				}
				return true

			})


		}

	})

	return nil, nil
}