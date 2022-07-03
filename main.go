package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
	var length int = 0
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", rec[0])

		var answer string

		fmt.Scanln(&answer)
		length++

		if answer == rec[1] {
			score++
		}
	}

	fmt.Printf("Good job, your total score is: %d out of %d", score, length)
}
