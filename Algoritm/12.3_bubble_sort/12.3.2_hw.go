package main

import (
	"fmt"
	"math/rand"
	"time"
)

//12.3.2 алгоритм сортировки пузырьком - bubble sort, отсортирует массив int по возрастанию.
//учебный пример
//https://play.golang.org/p/jwIZN_WL8sg

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := []int{6, 5, 3, 1, 8, 7, 2, 4} // пример из учебной главы
	//ar :=[]int{1, 5, 6, 0, 10, -7, 3, 8, 4, 2, 7}
	//ar := make([]int, 100)
	//for i := range ar {
	//	ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	//}

	fmt.Printf("Unsorted list.:\t%v\n", ar) //не сортированный массив
	fmt.Println("")
	bubbleSort(ar)
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", ar) //отсортированный масссив
}

//ваш код здесь sort on the left - по возрастанию
func bubbleSort(ar []int) { //сортировка
	for i := 0; i < (len(ar) - 1); i++ { //создаем переменную i типа int и присвоили ей значение 0
		// i - количество прогонов слайса,
		// i максимальное значние длина слайса минус 2

		fmt.Println("Прогон i:", i)
		//fmt.Println(len(ar))
		for j := 0; j < ((len(ar) - 1) - i); j++ { //создаем переменную j типа int и присвоили ей значени 0
			// j - индекс элемента слайса
			//j не больше чем len - 1 - i,  (длина слайса минус 2)

			fmt.Println("Шаг сравнения j:", j, (ar))

			if ar[j] > ar[j+1] { //если значение следующего меньше, предыдущего, то
				ar[j], ar[j+1] = ar[j+1], ar[j] // меняем элементы местами
			}

		}
		fmt.Printf("Итог прогона  ....:%v\n", ar) //визуальное представление сортировки
		fmt.Println("")
	}
	fmt.Println("Длина слайса:", len(ar))
}
