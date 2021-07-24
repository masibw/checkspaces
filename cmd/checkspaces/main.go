package main

import (
	"github.com/masibw/checkspaces"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(checkspaces.Analyzer) }
