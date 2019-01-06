package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	for i:=0;i<k;i++{
		*h=append(*h,nums[i])
	}
	heap.Init(h)

	nums=nums[k:]
	for i:=0;i<len(nums);i++{
		top:=(*h)[0]
		if nums[i] > top{
			heap.Pop(h)
			heap.Push(h,nums[i])
		}
		//fmt.Println("heap:",*h)
	}
	return (*h)[0]
}
func main() {

	//nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	//nums := []int{7, 10,  8}
	nums := []int{7, 10, 7, 8, 11, 8}

	//nums := []int{6, 5, 4, 3, 2, 1}
	res := findKthLargest(nums, 1)
	fmt.Println(res)
}
