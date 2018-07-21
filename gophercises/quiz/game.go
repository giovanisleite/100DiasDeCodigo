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

func main() {
	problems_file := flag.String("file", "problems.csv", "Path to the problems file")
    timelimit := flag.Int("timelimit", 30, "Integer - Timelimit (seconds) to answer the questions")
	flag.Parse()

	f, err := os.Open(*problems_file)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	file_reader := csv.NewReader(f)
	answer_reader := bufio.NewReader(os.Stdin)

    timer := time.NewTimer(timelimit * time.Second)
    score := 0
	fmt.Println("The quiz will start soon just as the timer, press any button to start the game.")
    _, _ = answer_reader.ReadString('\n')
    go func(){
        <-timer.C
        exit(fmt.Sprintf("You scored %v.", score))

    }()
	for {
		line, err := file_reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
            exit(err)
		}
        problem, answer := line[0], line[1]
		fmt.Println(problem)
		fmt.Print("Your answer is: ")
		user_answer, err := answer_reader.ReadString('\n')
		if err != nil {
			exit(err)
		}
		if strings.TrimSpace(answer) == strings.TrimSpace(user_answer) {
			score++
		}
	}
    timer.Stop()
	fmt.Printf("You scored: %v.\n", score)
}


func exit(message string){
    fmt.Printf("\n%s\n", message)
    os.Exit(1)
}
