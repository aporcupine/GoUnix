// implementation of the UNIX command "tree"

// currently, this only implements one level down
// recursively add in a counter for proper spacing depth
// and then recursively call the subDir function in itself

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	subDirToSkip := "bbb"
	fileCount := numFilesInDir(".")
	fileNum := 1

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip the entire dir if it's the specified flag entered
		if info.IsDir() && info.Name() == subDirToSkip {
			return filepath.SkipDir
		}

		// check for subdirectory
		if info.IsDir() == true && info.Name() != "." {
			// walk into the subdirectory
			subDir(info.Name())
			// skip the subdirectory as it's already been outputtted
			// within the function
			return filepath.SkipDir
		}

		if info.Name() == "." {
			fmt.Println(".")
		} else if fileNum != fileCount {
			fmt.Println("├──", info.Name())
		} else {
			fmt.Println("└──", info.Name())
		}

		fileNum++

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %v\n", err)
		return
	}
}

// numFilesInDir : counts the number of files in a directory
func numFilesInDir(dirName string) int {
	dir := "./" + dirName
	files, _ := ioutil.ReadDir(dir)
	// fileCount := len(files)
	return len(files)
}

// subDir : print out the files of a sub directory with Walk/WalkFunc
func subDir(dirName string) {
	fileCount := numFilesInDir(dirName)
	fileNum := 1

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Println("├──", info.Name())
		} else if fileNum != fileCount {
			fmt.Println("│   └──", info.Name())
		} else {
			fmt.Println("│   ├──", info.Name())
		}

		fileNum++

		return nil
	})

	if err != nil {
		return
	}
}
