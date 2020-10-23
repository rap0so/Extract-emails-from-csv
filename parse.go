package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"regexp"
)

func parse(fileName string) []string {
	csvFile, error := os.Open(fileName)
	checkError("Cannot open "+fileName, error)
	reader := csv.NewReader(csvFile)

	var matchedEmails []string

	for {
		valueOnRow, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		for _, cell := range valueOnRow {
			emailMatch, _ := regexp.MatchString(".*@.*", cell)
			if emailMatch {
				matchedEmails = append(matchedEmails, cell)
			}
		}
	}

	return matchedEmails
}
