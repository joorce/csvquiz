package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type QuizItem struct {
	Question string
	Answer   string
}

type QuizResponse struct {
	Question  string
	Answer    string
	IsCorrect bool
}

var quizItems []QuizItem

func getQuizItems() {
	filename := flag.String("f", "problems.csv", "the file with questions")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.Read()

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quizItem := QuizItem{row[0], row[1]}
		quizItems = append(quizItems, quizItem)

	}
}

var answers = make([]QuizResponse, 0)
var correctAnswers = 0
var incorrectAnswers = 0

func askQuestions() {
	fmt.Println("Answer the questions")
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(quizItems); i++ {
		quizItem := quizItems[i]
		fmt.Printf("%v: ", quizItem.Question)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		answer := strings.TrimRight(input, "\r\n")
		isCorrect := quizItem.Answer == answer

		if isCorrect {
			correctAnswers++
		} else {
			incorrectAnswers++
		}

		answers = append(answers, QuizResponse{
			Question:  quizItem.Question,
			Answer:    answer,
			IsCorrect: isCorrect})

		fmt.Printf("Correct: %v, Incorrect: %v\n", correctAnswers, incorrectAnswers)
	}
}

func main() {

	getQuizItems()
	askQuestions()
	fmt.Printf("%v correct answers from %v questions\n", correctAnswers, len(quizItems))
}
