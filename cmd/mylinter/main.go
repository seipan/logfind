package main

import (
	"github.com/seipan/logfind"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(logfind.Analyzer) }
