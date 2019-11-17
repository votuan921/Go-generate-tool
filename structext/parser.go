package structext

import (
	"bufio"
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// Parser is interface of go file parser
type Parser interface {
	// Parse parses a given file to structs
	Parse(structPath string) ([]*ParsedStruct, error)
}

// ParsedStruct represents a struct of a parsed struct
type ParsedStruct struct {
	StructName string
	IDType     string
	Fields     map[string]string //key: Field name, value: Field type
}

// NewParser returns a implemented instance of Parser interface
func NewParser() Parser {
	return &innerParser{}
}

type innerParser struct{}

func (i *innerParser) Parse(structPath string) (parsedStructs []*ParsedStruct, err error) {
	fStrs, err := readFileToString(structPath)
	if err != nil {
		return nil, err
	}

	fSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fSet, structPath, nil, 0)
	if err != nil {
		return nil, err
	}

	parsedStructs = make([]*ParsedStruct, 0)
	var ps *ParsedStruct
	for _, decl := range parsedFile.Decls {
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

			ps = &ParsedStruct{
				StructName: typeSpec.Name.Name,
				Fields:     make(map[string]string),
			}
			for _, field := range structType.Fields.List {
				fName := field.Names[0].Name
				fType := getStringFromNodePosition(fSet, fStrs, field.Type.Pos(), field.Type.End())
				if fName == "Id" {
					ps.IDType = fType
				}
				ps.Fields[fName] = fType
			}
			parsedStructs = append(parsedStructs, ps)
		}
	}

	return
}

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

func getStringFromNodePosition(fSet *token.FileSet, fStrs []string, pos token.Pos, end token.Pos) string {
	linePos, colPos := fSet.Position(pos).Line, fSet.Position(pos).Column
	lineEnd, colEnd := fSet.Position(end).Line, fSet.Position(end).Column
	var buf bytes.Buffer

	if linePos == lineEnd {
		buf.WriteString(fStrs[linePos-1][colPos-1 : colEnd-1])
	} else {
		buf.WriteString(fStrs[linePos-1][colPos-1:])
		for i := linePos + 1; i < lineEnd; i++ {
			buf.WriteString(fStrs[i])
		}
		buf.WriteString(fStrs[lineEnd-1][:colEnd-1])
	}

	return buf.String()
}
