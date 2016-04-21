package main

import "fmt"

func main() {
	res := f()
	fmt.Println(res)

}

func f() (ret int) {
	ret = 3
	arg := 8
	defer func(x int) {
		ret += arg
	}(arg)
	return
}
