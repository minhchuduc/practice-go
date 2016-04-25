package main

import "fmt"

func add(a interface{}, b interface{}) interface{} {
	//return a.(int) + b.(int)
	if ta, ok := a.(int); ok {
		if tb, ok := b.(int); ok {
			return ta + tb
		}
	} else {
		fmt.Println("Parameter's type is not supported!")
		return a
	}
	return nil
	// If don't have above line, Go will error. Don't know why.
	// Maybe because of "interface{}" return type.
}

func g(something interface{}) int {
	return something.(I).Get()

}

type I interface {
	Get() int
	//Put(int)
}

type S struct{ i int }

func (p *S) Get() int { return p.i }

//func (p *S) Put(v int) { p.i = v }

func main() {
	fmt.Println(add(5, 7))
	fmt.Println(add(6.6, 8.8))
	//fmt.Println(add(6, 7.7))
	stra := "the spice must flow"
	byts := []byte(stra)
	strb := string(byts)
	fmt.Println(stra, byts, strb)

	s := new(S)
	fmt.Println(g(s))

	i := 5
	fmt.Println(g(i)) // Compile-time is ok, but Run-time will fail!
}
