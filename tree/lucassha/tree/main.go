// implementation of the UNIX command "tree"

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	subDirToSkip := "bbb"

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip the entire dir if it's the specified flag entered
		if info.IsDir() && info.Name() == subDirToSkip {
			return filepath.SkipDir
		}

		if info.IsDir() == true && info.Name() != "." {
			subDir(info.Name())
			// filepath.Walk(info.Name(), subTraversal)
			return filepath.SkipDir
		}

		if info.Name() == "." {
			fmt.Println(".")
		} else {
			fmt.Println("└──", info.Name())
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %v\n", err)
		return
	}
}

// need to keep track of what depth you're at
func subDir(dirName string) {
	dir := "./" + dirName
	files, _ := ioutil.ReadDir(dir)
	fileCount := len(files)
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

// subTraversal : WalkFunc implemenetation that steps into subdirectories
// and prints all files
func subTraversal(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// files, _ := ioutil.ReadDir(".")
	// count := len(files)
	// fmt.Println("count:", count)

	if info.IsDir() {
		fmt.Println("├──", info.Name())
	} else {
		fmt.Println("│   ├──", info.Name())
	}

	return nil
}

// ├
// └
// ─

// .
// ├── aaa
// │   ├── a.txt
// │   └── b.txt
// ├── filepathExp.go
// └── main.go
