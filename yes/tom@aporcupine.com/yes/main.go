// yes is an implementation of the standard UNIX yes utility
package main

import (
	"flag"
	"fmt"
	"strings"
)

const versionNumber = 0.1

var version = flag.Bool("version", false, "output version information and exit")

func usage() {
	fmt.Println("Usage: yes [STRING]...")
	fmt.Println("  or:  yes OPTION")
	fmt.Println("Repeatedly output a line with all specified STRING(s), or 'y'.")
	fmt.Println()
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if *version {
		fmt.Printf("Version Number: %v\n", versionNumber)
		return
	}
	out := "y"
	if len(flag.Args()) > 0 {
		out = strings.Join(flag.Args(), " ")
	}
	for {
		fmt.Println(out)
	}
}
