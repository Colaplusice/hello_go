package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	a := ListNode{}
	a.Next = nil
	fmt.Println(&a.Next)
}
