package main

import (
	"fmt"
	"log"
	"time"
)

func fibonacci(nums chan<- int) {
	a := 1
	b := 1
	nums <- 1
	nums <- 1
	for {
		current := a + b
		a = b
		b = current
		// 'Yield' by sending to the channel
		nums <- current
	}
}

func main() {
	f := make(chan int)
	go fibonacci(f)
	for {
		ans := <-f
		log.Print(ans)
		time.Sleep(1 * time.Second)
	}
}
