package main

import "fmt"

//Definition for singly-linked list.
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var temp = &ListNode{}
	for head.Next != nil {
		fmt.Println(head.Next.Val)
		temp = head.Next
		head.Next = head.Next.Next
		temp.Next = head
		head = head.Next
	}
	return temp
}
func print_node(head *ListNode) {
	if head == nil {
		fmt.Println("head is nil")
		return
	}
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func is_Palindrome(head *ListNode) bool {
	one := head
	two := head
	for two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
	}
	reverse_one := reverse(one)
	print_node(reverse_one)
	//for one != nil {
	//	if one.Val != reverse_one.Val {
	//		return false
	//	}
	//	one = one.Next
	//	reverse_one = reverse_one.Next
	//}
	return true
}
func main() {
	a := ListNode{1, nil}
	b := ListNode{2, &a}
	c := ListNode{3, &b}
	is_Palindrome(&c)

}
