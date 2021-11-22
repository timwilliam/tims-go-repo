package main

import (
	"container/list"
	"fmt"
)

// package container/list is go implementation of double linked list data structure

func main() {
	data := list.New() // create a new double linked list

	data.PushFront("tim") // insert from the front
	// data.PushBack("tim") // insert from the back
	data.PushBack("william")
	data.PushBack("jr")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Back().Value)
	fmt.Println()

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println()

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
