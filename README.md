# eoldate

## About

An End of Life Date API SDK written in Go

This is a wrapper around the endoflife.date API
[Read the Docs](https://endoflife.date/docs/api)

## Installation

To install, just run the below command or download pre-compiled binary from the [releases page](https://github.com/mr-pmillz/eoldate/releases)

```bash
go install -v github.com/mr-pmillz/eoldate/cmd/eoldate@latest
```

## eoldate as a library

Integrate eoldate with other go programs

```go
package main

import (
	"fmt"
	"github.com/mr-pmillz/eoldate"
)

func main() {
	client := eoldate.NewClient(eoldate.EOLBaseURL)
	releaseVersions, err := client.GetProduct("php")
	if err != nil {
		panic(err)
	}
	fmt.Println(releaseVersions)
}
```