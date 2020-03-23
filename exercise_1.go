package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func returnMathTuple(mathOp string) (int64, int64) {
	math := strings.Split(mathOp, "+")
	num, _ := strconv.ParseInt(math[0], 10, 64)
	num2, _ = strconv.ParseInt(math[1], 10, 64)
	if err != nil {fmt.Println(err)}

	return num, num2
}

func main() {
	fileName := flag.String("filename","","the name of the file")
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

		if err != nil {fmt.Println(err)}

		firstNumber, secondNumber := returnMathTuple(math)
		expression:= firstNumber+secondNumber == answerNumber

		if expression {counter["right"]+= 1}
		if !expression {counter["left"]+=1}
	}

	fmt.Println(counter)

}
