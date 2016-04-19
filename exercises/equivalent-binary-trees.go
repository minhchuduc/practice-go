package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
)

/* tree.Tree struct
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t != nil {
		ch <- t.Value
		Walk(t.Left, ch)
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	res1 := StoreAndPrint(ch1, 10)
	fmt.Println("--------------------------------------")

	ch2 := make(chan int)
	go Walk(t2, ch2)
	res2 := StoreAndPrint(ch2, 10)

	sort.Ints(res1)
	sort.Ints(res2)
	same := true
	for i, _ := range res1 {
		if res1[i] != res2[i] {
			same = false
			fmt.Println("Checking element ", i)
			break
		}
	}
	return same
}

// just Print
func PrintChannel(ch chan int, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(<-ch, " ")
	}
}

func StoreAndPrint(ch chan int, total int) []int {
	res := make([]int, total)
	for i := 0; i < total; i++ {
		res[i] = <-ch
		fmt.Print(res[i], " ")
	}
	fmt.Println("")
	return res
}
func main() {
	t1 := tree.New(4)
	t2 := tree.New(5)

	fmt.Println(t1)
	fmt.Println(t2)
	/*
		ch1 := make(chan int)
		go Walk(t1, ch1)
		StoreAndPrint(ch1, 10)
		fmt.Println("--------------------------------------")

		ch2 := make(chan int)
		go Walk(t2, ch2)
		StoreAndPrint(ch2, 10)
	*/

	if Same(t1, t2) == true {
		fmt.Println("t1, t2 is the same.")
	} else {
		fmt.Println("t1, t2 is NOT the same!")
	}

}
