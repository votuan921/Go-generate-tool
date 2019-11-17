# struct-extend-generator

struct-extend-generator generate method from database struct with template using go/ast and go/generate

## Getting Started

If `struct.go` contains the `struct` types and `template.txt` you want to generate code, and assuming `GOPATH` is set to a reasonable value for an existing project (meaning that in this particular example if `struct.go` and `template.txt` is in the `myproject` directory, the project should be under `$GOPATH/src/myproject`), you can just run:
```
$ go get github.com/votuan921/struct-extend-generator
$ struct-extend-generator struct.go template.txt
```
Check `generated` folder for code genetared 

## Using struct-extend-generator

`struct-extend-generator` generates code based upon existing `struct` types with `template`.  For example, `struct-extend-generator struct.go template.txt` will by default create list new file `structName_extend.go` that contains serialization functions for all structs found in `struct.go` and have same template with `template.txt`
```
Usage of struct-extend-generator:

         struct-extend-generator [options] [structFile] [template1]..[templateN]

struct-extend-generator generates Go code for struct with template
   [options]
      -g=string group all structName extend to one file and named prefix with string(default: "")
      -o=string add suffix to all generated file with string(default: "_extend")
```
Your `structFile` code must be in a compilable state for `struct-extend-generator` to work. If you code doesn't compile `struct-extend-generator` will most likely exit with an error.

## Using struct-extend-generator with `go generate`

`struct-extend-generator` is fitable with `go generate`. It allows you to specify the `struct-extend-generator` command inside your individual struct files and run them all at once. This way you don't have to maintain a separate build file with the files you need to generate.

Add this comment anywhere inside your struct files:

```Go
//go:generate struct-extend-generator $GOFILE template.txt
```

then simply execute:

```sh
$ go generate
```

## How to create 'template' correct

`struct-extend-generator` template based on golang's package ['text/template'](https://golang.org/pkg/text/template).
Templates are executed by applying them to a data structure. Annotations in the template refer to elements of the data structure (typically a field of a struct or a key in a map) to control execution and derive values to be displayed.

```
Which variables could be used in template?
```
Execution of the template walks the structure and sets the cursor, represented by a period '.' and called "dot", to the value at the current location in the structure as execution proceeds.

Example we have parsed struct
```
type ParsedStruct struct {
        StructName string
        IDType     string
        Fields     map[string]string
}
```
then we can create template like this
```
type {{.StructName}}ID {{.IDType}}
type {{.StructName}}Slice []*{{.StructName}}
// GroupByID returns a map and a slice of given ItemSlice
func (ss {{.StructName}}Slice) GroupByID() (grouped map[{{.StructName}}ID]*{{.StructName}}, ids []{{.StructName}}ID){
    if len(ss) == 0 {
        return
    }
    grouped = make(map[{{.StructName}}ID]*{{.StructName}})
    ids = make([]{{.StructName}}ID, len(ss))
    for idx, i := range ss {
        grouped[i.Id] = i
        ids[idx] = {{.StructName}}ID(i.Id)
    }
    return
}
```

## License
 
 This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
