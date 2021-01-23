package main

import (
	"errors"
	"fmt"
	"log"
)

func main()  {
	//fmt.Println(errWrapExample())
	log.Fatal(errWrapExample())
}

func errWrapExample() error {
	err := errExample()

	return fmt.Errorf("wropad error %w", err)
}

func errExample() error {
	return errors.New("err example")
}