package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/glaslos/tlsh"
)

var (
	// VERSION is set by the makefile
	VERSION = "v0.0.0"
	// BUILDDATE is set by the makefile
	BUILDDATE = ""
)

func main() {
	var file = flag.String("f", "", "path to the `file` to be hashed")
	var raw = flag.Bool("r", false, "set to get only the hash")
	var version = flag.Bool("version", false, "print version")
	flag.Parse()
	if *version {
		fmt.Printf("%s %s\n", VERSION, BUILDDATE)
		return
	}
	if *file == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s [-f <file>]\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println()
		return
	}
	hash, err := tlsh.Hash(*file)
	if err != nil {
		fmt.Println(err)
		return
	}
	if *raw {
		fmt.Println(hash)
	} else {
		fmt.Printf("%s  %s\n", hash, *file)
	}
}
