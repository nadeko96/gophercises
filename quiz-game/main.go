package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Problem struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func readCSVFile() []Problem {
	csvFile, _ := os.Open("problems.csv")
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var problems []Problem
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		problems = append(problems, Problem{
			Question: line[0],
			Answer:   line[1],
		})
	}
	return problems
}

func askQuestions(problems []Problem) (int, int) {
	var count int

	ticker := time.NewTicker(time.Second * 5)
	for i := range problems {
		ch := make(chan bool)
		go func() {
			var r int

			answer, err := strconv.Atoi(problems[i].Answer)
			if err == nil {
				fmt.Println(problems[i].Question)
				fmt.Scanf("%d", &r)
				if r == answer {
					count++
					fmt.Println("Correct")
				} else {
					fmt.Println("Wrong!", r)
				}

			}
			ch <- true
		}()
		select {
		case <-ch:
			return
		case <-ticker.C:
			fmt.Println("Too slow")
		}
	}

	return count, len(problems)
}

func main() {
	problems := readCSVFile()
	correct, total := askQuestions(problems)

	fmt.Printf("You got %d out of %d\n", correct, total)
}
