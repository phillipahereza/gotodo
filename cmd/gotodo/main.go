package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/phillipahereza/gotodo"
)

func main() {
	singlechecker.Main(gotodo.Analyzer)
}