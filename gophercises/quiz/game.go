package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	problems_file := flag.String("file", "problems.csv", "String - Path to the problems file")
	timelimit := flag.Int("timelimit", 30, "Integer - Timelimit (seconds) to answer the questions")
	shuffle := flag.Bool("shuffle", false, "Bool - Change the orger of the questions")
	flag.Parse()

	problems := ReadProblems(*problems_file)

	if *shuffle {
		problems = Shuffle(problems)
	}

	answer_reader := bufio.NewReader(os.Stdin)

	fmt.Println("The quiz will start soon just as the timer, press any button to start the game.")
	_, _ = answer_reader.ReadString('\n')

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	score := 0
	go func() {
		<-timer.C
		exit(fmt.Sprintf("You scored %v out of %v.\n", score, len(problems)))
	}()

	for _, problem := range problems {
		question, answer := problem[0], problem[1]
		fmt.Println(question)
		fmt.Print("Your answer is: ")
		user_answer, err := answer_reader.ReadString('\n')
		if err != nil {
			exit("Cannot read the user answer")
		}
		if strings.TrimSpace(answer) == strings.TrimSpace(user_answer) {
			score++
		}
	}
	timer.Stop()
	fmt.Printf("You scored %v out of %v.\n", score, len(problems))
}

func ReadProblems(problems_filename string) [][]string {
	f, err := os.Open(problems_filename)
	defer f.Close()
	if err != nil {
		exit("Cannot open the file")
	}

	file_reader := csv.NewReader(f)
	problems, err := file_reader.ReadAll()
	if err != nil {
		exit("Cannot read the file")
	}
	return problems
}

func Shuffle(problems [][]string) [][]string {
	perm := rand.Perm(len(problems))
	shuffled := make([][]string, len(problems))
	for i, v := range perm {
		shuffled[i] = problems[v]
	}
	return shuffled
}

func exit(message string) {
	fmt.Printf("\n%s\n", message)
	os.Exit(1)
}
