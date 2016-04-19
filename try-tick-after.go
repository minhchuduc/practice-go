package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c2 := make(chan int)
	select {
	case m := <-c2:
		handle(m)
	case <-time.After(5 * time.Second):
		fmt.Println("timed out")
	}

	defer fmt.Println("Goodbye")
	// The block below will run forever because of time.Tick()
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, fmt.Sprint("Fake returN"))
		r := rand.Intn(5)
		fmt.Println(r)
		if r >= 4 {
			break
		}
	}

}

func handle(m int) {
	fmt.Println(m)
}
