# go-folder

[![GoDoc](https://godoc.org/github.com/unjx-de/go-folder?status.svg)](https://godoc.org/github.com/unjx-de/go-folder)
[![Go](https://github.com/unjx-de/go-folder/actions/workflows/go.yml/badge.svg)](https://github.com/unjx-de/go-folder/actions/workflows/go.yml)
[![CodeQL](https://github.com/unjx-de/go-folder/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/unjx-de/go-folder/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/unjx-de/go-folder)](https://goreportcard.com/report/github.com/unjx-de/go-folder)

## Usage

```go
package main

import (
	"fmt"
	folder "github.com/unjx-de/go-folder"
	"os"
)

func main() {
	err := folder.CreateFolders([]string{"exampleFolder", "anotherFolder"}, 0755)
	if err != nil {
		panic(err)
	}

	err = folder.RemoveFolders([]string{"exampleFolder", "anotherFolder"})
	if err != nil {
		panic(err)
	}
}
```
