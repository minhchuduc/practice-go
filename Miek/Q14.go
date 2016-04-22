/* Buble sort */
package main

import "fmt"

func BubbleSort(ar []int) []int {
	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			//fmt.Println(i, " ", j)
			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
	return ar
}

func main() {
	input := []int{3, 1, 4, 2, 5}
	fmt.Println(BubbleSort(input))
}
