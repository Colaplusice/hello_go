package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	fake_node := ListNode{Val: 2, Next: head}
	for fake_node.Next != nil {
		if fake_node.Next.Val == val {
			fake_node.Next = fake_node.Next.Next
		}
		fake_node.Next = fake_node.Next.Next
	}
	return head

}

func main() {

	a := ListNode{Val: 1}
	return_val := removeElements(a, 1)

	print(a)

}
