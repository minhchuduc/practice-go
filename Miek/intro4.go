package main

import (
	"fmt"
	"unicode/utf8"
)

const str = "asSASA ddd dsjkdsjs dk"

func main() {
	fmt.Println(str)
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		value, ok := m[str[i]]
		//fmt.Println(value, ok, byte(i))
		if ok == false {
			m[str[i]] = 1
		} else {
			m[str[i]] = value + 1
		}
	}
	for i := range m {
		fmt.Printf("Number of char '%v' is: %v\n", string(i), m[i])
	}

	fmt.Printf("Number of byte = %v\n", utf8.RuneCountInString(str))

	arr := []byte(str)
	fmt.Println("arr =", arr)
	var bonus string = "abc"
	bonus_arr := []byte(bonus)
	fmt.Println("bonus_arr =", bonus_arr)

	head := arr[:4]
	tail := arr[4:]

	var res []byte
	res = append(res, head...)
	res = append(res, bonus_arr...)
	res = append(res, tail...)
	fmt.Println(res)
	fmt.Println(string(res))

	fb := `foobar`
	rev := make([]byte, len(fb)+1)
	for i := range fb {
		rev[len(fb)-i] = fb[i]
	}
	fmt.Println(string(rev))

	// Way to Replace
	s := "ðå ぽ ai no Æl"
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", s)
	fmt.Printf("After (replace): %s\n", string(r))

}
