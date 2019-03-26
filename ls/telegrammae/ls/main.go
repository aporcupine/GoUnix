package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const versionNumber = 0.1

var version = flag.Bool("v", false, "output version information")
var all = flag.Bool("a", false, "list all items beginning with '.'")

func usage() {
	fmt.Println("Usage:")
	fmt.Println("List the contents of a directory.")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	flag.Parse()
	if *version {
		fmt.Printf("Version Number: %v\n", versionNumber)
		return
	}

	var dir string
	lenArgs := len(os.Args)
	if lenArgs > 1 {
		dir = os.Args[lenArgs-1]
	} else {
		dir = "."
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range files {
		if !*all && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		fmt.Println(f.Name())
	}
}
