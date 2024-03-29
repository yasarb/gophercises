package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

// DefaultFile indicates the default filepath containing problems.
// If user specifies a different filepath via arguments, this will be overriden.
const DefaultFile = "./problems.csv"

// Time limit for a quiz in seconds
const DefaultTimeLimit = 30

var filePath string
var timeLimit int
var problems []Problem
var score int

func main() {

	parseArguments()
	readProblems()
	startQuiz()

}

func parseArguments() {
	flag.StringVar(&filePath, "file", DefaultFile, "Path of CSV file containing problems")
	flag.IntVar(&timeLimit, "limit", DefaultTimeLimit, "Time limit for a quiz in seconds")
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

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	go onTimerExpired(timer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(problems); i++ {
		fmt.Printf("Problem #%d: %-4s= ? ", i+1, problems[i].question)
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)

		if answer == problems[i].answer {
			score++
		} else {
			fmt.Printf("You scored %d out of %d", score, len(problems))
			timer.Stop()
			os.Exit(1)
		}
	}

	fmt.Printf("Perfect! You scored %d out of %d", score, len(problems))
	timer.Stop()
}

func onTimerExpired(timer *time.Timer) {
	func() {
		<-timer.C
		fmt.Printf("\nYou scored %d out of %d\n", score, len(problems))
		os.Exit(1)
	}()
}
