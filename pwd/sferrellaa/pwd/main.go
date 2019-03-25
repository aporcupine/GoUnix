// pwd is a go implementation of the unix tool "pwd"
// created 23 March 2019 by SferrellaA
// edited  23 March 2019 by SferrellaA
package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

// errFail() is a helper function to fail on errors
func errFail(err error) {
	if nil != err {
		panic(err)
	}
}

// absolute() returns the current absolute/logical PWD
func absolute() string {
	abs, err := filepath.Abs(".")
	errFail(err)
	return abs
}

func main() {

	// handle flags
	version := flag.Bool("version", false, "version information")
	logical := flag.Bool("L", false, "display logical path (allows symlinks)")
	physical := flag.Bool("P", true, "display physical path (literal path, no symlinks)")
	flag.Parse()

	// TODO handle -LP and -PL flags
	// TODO handle -l, -p flags
	// TODO handle --logical and --physical flags
	// TODO adjust usage text

	switch {
	// display version information and exit
	case *version:
		fmt.Println("GoUnix ls implementation by SferrellaA")
		return

	// display the logical (symlinks allowed) path
	case *logical:
		pwd := absolute()
		fmt.Println(pwd)
		return

	// display the physical (no symlinks) path
	case *physical:
		pwd, err := filepath.EvalSymlinks(absolute())
		errFail(err)
		fmt.Println(pwd)
		return
	}
}
