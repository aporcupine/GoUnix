// rm is a limited implementation of the standard UNIX rm utility
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/aporcupine/multiflag"
)

var (
	recursive      = multiflag.Bool(false, "remove directories and their contents recursively", "r", "R", "recursive")
	dir            = multiflag.Bool(false, "remove empty directories", "d", "dir")
	verbose        = multiflag.Bool(false, "explain what is being done", "v", "verbose")
	version        = flag.Bool("version", false, "output version information and exit")
	noPreserveRoot = flag.Bool("no-preserve-root", false, "do not treat '/' specially")
	promptEvery    = flag.Bool("i", false, "prompt before every removal")
	exitCode       int
)

func usage() {
	fmt.Println("Usage: rm [OPTION]... [FILE]...")
	fmt.Println("Remove (unlink) the FILE(s).")
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("By default, rm does not remove directories.  Use the --recursive (-r or -R)")
	fmt.Println("option to remove each listed directory, too, along with all of its contents.")
}

func main() {
	defer func() {
		os.Exit(exitCode) // Exit with value of the exitCode var rather than 0
	}()
	flag.Usage = usage
	flag.Parse()
	paths := flag.Args()
	if *version {
		fmt.Println("rm 1.0")
		fmt.Println("Limited implementation of the standard UNIX rm utility")
		fmt.Println("Written by Tom Hanson for https://github.com/aporcupine/GoUnix")
		return
	}
	if len(paths) == 0 {
		fmt.Fprintln(os.Stderr, "missing operand")
		fmt.Fprintln(os.Stderr, "Try 'rm --help' for more information.")
		exitCode = 1
		return
	}
	for _, p := range paths {
		removePath(p)
	}
}

// Removes the provided path, recursively if the recursive flag is true
func removePath(p string) {
	p = path.Clean(p)
	f, err := os.Lstat(p)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
		return
	}

	if !*noPreserveRoot && p == "/" {
		fmt.Println("it is dangerous to operate recursively on '/'")
		fmt.Println("use --no-preserve-root to override this failsafe")
		exitCode = 1
		return
	}

	if f.IsDir() && *recursive { // Descend into directory if recursive
		if *promptEvery {
			c, err := promptUser(fmt.Sprintf("descend into directory '%v'?", p))
			if err != nil {
				fmt.Fprintf(os.Stderr, "error prompting user: %v\n", err)
				exitCode = 1
				return
			}
			if !c {
				return
			}
		}
		files, _ := ioutil.ReadDir(p)
		for _, file := range files {
			removePath(fmt.Sprintf("%v/%v", p, file.Name()))
		}
	}
	if !f.IsDir() || *dir || *recursive { // Remove files and empty directories if -d or recursive
		if *promptEvery {
			ft := fileType(f)
			prompt := fmt.Sprintf("remove %v '%v'?", ft, p)
			c, err := promptUser(prompt)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error prompting user: %v\n", err)
				exitCode = 1
				return
			}
			if !c {
				return
			}
		}
		if err := os.Remove(p); err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitCode = 1
		} else if *verbose {
			t := "removed '%v'\n"
			if f.IsDir() {
				t = "removed directory '%v'\n"
			}
			fmt.Printf(t, p)
		}
	} else {
		fmt.Fprintf(os.Stderr, "cannot remove '%v': Is a directory\n", p)
		exitCode = 1
		return
	}
}

// Prompts the user to confirm whether to continue
// Returns true if response is yes or y, else false
func promptUser(prompt string) (bool, error) {
	fmt.Print(prompt + " ")
	var r string
	_, err := fmt.Scanln(&r)
	if err != nil {
		return false, err
	}
	if r == "y" || r == "yes" {
		return true, nil
	}
	return false, nil
}

// Given a os.FileInfo argument, returns the type or file
func fileType(f os.FileInfo) string {
	switch mode := f.Mode(); {
	case mode.IsRegular():
		return "regular file"
	case mode.IsDir():
		return "directory"
	case mode&os.ModeSymlink != 0:
		return "symbolic link"
	case mode&os.ModeNamedPipe != 0:
		return "named pipe"
	case mode&os.ModeDevice != 0:
		return "device file"
	default:
		return "file"
	}
}
