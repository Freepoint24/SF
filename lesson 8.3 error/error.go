package main

import (
	"errors"
	"fmt"
)

func NewCustomErr(message string) *CustomErr {
	return &CustomErr{message: message}
}

type CustomErr struct {
	message string
}

func (e CustomErr) Error() string {
	return e.message
}

func main() {
	//i, err := someFunc(0)
	//fmt.Println(i, err)
	//
	//i, err = someFunc(10)
	//fmt.Println(i, err)

	//testIs()
	testAs()
}

func testIs() {
	var errNotFound = errors.New("not found")

	err := fmt.Errorf("wrap: %w", errNotFound)

	//  тоже самое, что:
	//   if err == errNotFound { … }
	if errors.Is(err, errNotFound) {
		fmt.Println("got error with Is")
	}
}


func testAs() {
	err := NewCustomErr("some error message")
	var pattern *CustomErr

	//  тоже самое, что:
	//   if e, ok := err.(CustomErr); ok { … }
	if errors.As(err, &pattern) {
		fmt.Println("got error with As")
	}
}

func someFunc(i int) (int, error) {
	j, err := funcReturningError(i)
	if err != nil {
		return j, fmt.Errorf("wrap error: %w", err)
	}

	return i, nil
}

func funcReturningError(i int) (int, error) {
	if i == 0 {
		return 0, fmt.Errorf("got %d", i)
	}

	return i, nil
}
