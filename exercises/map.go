package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	substrings := strings.Fields(s)
	fmt.Println(substrings)
	ret := make(map[string]int)
	for _, substr := range substrings {
		ret[substr] = ret[substr] + 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
