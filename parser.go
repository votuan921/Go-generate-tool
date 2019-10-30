package main

import "fmt"
import "go/token"
import "go/parser"
import "go/ast"

const PARSER_FILE_PATH = "db.go"

func  main()  {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, PARSER_FILE_PATH, nil, 0)

	if err != nil {
		fmt.Println("Faild Parsing", err)
		return
	}

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			fmt.Println("Struct: ", typeSpec.Name)
			fields := structType.Fields.List
			for _, field := range fields {
				fmt.Printf("--Field name: %s\n", field.Names)
			}
		}
	}
}