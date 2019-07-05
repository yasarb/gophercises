package main

import (
	"flag"
	"fmt"
	"os"
)

// DefaultFile indicates the default filepath containing problems.
// If user specifies a different filepath via arguments, this will be overriden.
const DefaultFile = "./problems2.csv"

var filepath string

func main() {

	parseArguments()
	fmt.Printf("Filepath: %s\n", filepath)

}

func parseArguments() {
	flag.StringVar(&filepath, "file", "", "path of CSV file containing problems.\nIf not specified, default file is \"./problems.csv\"")
	flag.Parse()

	if filepath == "" {

		a, err := os.Stat(DefaultFile)
		fmt.Println(a)
		fmt.Println(err)

		if _, err := os.Stat(DefaultFile); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "Error: Problems file couldn't be found. Filepath must be specified")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
}
