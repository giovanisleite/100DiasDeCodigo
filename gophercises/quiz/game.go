package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	problems_file := flag.String("file", "problems.csv", "Path to the problems file")
	flag.Parse()

	f, err := os.Open(*problems_file)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	file_reader := csv.NewReader(f)
	answer_reader := bufio.NewReader(os.Stdin)

	fmt.Println("Quiz:")
	score, fails := 0, 0
	for {
		line, err := file_reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
            break
		}
        problem, answer := line[0], line[1]
		fmt.Println(problem)
		fmt.Print("Your answer is: ")
		user_answer, err := answer_reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
            break
		}
		if strings.TrimSpace(answer) == strings.TrimSpace(user_answer) {
			score++
		} else {
			fails++
		}
	}
	fmt.Printf("Successes: %v. Fails: %v.\n", score, fails)
}
