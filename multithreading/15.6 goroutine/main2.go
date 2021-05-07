package main

//У горутин нет идентификатора, по которому мы могли бы к ним обращаться.
//Однако мы можем обмениваться данными с помощью встроенного средства Golang — каналов (channels).
import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	fmt.Println(<-c + <-c) // 12
}
