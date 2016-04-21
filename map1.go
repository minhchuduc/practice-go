package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(power, exists)
	delete(lookup, "goku2") // can delete a non-exist key
	fmt.Println(lookup)

	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = &Saiyan{"Krillin monkey", nil} //todo load or create Krillin
	fmt.Println(goku)

	lookup2 := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}
	//array2 := array[]
	fmt.Println(lookup2)
}

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}
