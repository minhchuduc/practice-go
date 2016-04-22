// Variadic arg
package main

import "fmt"

func f(args ...int) {
	for _, val := range args {
		fmt.Println(val)
	}
	fmt.Println("------------------------------")
}

func main() {
	f(1, 2)
	f(3, 4, 5, 6)
}
