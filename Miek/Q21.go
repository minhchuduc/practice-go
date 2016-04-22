package main

import (
	"container/list"
	"errors"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}

	fmt.Println("Now test my List implementation")
	ml := new(MyList)
	ml.Push(1)
	ml.Push(2)
	ml.Push(4)

	for g := ml.Front(); g != nil; g = g.Next() {
		fmt.Printf("%v\n", g.Value)
	}
	fmt.Println("===============")

	for h, err := ml.Pop(); err == nil; h, err = ml.Pop() {
		fmt.Printf("%v\n", h)
	}
}

/* Self-implement double-linked list from here */
type Value int

type Node struct {
	Value
	prev, next *Node
}

type MyList struct {
	head, tail *Node
}

func (l *MyList) Front() *Node {
	return l.head
}

func (l *Node) Next() *Node { //NOTICE: it's NOT symmetric with Front()
	return l.next
}

func (l *MyList) Push(v Value) *MyList {
	n := &Node{Value: v} // create new Node with Value = v
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n

	return l
}

var errEmpty = errors.New("MyList is empty.")

func (l *MyList) Pop() (v Value, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.Value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}

	return v, err // "return" also ok because of named-result
}
