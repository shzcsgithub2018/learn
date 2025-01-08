package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"

	"golang.org/x/tools/go/ast/astutil"
)

var src string = `package main

import(
	"errors"
)

func main(){
	errors.New("xxxxx")
	errors.New("-----")
	errors.New(Foo("asdfasdf"))
}

func Ts(){
	errors.New("zzz")
}

func Foo(s string){}
`

func Rewrite() {
	// 设置可写入文件
	fset := token.NewFileSet()
	// 解析文件
	file, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	//ast.Print(fset, file)
	astutil.RewriteImport(fset, file, "errors", "syserrors")

	// 导入file，最终返回file。 .Apply方法类似ast.Inspect
	astutil.Apply(file, nil, func(c *astutil.Cursor) bool {

		n := c.Node()
		switch t := n.(type) {
		case *ast.CallExpr:
			if sl, ok := t.Fun.(*ast.SelectorExpr); ok {
				// 检查语句是否是errors.New
				if sl.Sel.Name != "New" {
					break
				}
				// 断言是否有子属性 sl.X.(*ast.Ident)
				if cl, ok := sl.X.(*ast.Ident); ok {
					// 判断是否是errors，否则跳过
					if cl.Name != "errors" {
						break
					}
					// 替换成syserror
					cl.Name = "syserror"
					// 拼接两个数组
					t.Args = append([]ast.Expr{&ast.BasicLit{Kind: token.IDENT, Value: "traceID"}}, t.Args...)
				}
			}
		}
		return true
	})

	if err := format.Node(os.Stdout, token.NewFileSet(), file); err != nil {
		log.Fatalln("Error:", err)
	}
}