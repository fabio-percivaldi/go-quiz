package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	filename := flag.String("filename", "problem.csv", "name of the file containing quiz questions")
	flag.Parse()
	fmt.Println(*filename)

	f, err := os.Open(*filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	var score int = 0
	filedata, err := csvReader.ReadAll()
	totalQuestions := len(filedata)
	endTime := time.Now().Add(10 * time.Second)
loop:
	for _, rec := range filedata {
		fmt.Printf("%+v\n", rec[0])

		answer := make(chan string, 1)

		go func() {
			var userInput string
			fmt.Scanln(&userInput)
			answer <- userInput
		}()
		t := time.Now()
		remainingTime := endTime.Sub(t)
		fmt.Printf("time left: %f\n", math.Round(remainingTime.Seconds()))

		select {
		case res := <-answer:
			if res == rec[1] {
				score++
			}
		case <-time.After(time.Duration(remainingTime.Seconds()) * time.Second):
			fmt.Println("timeout reached")
			break loop
		}
	}

	fmt.Printf("Good job, your total score is: %d out of %d", score, totalQuestions)
}
