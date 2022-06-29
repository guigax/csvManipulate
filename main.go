package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Output struct {
	Line string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLine(input [][]string, columnToKeep int) []string {
	var lines []string
	for _, line := range input {
		lines = append(lines, line[columnToKeep])
	}
	return lines
}

func main() {
	csvFilename := flag.String("csv", "example.csv", "a csv file that you wish to keep only one column in it")
	columnToKeep := flag.Int("column", 0, "column to remain on the resulting csv file")

	flag.Parse()

	// open the csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit("Failed to open the CSV file: " + *csvFilename)
	}

	// close the file at the end of the execution
	defer file.Close()

	// read the csv using csv.Reader
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	data, err := csvReader.ReadAll()
	if err != nil {
		exit("Failed to parse the data of the CSV file: " + err.Error())
	}

	// convert records to array of struct
	output := parseLine(data, *columnToKeep)

	newFile, err := os.Create("output.csv")
	if err != nil {
		exit("Failed to create new file")
	}
	defer newFile.Close()

	w := bufio.NewWriter(newFile)
	for _, line := range output {
		fmt.Fprintln(w, line)
	}
	w.Flush()
}
