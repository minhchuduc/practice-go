package main

import "fmt"

func nextFibonacci(ch chan int) {
	a1 := <-ch
	a2 := <-ch
	fmt.Println(a2)
	ch <- a2
	ch <- a1 + a2
}

func Fibochain(len int) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 1
	for i := 0; i < len; i++ {
		nextFibonacci(ch)
	}
}

func main() {
	i := 9
	Fibochain(i)
}
