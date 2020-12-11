package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := joinStringsWithSpace("John", "Smith")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result, err = joinStringsWithSpace("", "apple")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func joinStringsWithSpace(str1, str2 string) (string, error) {
	if str1 == "" {
		return "", errors.New("первая строка пустая")
	}
	if str2 == "" {
		return "", errors.New("вторая строка пустая")
	}

	return str1 + " " + str2, nil
}