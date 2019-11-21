package structext

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

// NewGenerator inits and returns *Generator with validated params
func NewGenerator(structPath, tplPath, outDir, outFileExt string, parser Parser) (*Generator, error) {
	if !FileExisting(structPath) {
		return nil, errors.Errorf("struct path not existing: %s", structPath)
	}

	if !FileExisting(tplPath) {
		return nil, errors.Errorf("template path not existing: %s", tplPath)
	}

	if parser == nil {
		return nil, errors.New("no parser given")
	}

	fStructName := FileNameWithoutExt(FileName(structPath))
	if outDir == "" { // no out dir given, default to same struct dir
		outDir = FileDir(structPath)
	} else {
		if Mkdirp(outDir) != nil {
			return nil, errors.Errorf("cannot create dir: %s", outDir)
		}
	}

	if outFileExt == "" { // no out file ext given, default to ".extend.go"
		outFileExt = ".extend.go"
	}

	return &Generator{
		structPath:           structPath,
		structNameWithoutExt: fStructName,
		tplPath:              tplPath,
		outDir:               outDir,
		outFileExt:           outFileExt,
		parser:               parser,
	}, nil
}

// Generator contains all info to generate file
type Generator struct {
	structPath           string
	structNameWithoutExt string
	tplPath              string
	outFileExt           string
	outDir               string
	parser               Parser
}

// Generate generates the extended file with given struct and template file path
func (g *Generator) Generate() (outFilePath string, err error) {
	parsedStructs, err := g.parser.Parse(g.structPath)
	if err != nil {
		return "", errors.Wrapf(err, "cannot parse struct file: %s", g.structPath)
	}

	var buf bytes.Buffer
	var outFileName string
	generatedTmpl := template.Must(template.New(filepath.Base(g.tplPath)).ParseFiles(g.tplPath))
	err = generatedTmpl.Execute(&buf, parsedStructs)
	if err != nil {
		return "", errors.Wrapf(err,
			"cannot generate code from struct: %s with template: %s",
			g.structPath,
			g.tplPath)
	}

	outFileName = fmt.Sprintf("%s.%s%s",
		g.structNameWithoutExt,
		FileNameWithoutExt(FileName(g.tplPath)),
		g.outFileExt)
	outFilePath = DirJoin(g.outDir, outFileName)
	err = g.writeToFile(buf, outFilePath)
	if err != nil {
		return "", errors.Wrapf(err, "cannot write to file: %s", outFilePath)
	}

	return
}

func (g *Generator) writeToFile(buf bytes.Buffer, outputPath string) error {
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "cannot format source")
	}

	if err := ioutil.WriteFile(outputPath, src, 0644); err != nil {
		return errors.Wrapf(err, "cannot write to file: %s", outputPath)
	}
	return nil
}
