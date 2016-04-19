package main

import "fmt"

func main() {
	var scores [10]int
	scores[7] = 777
	fmt.Println(scores)

	scores2 := [4]int{9001, 9333, 212, 33}
	for index, value := range scores2 {
		fmt.Println(index, value)
	}

	scores3 := []int{}
	scores3 = scores2[:]
	scores3 = append(scores3, 66)
	fmt.Println(scores3)
	fmt.Printf("%p\n", &scores3)

	test := make([]int, 10)
	fmt.Println(test)
}
