package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/votuan921/struct-extend-generator/structext"
)

func usage() {
	helpStr := `
	Extended methods generator for Go struct.
	Usage: $struct-extend-generator [options] /absolute/path/to/structs_file.go /absolute/path/to/template.tpl
	Options:
			-e: output file extension. Default: ".extend.go"
			-o: absolute path to output file directory. Default is struct file dir.
`
	fmt.Fprint(os.Stderr, helpStr)
}

func main() {
	outputExt := flag.String("e", ".extend.go", "output file extension. Default: \".extend.go\"")
	outputDir := flag.String("o", "", "absolute path to output file directory. Default is struct file dir.")
	flag.Parse()

	structAndTemplateFiles := flag.Args()
	if len(structAndTemplateFiles) != 2 {
		usage()
		os.Exit(1)
	}
	structPath := structAndTemplateFiles[0]
	tplPath := structAndTemplateFiles[1]

	if *outputDir == "" { // no "-o" option given, default to struct dir
		structDir := structext.FileDir(structPath)
		outputDir = &structDir
	}

	parser := structext.NewParser()
	generator, err := structext.NewGenerator(structPath, tplPath, *outputDir, *outputExt, parser)
	if err != nil {
		fmt.Fprint(os.Stderr, "cannot init generator, err: ", err)
		usage()
		os.Exit(1)
	}
	outFilePath, err := generator.Generate()
	if err != nil {
		fmt.Fprint(os.Stderr, "cannot generate file, err: ", err)
		usage()
		os.Exit(1)
	}

	fmt.Println("Generated file: ", outFilePath)
}
