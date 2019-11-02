package main

import "fmt"
import "go/token"
import "go/parser"
import "go/ast"
import "bytes"
import "bufio"
import "os"

type ParsedStruct struct {
	StructName string
	IDType string
	Fields map[string]string
}

const PARSER_FILE_PATH = "db.go"

var parsedStructs []*ParsedStruct

func readFileToString(path string) ([]string, error) {
	file, err := os.Open(path)
	
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func getStringFromNodePosition(fset *token.FileSet, fStr []string, pos token.Pos, end token.Pos) string {
	linePos, colPos := fset.Position(pos).Line, fset.Position(pos).Column
	lineEnd, colEnd := fset.Position(end).Line, fset.Position(end).Column
	buf := bytes.Buffer{}

	if linePos == lineEnd {
		buf.WriteString(fStr[linePos-1][colPos-1:colEnd-1])
	} else {
		buf.WriteString(fStr[linePos-1][colPos-1:])
		for i:= linePos+1; i < lineEnd; i++ {
			buf.WriteString(fStr[i])
		}
		buf.WriteString(fStr[lineEnd-1][:colEnd-1])
	}
	
	return buf.String()
}

func  main()  {
	src, err := readFileToString(PARSER_FILE_PATH)

	if err != nil {
		fmt.Println("Reading file error ", err)
	}
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
			
			var nStruct *ParsedStruct
			nStruct = new(ParsedStruct)
			nStruct.StructName = typeSpec.Name.Name
			nStruct.Fields = make(map[string]string)

			fields := structType.Fields.List
			for _, field := range fields {
				fname := field.Names[0].Name
				ftype := getStringFromNodePosition(fset, src, field.Type.Pos(), field.Type.End())
				if fname == "Id" {
					nStruct.IDType = ftype
				}
				nStruct.Fields[fname] = ftype
			}
			parsedStructs = append(parsedStructs, nStruct)
		}
	}
}