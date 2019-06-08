// pwd is a go implementation of the unix tool "pwd"
// created 23 March 2019 by SferrellaA
// edited  28 March 2019 by SferrellaA
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

// usage() defines the behavior of `pwd -h`
func usage() {
	fmt.Println("Usage: pwd [-LP]")
	fmt.Println("   Print the name of the current working directory")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("   -L Logical path (symlinks allowed) (default true)")
	fmt.Println("   -P Physical/literal path (no symlinks)")
	fmt.Println("")
	fmt.Println("`pwd` assumes `-L` if no flag is given")
	fmt.Println("`-P` will override `-L` if both are given")
}

// getFlags() configures and parses flags, returns pwd mode
func getFlags() bool {

	// Additional, hidden flags prevent user headache

	// Logical is assumed by default, so value not stored
	flag.Bool("L", false, "")
	flag.Bool("l", false, "")
	flag.Bool("logical", false, "")

	// Physical trumps logical
	P := flag.Bool("P", false, "")
	p := flag.Bool("p", false, "")
	physical := flag.Bool("physical", false, "")

	// Read flags
	flag.Usage = usage
	flag.Parse()

	// Return true for -P, false for -L
	if *P || *p || *physical {
		return true
	}
	return false
}

// absolute() returns the current absolute/logical PWD
func absolute() string {
	abs, err := filepath.Abs(".")
	errFail(err)
	return abs
}

// PWD() provides `pwd` functionality as a callable function
func PWD(physical bool) string {

	// Get the absolute (logical) path
	pwd := absolute()

	// If change logical to physical if argued for
	if physical {
		var err error
		pwd, err = filepath.EvalSymlinks(pwd)
		errFail(err)
	}

	// Return pwd
	return pwd
}

// main() calls on existing functions
func main() {
	mode := getFlags()
	pwd := PWD(mode)
	fmt.Println(pwd)
}
