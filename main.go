package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	inputFolder := "./input/"
	// receive a list of files from ReadDir
	inputFiles, err := ioutil.ReadDir(inputFolder)
	checkError("Cannot create file", err)
	var emails []string

	// for each file in range of inputFiles
	for _, file := range inputFiles {
		parsedEmails := parse(inputFolder + file.Name())

		emails = append(emails,
			parsedEmails...,
		)
	}

	result := []byte(strings.Join(emails, ", \n"))

	ioutil.WriteFile("./output/result.csv", result, 0644)
}
