package struct_extend

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Generator struct {
	structPath       string
	templatePaths    []string
	generatedDir     string
	generatedFileExt string
	parser           Parser
}

func (g *Generator) Generate() (files []string, err error) {

}

func NewGenerator() {

}

func writeToFile(buf bytes.Buffer, outputPath string) {
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

func write() {
	parsedStructs, err := se.ParseStruct(structPath)
	for idx := 1; idx < len(extra); idx++ {
		tmplPath, err := filepath.Abs(extra[idx])

		if err != nil {
			log.Printf("unable detect template path: %s: %v", tmplPath, err)
			continue
		}
		var bufs bytes.Buffer
		for _, parsedStruct := range parsedStructs {
			buf := bytes.Buffer{}
			generatedTmpl := template.Must(template.New(filepath.Base(tmplPath)).ParseFiles(tmplPath))
			err := generatedTmpl.Execute(&buf, parsedStruct)

			if err != nil {
				log.Fatalf("generating code: %v", err)
			}
			bufs.Write(buf.Bytes())
			if *groupPrefix == "" {
				outputPath := filepath.Join(dir,
					"generated",
					filepath.Base(tmplPath),
					strings.ToLower(parsedStruct.StructName+*outputSuffix+".go"))
				se.writeToFile(buf, outputPath)
			}
		}
		if *groupPrefix != "" {
			outputPath := filepath.Join(dir,
				"generated",
				filepath.Base(tmplPath),
				strings.ToLower(*groupPrefix+*outputSuffix+".go"))
			se.writeToFile(bufs, outputPath)
		}
	}
}
