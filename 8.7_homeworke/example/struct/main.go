package main

import "models/models"

func main() {
	user := models.NewUser("Bulat", "Zamalat", 100)
	user.GetSalary()
}
