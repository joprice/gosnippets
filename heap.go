package main

import (
	"container/heap"
	"fmt"
)

// example from go heap docs

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func main() {
	h := &IntHeap{82, 12, 26}
	heap.Init(h)
	heap.Push(h, 10)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
