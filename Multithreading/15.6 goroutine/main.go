package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("Hello, goroutine!")
			time.Sleep(time.Second)
		}()
	}
	time.Sleep(time.Second)
}

//Здесь три горутины выводят сообщения и ждут 1 секунду.
//Но программа завершит работу не через 4 секунды, а примерно через 1, так как горутины работают параллельно.
