package main

import (
	"io/ioutil"
	"os"
)

func readFiles(folder string) []os.FileInfo {
	inputFiles, err := ioutil.ReadDir(folder)
	checkError("Cannot create file", err)

	return inputFiles
}
