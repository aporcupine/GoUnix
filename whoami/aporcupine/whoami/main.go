// whoami is an implementation of the standard UNIX whoami utility
package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
)

const versionNumber = 0.1

var version = flag.Bool("version", false, "output version information and exit")

func usage() {
	fmt.Println("Usage: whoami [OPTION]...")
	fmt.Println("Print the user name associated with the current effective user ID.")
	fmt.Println("Same as id -un.")
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
	cu, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cu.Username)
}
