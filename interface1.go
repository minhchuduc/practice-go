package main

import "fmt"

func add(a interface{}, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func main() {
	fmt.Println(add(5, 7))
	stra := "the spice must flow"
	byts := []byte(stra)
	strb := string(byts)
	fmt.Println(stra, byts, strb)
}
