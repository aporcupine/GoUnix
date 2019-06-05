package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New : returns new text error
func New(text string) error {
	return &errorString{text}
}

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}

// octal : convert string to octal bits for proper dir permissions
func octal(octStr string) uint32 {
	res, _ := strconv.ParseInt(octStr, 8, 32)
	return uint32(res)
}

// CreateDirs : Create directories for the passed in os.Args
func CreateDirs(args []string) error {
	// flags
	var mode string
	var p string
	flag.StringVar(&mode, "mode", "0755", "set the file permission bits")
	flag.StringVar(&p, "p", "", "Create intermediate directories. Not implemented yet.")
	flag.Parse()

	// make sure dir/dirs not empty
	if len(args) == 0 {
		// fmt.Println("usage: mkdir [-pv] directory ...")
		return errors.New("usage: mkdir [-pv] directory ...")
	}

	for _, dir := range args {
		// create each directory
		err := os.Mkdir(dir, os.FileMode(octal(mode)))
		if err != nil {
			// fmt.Println("mkdir:", dir, "File exists")
			fmt.Printf("mkdir: %s: File exists\n", dir)
		}
	}

	return nil
}

func main() {
	err := CreateDirs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
