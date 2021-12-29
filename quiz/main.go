package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse "+filePath, err)
	}

	return records
}

func getAnswer(question string) string {
	fmt.Println("Question: " + question)
	var answer string
	fmt.Scan(&answer)
	answer = strings.ToLower(strings.TrimSpace(answer))
	return answer
}

type parameters struct {
	source  string
	timeout int
}

func getPars() parameters {
	var p parameters
	flag.StringVar(&p.source, "source", "problems.csv", "Input file")
	flag.IntVar(&p.timeout, "timeout", 30, "Timeout in seconds")
	flag.Parse()
	return p
}

type quiz struct {
	question string
	answer   string
}

func loadQuiz(source string) []quiz {
	data := readCsvFile(source)
	var q []quiz
	for _, line := range data {
		q = append(q, quiz{
			question: line[0],
			answer:   line[1],
		})
	}
	return q
}

func main() {
	log.Println("Start")
	log.Println("Read args")

	p := getPars()

	log.Println("Reading from " + p.source)
	q := loadQuiz(p.source)

	correct := 0
	wrong := 0
	for _, this_q := range q {

		// read answer
		resultCh := make(chan string)
		timeoutCh := time.After(time.Duration(p.timeout) * time.Second)

		go func() {
			resultCh <- getAnswer(this_q.question)
		}()

		select {
		case answer := <-resultCh:
			// tracking
			if this_q.answer == answer {
				correct++
			} else {
				wrong++
			}
		case <-timeoutCh:
			fmt.Println("Timeout!")
			wrong++
		}
	}
	fmt.Println("Correct: " + fmt.Sprint(correct))
	fmt.Println("Wrong: " + fmt.Sprint(wrong))
}
