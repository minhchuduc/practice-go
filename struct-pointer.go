package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v1 := Vertex{X: 1, Y: 2}
	v2 := Vertex{3, 4}
	fmt.Println(v1, v2)

	p := &v1
	p.X = 1e5
	fmt.Println(v1)

	// "&" operator generates a pointer to its operand (in this case, operand = Vertex{6, 7} )
	q := &Vertex{6, 7} // q has type *Vertex - mean a pointer to value that have Vertex type
	t := &Vertex{6, 7}
	q.X = 9
	fmt.Println(q, t)
}
