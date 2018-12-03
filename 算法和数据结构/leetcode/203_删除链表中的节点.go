package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func OutList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Print(temp.Val, " ")
		temp = temp.Next
	}
	fmt.Println()

}
func removeElements(head *ListNode, val int) *ListNode {
	for ; head != nil && head.Val == val;
	head = head.Next {
	}
	previous := head
	for previous != nil && previous.Next != nil {
		if previous.Next.Val == val {
			previous.Next = previous.Next.Next
		} else {
			previous = previous.Next
		}

	}
	return head
}

func main() {

	a := ListNode{Val: 1}
	b := ListNode{Val: 2}
	c := ListNode{Val: 2}
	d := ListNode{Val: 1}
	e := ListNode{Val: 5}
	f := ListNode{Val: 6}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = &f
	OutList(&a)
	return_val := removeElements(&a, 2)
	OutList(return_val)

}
