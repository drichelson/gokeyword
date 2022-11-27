package main

import (
	"github.com/drichelson/gokeyword"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(gokeyword.New()) }
