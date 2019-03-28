package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const versionNumber = 0.1

var version = flag.Bool("v", false, "output version information")
var path = flag.String("p", "", "create a path of directories")

func usage() {
	fmt.Println("Usage:")
	fmt.Println("Create a directory.")
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
		usage()
		return
	}

	if *path == "" {
		err := os.Mkdir(dir, os.ModePerm)
		handleCreate(err)
	} else {
		// TODO Handle Windows paths
		dirs := strings.Split(*path, "/")
		var p string
		for _, sub := range dirs {
			p += sub + "/"
			err := os.Mkdir(p, os.ModePerm)
			handleCreate(err)
		}
	}
}

func handleCreate(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
