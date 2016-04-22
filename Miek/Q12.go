package main

import "fmt"

func main() {
	fmt.Println(mapx(double, []int{1, 2, 3}))

	// Must use []e, can not use []string
	fmt.Println(mapArbitrary(double2, []e{"asdf", "uop"}))
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

type e interface{}

func double2(j e) e {
	switch j.(type) {
	case int:
		return j.(int) * 2
	case string:
		return j.(string) + j.(string)
	case float64:
		return j.(float64) * 2
	}
	return j // if don't match "int" / "string" / "float64"
}

func mapArbitrary(f func(e) e, ar []e) []e {
	ret := make([]e, len(ar))
	for i, v := range ar {
		ret[i] = f(v)
	}
	return ret
}
