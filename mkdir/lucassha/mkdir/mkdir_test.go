package main

import (
	"os"
	"testing"
)

func TestCreateDir(t *testing.T) {

	t.Run("Creating 2 directories", func(t *testing.T) {
		testDirs := []string{"test1", "test2"}
		err := CreateDirs(testDirs)
		if err == nil {
			for _, dir := range testDirs {
				err := os.Remove(dir)
				if err != nil {
					t.Errorf("directories not created")
				}
			}
		}
	})

}
