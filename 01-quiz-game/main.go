package _1_quiz_game

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Problem struct {
	question string
	answer   string
}

// DefaultFile indicates the default filepath containing problems.
// If user specifies a different filepath via arguments, this will be overriden.
const DefaultFile = "./problems.csv"

var filePath string
var problems []Problem
var score int

func main() {

	parseArguments()
	readProblems()
	startQuiz()

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

		problems = append(problems, Problem{record[0], record[1]})
	}
}

func startQuiz() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(problems); i++ {
		fmt.Printf("Problem #%d: %-4s= ? ", i+1, problems[i].question)
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)

		if answer == problems[i].answer {
			score++
		} else {
			fmt.Printf("You scored %d out of %d", score, len(problems))
			break
		}
	}

	fmt.Printf("Perfect! You scored %d out of %d", score, len(problems))
}
