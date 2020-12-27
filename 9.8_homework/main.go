package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	var people = map[string]Man{
		"Саша":  {Name: "Саша", LastName: "Иванов", Age: 39, Gender: "М", Crimes: 4},
		"Антон": {Name: "Антон", LastName: "Петров", Age: 556, Gender: "М", Crimes: 1},
		"Коля":  {Name: "Коля", LastName: "Сидров", Age: 45, Gender: "М", Crimes: 5},
		"Миша":  {Name: "Миша", LastName: "Захаров", Age: 31, Gender: "М", Crimes: 12},
		"Наум":  {Name: "Наум", LastName: "Чернов", Age: 23, Gender: "М", Crimes: 3},
		"Витя":  {Name: "Витя", LastName: "Брич", Age: 36, Gender: "М", Crimes: 1},
		"Иван":  {Name: "Иван", LastName: "Захаров", Age: 34, Gender: "М", Crimes: 20},
		"Ольга": {Name: "Ольга", LastName: "Фролова", Age: 25, Gender: "Ж", Crimes: 0},
	}

	suspects := []string{
		"Саша",
		"Коля",
		"Женя",
	}

	var (
		mostCrimes   Man
		mostCrimesOk bool
	)

	for _, name := range suspects {
		human, ok := people[name]
		if !ok { //! преобразует текущее выражение в ложное
			continue
		}

		if human.Crimes > mostCrimes.Crimes {
			mostCrimes = human
			mostCrimesOk = true
			//}
		}
	}
	if mostCrimesOk {
		fmt.Printf("Наибольшее количество совершённых преступлений: \n%s %s", mostCrimes.Name, mostCrimes.LastName)
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
}
