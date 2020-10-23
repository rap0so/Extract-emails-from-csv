package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
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
			if regexEmailValidator(cell) {
				matchedEmails = append(matchedEmails, cell)
			}
		}
	}

	return matchedEmails
}
