package main

import "fmt"

func main() {
	fmt.Println(mapx(double, []int{1, 2, 3}))
}

// TODO: improve it to work on arbitrary function/array
func mapx(f func(i int) int, ar []int) (ret []int) {
	ret = make([]int, len(ar))
	for i, v := range ar {
		ret[i] = f(v)
	}
	return ret
}

func double(i int) int {
	return i + i
}
