package main

import (
	"fmt"
	"math"
	"strconv"
)

type I interface {
	M()
	F1(tail string)
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}
func (t T) F1(tail string) {
	t.S = t.S + tail
	fmt.Println("now = ", t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func (f F) F1(tail string) {
	//tmp_f := float64(f)
	//f_str := strconv.FormatFloat(tmp_f, 'E', -1, 32)
	f_str := f.String() + tail
	fmt.Println("now = ", f_str)
}
func (f F) String() string { // "Stringer" interface defined by fmt package
	f_str := strconv.FormatFloat(float64(f), 'E', -1, 32)
	return fmt.Sprintf("%v", f_str)
}

func describe(i I) {
	fmt.Printf("Value = %v, Type = %T \n", i, i)
}
func main() {
	var i I
	i = T{"hello you"}
	i.F1(" ZZZ")
	describe(i)
	//i.M()
	fmt.Println("----------------------------------")
	i = F(math.Pi)
	i.F1(" ZZZ")
	describe(i)
	//i.M()

}
