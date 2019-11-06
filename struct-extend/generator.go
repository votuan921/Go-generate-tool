package struct_extend

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

func NewGenerator(structPath string, templatePaths []string, outDir, outFileExt string, parser Parser) (*Generator, error) {
	if parser == nil {
		return nil, errors.New("no parser given")
	}
	if len(templatePaths) == 0 {
		return nil, errors.New("no template given")
	}

	fStructName := FileNameWithoutExt(FileName(structPath))
	if outDir == "" { // no out dir given, default to same struct dir
		outDir = FileDir(structPath)
	}
	if outFileExt == "" { // no out file ext given, default to ".extend.go"
		outFileExt = ".extend.go"
	}
	return &Generator{
		structPath:           structPath,
		structNameWithoutExt: fStructName,
		templatePaths:        templatePaths,
		outDir:               outDir,
		outFileExt:           outFileExt,
		parser:               parser,
	}, nil
}

type Generator struct {
	structPath           string
	structNameWithoutExt string
	templatePaths        []string
	outFileExt           string
	outDir               string
	parser               Parser
}

func (g *Generator) Generate() (outFilePaths []string, err error) {
	parsedStructs, err := g.parser.Parse(g.structPath)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse struct file: %s", g.structPath)
	}

	var bufs bytes.Buffer
	var outFile, outPath string
	outFilePaths = make([]string, 0, len(g.templatePaths))
	for _, tpl := range g.templatePaths {
		buf := bytes.Buffer{}
		generatedTmpl := template.Must(template.New(filepath.Base(tpl)).ParseFiles(tpl))
		err := generatedTmpl.Execute(&buf, parsedStructs)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot generate code from template: %s", tpl)
		}
		outFile = fmt.Sprintf("%s.%s.%s", g.structNameWithoutExt, FileNameWithoutExt(FileName(tpl)), g.outFileExt)
		outPath = DirJoin(g.outDir, outFile)
		g.writeToFile(buf, outPath)
		outFilePaths = append(outFilePaths, outPath)
		bufs.Reset()
	}

	return
}

func (g *Generator) writeToFile(buf bytes.Buffer, outputPath string) {
	src, err := format.Source(buf.Bytes())

	if err != nil {
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		src = buf.Bytes()
	}
	os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
	err = ioutil.WriteFile(outputPath, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
