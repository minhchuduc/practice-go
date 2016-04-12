package main

import "fmt"

type Quackable interface {
	Quack()
	Feathers()
}

type Duck int

func (this Duck) Quack() {
	fmt.Println("Quaaaaaack!")
}

func (this Duck) Feathers() {
	fmt.Println("The duck has white and gray feathers.")
}

type Person int

func (this Person) Quack() {
	fmt.Println("The person imitates a duck.")
}

func (this Person) Feathers() {
	fmt.Println("The person takes a feather from the ground and shows it.")
}

// May also be defined as func InTheForest(q Quackable) for compile-time checking
func InTheForest(i interface{}) {
  q := i.(Quackable)

	q.Quack()
	q.Feathers()
}

func main() {
	var d Duck
	var p Person

	InTheForest(d)
	InTheForest(p)
}
