package main

import (
        "bytes"        
        "flag"
        "go/format"    
        "io/ioutil"    
        "log"
        "os"
        "path/filepath"
        "strings"
        "github.com/votuan921/struct-extend-generator/parser"
        "text/template"
)

var (
        groupPrefix  = flag.String("g", "", "Group all struct generated to one file with prefix name, default empty meaning separate into multi file with StructName prefix")
        outputSuffix = flag.String("o", "_extend", "suffix to be added to the output file")
)

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

func main() {
        flag.Parse()
        extra := flag.Args()
        dir, err := filepath.Abs("./")

        if len(extra) < 2 {
                log.Fatalf("Not enough")
        }
        structPath, err := filepath.Abs(extra[0])

        if err != nil {
                log.Fatalf("unable detect struct file path: %s: %v", structPath, err)
        }
        parsedStructs, err := parser.ParseStruct(structPath)

        if err != nil {
                log.Fatalf("Parsing Struct: %v", err)
        }
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
                                        strings.ToLower(parsedStruct.StructName + *outputSuffix + ".go"))
                                writeToFile(buf, outputPath)
                        }
                }
                if *groupPrefix != "" {
                        outputPath := filepath.Join(dir,
                                "generated", 
                                filepath.Base(tmplPath), 
                                strings.ToLower(*groupPrefix + *outputSuffix + ".go"))
                        writeToFile(bufs, outputPath)
                }
        }
}