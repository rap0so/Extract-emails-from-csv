package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	inputFolder := "./input/"
	var emails []string

	// for each file in range of inputFiles
	for _, file := range readFiles(inputFolder) {
		parsedEmails := parse(inputFolder + file.Name())

		emails = append(emails,
			parsedEmails...,
		)
	}

	result := []byte(strings.Join(emails, ", \n"))

	ioutil.WriteFile("./output/result.csv", result, 0644)
}
