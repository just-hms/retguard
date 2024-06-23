package main

import (
	"github.com/just-hms/retguard/pkg/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analysis.RetGuard)
}
