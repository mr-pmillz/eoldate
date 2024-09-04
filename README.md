# eoldate

[![Go Report Card](https://goreportcard.com/badge/github.com/mr-pmillz/eoldate)](https://goreportcard.com/report/github.com/mr-pmillz/eoldate)
![GitHub all releases](https://img.shields.io/github/downloads/mr-pmillz/eoldate/total?style=social)
![GitHub repo size](https://img.shields.io/github/repo-size/mr-pmillz/eoldate?style=plastic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mr-pmillz/eoldate?style=plastic)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mr-pmillz/eoldate?style=plastic)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/mr-pmillz/eoldate?style=plastic)
[![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fmr-pmillz%2Feoldate)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fmr-pmillz%2Feoldate)
[![CI](https://github.com/mr-pmillz/eoldate/actions/workflows/ci.yml/badge.svg)](https://github.com/mr-pmillz/eoldate/actions/workflows/ci.yml)

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