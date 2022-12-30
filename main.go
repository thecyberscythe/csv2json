package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// Set the user-defined location for the output JSON file
	outputLocation := "./bunchofstuff.json"

	// Open the CSV files and read all of the records
	records := make(map[string]map[string]string)
	filepath :=
		"~/Desktop"
	for _, file := range []string{filepath + "Bunchofstuff.csv", "Bunchofstuff2.csv", "Bunchofstuff3.csv"} {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Error opening CSV file: %s\n", err)
			return
		}
		defer f.Close()

		r := csv.NewReader(f)
		headers, err := r.Read()
		if err != nil {
			fmt.Printf("Error reading CSV Headers: %s\n", err)
			return
		}
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Error reading CSV record: %s\n", err)
				return
			}
			// Create a map for the current record, with the headers as keys and the values as values
			recordMap := make(map[string]string)
			for i, header := range headers {
				recordMap[header] = record[i]
			}

			records[fmt.Sprintf("record%d", len(records)+1)] = recordMap
		}
	}

	// Convert the records to JSON
	jsonData, err := json.MarshalIndent(records, "", "    ")
	if err != nil {
		fmt.Printf("Error converting records to JSON: %s\n", err)
		return
	}

	// Write the JSON to a file at the output location
	f, err := os.Create(outputLocation)
	if err != nil {
		fmt.Printf("Error creating output file: %s\n", err)
		return
	}
	defer f.Close()

	_, err = f.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing JSON to output file: %s\n", err)
		return
	}
}
