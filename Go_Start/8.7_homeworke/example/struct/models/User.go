package models

import (
	"fmt"
	"math/rand"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	salary    int
}

func NewUser(name string, surname string, salary int) *User {
	return &User{
		ID:        rand.Int63(),
		FirstName: name,
		LastName:  surname,
		salary:    salary,
	}
}

func (u *User) GetSalary() {
	fmt.Println(u.salary + 2000)
}
