package main

import "fmt"

const limit = 3

type Stack []int

var stack = make(Stack, limit, limit)
var index int // 'index' point to the next available

func (stack Stack) String() string {
	var ret string
	for i, v := range stack {
		ret = ret + fmt.Sprintf("[%v:%c] ", i, rune(v))
	}
	return ret
}

func main() {
	fmt.Println("Initial stack =", stack)
	push(105)
	push(106)
	push(107)
	push(108)
	push(109)
	fmt.Printf("My stack = %v\n", stack)
	pop()
	pop()
	pop()
	pop()

}

func push(k int) bool {
	if index < limit {
		stack[index] = k
		index++
		fmt.Println(stack)
		return true
	} else {
		fmt.Println("Stack Overflow")
		return false
	}
}

func pop() (int, bool) {
	if index > 0 {
		err := false
		index--
		fmt.Println("Pop", stack[index])
		return stack[index], err
	} else {
		err := true
		fmt.Println("Stack is empty!")
		return 0, err
	}
}
