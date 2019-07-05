package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

// DefaultFile indicates the default filepath containing problems.
// If user specifies a different filepath via arguments, this will be overriden.
const DefaultFile = "./problems.csv"

var filePath string
var problems []string

func main() {

	parseArguments()
	readProblems()

}

func parseArguments() {
	flag.StringVar(&filePath, "file", "", "path of CSV file containing problems. (default: \"./problems.csv\")")
	flag.Parse()

	if filePath == "" {
		if _, err := os.Stat(DefaultFile); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "Error: Problems file couldn't be found. Filepath must be specified")
			flag.PrintDefaults()
			os.Exit(1)
		}
		filePath = DefaultFile
	}
}

func readProblems() {
	f, _ := os.Open(filePath)
	r := csv.NewReader(bufio.NewReader(f))
	defer f.Close()

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		problems = append(record)
	}
}
