# struct-extend-generator

Mini tool generate method from database struct with template using go/ast and go/generate

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them
- Go[1.12+](https://golang.org/dl/)
- Go's Package:
```
   "log"
   "go/token"
   "go/parser"
   "go/ast"
   "bytes"
   "bufio"
   "os"
```

### Installing

To get started, clone the repo and then install the needed packages:

```
$ git clone git@github.com:votuan921/struct-extend-generator.git
$ cd struct-extend-generator
```

Replace file db.go to your database file you want generate or change PARSER_FILE_PATH to your file name in parser.go and run command

```
$ go run parser.go
```

### Incoming features

- Generate with template and command
- Compress to lib file
