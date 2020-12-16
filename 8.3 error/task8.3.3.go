package main

import (
	"errors"
	"fmt"
)

var (
	veryExpensiveCarErr = errors.New("слишком дорогая машина")
	carNotExistErr      = errors.New("машина не существует")
)

type car struct {
	brand string
	model string
}

func main() {
	car, err := buyCar("BMW", "X6", 3000000)
	switch err {
	case veryExpensiveCarErr:
		fmt.Println("Нужно поднакопить денег")
	case carNotExistErr:
		fmt.Println("Выберите другую машину")
	case nil:
		fmt.Printf("Поздравляем с покупкой %s %s", car.brand, car.model)
	}
}

func buyCar(desiredBrand, desiredModel string, moneyAmount int) (car, error) {
	carPrice, err := getCarPrice(desiredBrand, desiredModel)
	if err != nil {
		return car{}, err
	}
	if carPrice > moneyAmount {
		return car{}, veryExpensiveCarErr
	}
	return car{brand: desiredBrand, model: desiredModel}, nil
}

func getCarPrice(brand, model string) (int, error) {
	if brand == "BMW" && model == "X6" {
		return 6000000, nil
	}
	if brand == "Audi" && model == "A7" {
		return 4000000, nil
	}
	if brand == "Mazda" && model == "6" {
		return 2000000, nil
	}
	return 0, carNotExistErr
}