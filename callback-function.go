package main

import "fmt"

func main() {
	callback(4, printit)
}

func printit(x int) {
	fmt.Printf("%v\n", x)
}

func callback(y int, f func(int)) {
	y = y * y
	f(y)
}
