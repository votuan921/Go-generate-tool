# struct-extend-generator

struct-extend-generator parses the given go structs and generates new file based on given template and parsed structs.

# Usage
- Firstly, the go structs file and template file can be created as you want. Especially you have structs of tables of database schema, then some utility methods like group by ID, or filter to get slice of colum.
- Then install the tool:
```bash
$go install github.com/votuan921/struct-extend-generator
```
- And generates what you want by given go structs file and template path:
```
Extended methods generator for Go struct.
	Usage: $struct-extend-generator [options] /absolute/path/to/structs_file.go /absolute/path/to/template.tpl
	Options:
			-e: output file extension. Default: ".extend.go"
			-o: absolute path to output file directory. Default is struct file dir.
```
- Usage of parsed structs in template could be:
```
// ParsedStruct represents a struct of a parsed struct
type ParsedStruct struct {
	StructName string
	IDType     string
	Fields     map[string]string //key: Field name, value: Field type
}
```
-  Refer to example dir to more details of usage.

## Example
```bash
$go get -u github.com/votuan921/struct-extend-generator
$cd $GOPATH/src/github.com/votuan921/struct-extend-generator
$struct-extend-generator $GOPATH/src/github.com/votuan921/struct-extend-generator/example/struct.go.example $GOPATH/src/github.com/votuan921/struct-extend-generator/example/struct.tpl
```

## License
 
 This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
