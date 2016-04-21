package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := &Saiyan{"Goku", 9000}
	/*
	   goku := new(Saiyan)
	   goku.Power = 9000
	*/
	Super(goku)
	fmt.Println(goku.Power)
	goku.Super2()
	fmt.Println(goku.Power)

	goku2 := &Saiyan2{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku2.Introduce()
	goku2.Person.Introduce() // Function does NOT overloading
	fmt.Println(goku2.Name)
	fmt.Println(goku2.Person.Name) // the same as above line
}

func Super(s *Saiyan) { s.Power += 10000 }

func (s *Saiyan) Super2() { s.Power += 5000 }

func NewSaiyan(name string, power int) *Saiyan { // a type that "pointer to value of Saiyan type"
	return &Saiyan{ // address of a value have type is Saiyan
		Name:  name,
		Power: power,
	}
}

type Person struct {
	Name string
}

func (p *Person) Introduce() { // p as receiver of Introduce() func
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan2 struct {
	*Person // Composition
	Power   int
}

func (s *Saiyan2) Introduce() { // this func will overwrite func of composed type (Person)
	fmt.Printf("Hi, I'm %s. Yaaaaaaaaaaa!\n", s.Name)
}

func extractPowers(saiyans []*Saiyan) []int {
	powers := make([]int, len(saiyans))
	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}
func extractPowers2(saiyans []*Saiyan) []int {
	powers := make([]int, 0, len(saiyans))
	for _, saiyan := range saiyans {
		powers = append(powers, saiyan.Power)
	}
	return powers
}
