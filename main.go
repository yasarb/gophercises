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

func main() {

	parseArguments()

	f, _ := os.Open(filePath)
	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

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
