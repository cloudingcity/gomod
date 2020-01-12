# Gomod

[![Build Status](https://travis-ci.com/cloudingcity/gomod.svg?branch=master)](https://travis-ci.com/cloudingcity/gomod)
[![codecov](https://codecov.io/gh/cloudingcity/gomod/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/gomod)


Parse `go.mod` into a struct

## Install

```
go get -u github.com/cloudingcity/gomod
```

## Example

```go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/cloudingcity/gomod"
)

func main() {
	file, err := ioutil.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}

	mod, err := gomod.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", mod)
}
```

## Struct
```go
type Module struct {
	Path    string
	Version string
}

type GoMod struct {
	Module  Module
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
}

type Require struct {
	Path     string
	Version  string
	Indirect bool
}

type Replace struct {
	Old Module
	New Module
}
```
