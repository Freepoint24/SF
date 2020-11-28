package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./calc"
)

func main() {

	knownOperators := []string{"+", "-", "*", "/", "^", "%"}
	var operatorsFound []string
	var numbersFound []string
	var decimals bool
	fmt.Println("Введите операцию: ")
	reader := bufio.NewReader(os.Stdin)

	operation, err := reader.ReadString('\n')

	if err != nil || operation == "\r\n" {
		fmt.Printf("Operation was not recognised as valid, %s", questionmark(err))
		return
	}

	for _, r := range operation {
		c := string(r)
		if contains(knownOperators, c) {
			operatorsFound = append(operatorsFound, c)
		}
		if c == "." {
			decimals = true
		}
	}

	//Remove trailing new-line
	operation = strings.TrimSuffix(operation, "\r\n")
	//Remove spaces if any
	operation = strings.Map(
		func(c rune) rune {
			if c == ' ' {
				return -1
			}
			return c
		}, operation)

	if len(operatorsFound) > 1 || len(operatorsFound) == 0 {
		fmt.Println("Please, enter a valid operation")
		return
	}

	operator := operatorsFound[0]
	numbersFound = strings.Split(operation, operator)
	var numbers []float64

	for i := 0; i < len(numbersFound); i++ {
		number, err := strconv.ParseFloat(numbersFound[i], 64)
		if err != nil {
			fmt.Printf("Not a valid number: %s", numbersFound[i])
			return
		}
		numbers = append(numbers, number)
	}

	calculator := calc.NewCalculator()

	result := calculator.Calculate(numbers[0], numbers[1], operator)

	if decimals {
		//Remove the trailing zeros? Doesn't work as indended
		// fmt.Printf("Result: %s", strconv.FormatFloat(result.(float64), 'f', -1, 64))
		fmt.Printf("Result: %.2f", result)
	} else {
		fmt.Printf("Result: %.0f", result)
	}
}

func contains(slice []string, value string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return true
		}
	}
	return false
}

func questionmark(err error) string {
	if err != nil {
		return err.Error()
	}
	return "The line is empty"
}
