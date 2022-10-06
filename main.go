package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type QuizItem struct {
	Question string
	Answer   string
}

type QuizResponse struct {
	Question string
	Answer   string
}

var quizItems []QuizItem

func getQuizItems(){
args := os.Args[1:]

	

	if len(args) == 0 {
		println("No argument")
		return
	}

	csvfilename := args[0]
	println(csvfilename)

	f, err := os.Open(csvfilename)
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

		fmt.Printf("%+v\n", quizItem)
    	
	}
}

func main() {
	getQuizItems()
	fmt.Printf("%+v\n", quizItems)

}
