package main

import (
	"os"

	"github.com/vivekprm/go-best-practices/pkg/gopher"
)

func main() {
	gopher := &gopher.Gopher{
		Name:     "vivek",
		AgeYears: 38,
	}
	gopher.WriteTo(os.Stdout)
}
