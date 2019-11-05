package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	helpStr := `
	Extended methods generator for Go struct.
	Usage: $struct-extend-generator [options] /absolute/path/to/structs_file.go /absolute/path/to/template.tpl [/absolute/path/to/template2.tpl ...]
	Options:
			-e: output file extension. Default: ".extend.go"
			-o: absolute path to output file directory. Default is struct file dir.
`
	fmt.Fprint(os.Stderr, helpStr)
}

func fileExisting(filePath string) bool {

}

func main() {
	outputExt := flag.String("e", ".extend.go", "output file extension. Default: \".extend.go\"")
	outputDir := flag.String("o", "", "absolute path to output file directory. Default is struct file dir.")
	flag.Parse()

	structAndTemplateFiles := flag.Args()
	if len(structAndTemplateFiles) < 2 {
		usage()
		os.Exit(1)
	}
	structPath := structAndTemplateFiles[0]
	templatePaths := structAndTemplateFiles[1:]
	if !fileExisting(structPath) {
		fmt.Fprintf(os.Stderr, "invalid struct file: %s", structPath)
		os.Exit(1)
	}
	for _, t := range templatePaths {
		if !fileExisting(t) {
			fmt.Fprintf(os.Stderr, "invalid template file: %s", t)
			os.Exit(1)
		}
	}

	if *outputDir == "" { // no "-o" option given, default to struct dir

	}
}
