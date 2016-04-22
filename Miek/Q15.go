package main

import "fmt"

func main() {
	p2 := plusX(2)
	fmt.Println(p2(5))
}

func plusX(x int) func(int) int {
	f := func(y int) int {
		return x + y
	}
	return f
}
