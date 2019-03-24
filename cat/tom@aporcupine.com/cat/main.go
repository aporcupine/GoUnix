// cat is a limited implementation of the standard UNIX cat utility
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var numberNonblank = flag.Bool("number-nonblank", false, "number nonempty output lines, overrides -n")
var showEnds = flag.Bool("show-ends", false, "display $ at end of each line")
var number = flag.Bool("number", false, "number all output lines")
var squeezeBlank = flag.Bool("squeeze-blank", false, "suppress repeated empty output lines")
var showTabs = flag.Bool("show-tabs", false, "display TAB characters as ^I")

func usage() {
	fmt.Println("Usage: cat [FLAG]... [FILE]...")
	fmt.Println("Concatenate FILE(s) to standard output")
	fmt.Println()
	fmt.Println("With no FILE, or when FILE is -, read standard input.")
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  cat f - g  Output f's contents, then standard input, then g's contents.")
	fmt.Println("  cat Copy standard input to standard output.")
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "-")
	}
	for _, arg := range args {
		file := os.Stdin
		if arg != "-" {
			var err error
			file, err = os.Open(arg)
			if err != nil {
				log.Fatal(err)
			}
		}
		if *numberNonblank || *showEnds || *number || *squeezeBlank || *showTabs {
			cat(file)
		} else {
			simpleCat(file)
		}
		file.Close()
	}
}

// Simple copy from file to stdout
func simpleCat(file io.Reader) {
	_, err := io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatal(err)
	}
}

// Read from file, complete some analysis and write to stdout
func cat(file io.Reader) {
	reader := bufio.NewReader(file)
	line := 1
	partial := false
	lastBlank := false
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil && err != bufio.ErrBufferFull && err != io.EOF {
			log.Fatal(err)
		}
		// Handle squeeze blank if requested
		if *squeezeBlank {
			lineBlank := len(b) > 0 && b[0] == '\n'
			if lastBlank && lineBlank {
				continue
			}
			lastBlank = lineBlank
		}
		// Add line numbers if required and not partial line suffix
		if (*number || *numberNonblank) && !partial && len(b) > 0 {
			if !(*numberNonblank && b[0] == '\n') {
				ln := fmt.Sprintf("%6v\t", line)
				b = append([]byte(ln), b...)
				line++
			}
		}
		partial = false
		// Add $ to end of lines if requested and partial line prefix
		if err == bufio.ErrBufferFull {
			partial = true
		} else if *showEnds && err != io.EOF {
			b = append(b[:len(b)-1], []byte("$\n")...)
		}
		// Replace tabs with ^I if requested
		if *showTabs {
			b = bytes.Replace(b, []byte("\t"), []byte("^I"), -1)
		}
		os.Stdout.Write(b)
		if err == io.EOF {
			break
		}
	}
}
