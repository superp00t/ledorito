package main

import (
	"fmt"
	"os"

	"github.com/ogier/pflag"
	"github.com/superp00t/ledorito"
)

func usage() {
	fmt.Printf("usage: %s <input .dat file> <output chunk directory>\n", os.Args[0])
	os.Exit(0)
}

func main() {
	pflag.Parse()

	in := pflag.Arg(0)
	out := pflag.Arg(1)

	if in == "" || out == "" {
		usage()
	}

	if err := ledorito.Extract(in, out); err != nil {
		fmt.Println(err)
	}
}
