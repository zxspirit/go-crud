package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ch2 <- i * i
			time.Sleep(time.Second)
		}
		close(ch2)
	}()

	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 is closed")
				return
			}
			fmt.Println("ch1:", v)
		case v, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 is closed")
				return
			}
			fmt.Println("ch2:", v)
		}
	}
}
