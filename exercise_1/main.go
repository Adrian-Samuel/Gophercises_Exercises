package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func stringToNumber(mathArg string) int64 {
	num, err := strconv.ParseInt(mathArg, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return num
}

func main() {
	fileName := flag.String("filename", "", "the name of the file")
	flag.Parse()
	content, err := ioutil.ReadFile(*fileName)

	if err != nil {
		fmt.Println(err)
	}

	csv := string(content)

	row := strings.Split(csv, "\n")

	counter := map[string]int{"right": 0, "wrong": 0}

	for _, value := range row {

		record := strings.Split(value, ",")
		math, answer := record[0], record[1]

		answerNumber, err := strconv.ParseInt(answer, 10, 64)

		if err != nil {
			fmt.Println(err)
		}
		mathArguments := strings.Split(math, "+")

		firstNumber := stringToNumber(mathArguments[0])
		secondNumber := stringToNumber(mathArguments[1])

		checkSum := firstNumber+secondNumber == answerNumber

		if checkSum {
			counter["right"] += 1
		}
		if !checkSum {
			counter["left"] += 1
		}
	}

	fmt.Println(counter)

}
