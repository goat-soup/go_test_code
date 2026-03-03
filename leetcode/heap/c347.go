package heap

import "container/heap"

type A struct {
	v int // value
	p int // priority
}

type NHeap []A

func (a NHeap) Len() int {
	return len(a)
}

func (a NHeap) Less(i, j int) bool {
	return a[i].p < a[j].p
}

func (a NHeap) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
	//
}

func (a *NHeap) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]
	return x
}

func (a *NHeap) Push(x interface{}) {
	*a = append(*a, x.(A))
}

func GetKNums(nums []int, k int) []int {
	a := &NHeap{}
	heap.Init(a)
	temp := make(map[int]int)
	for _, v := range nums {
		temp[v]++
	}
	for v, p := range temp {
		at := A{v: v, p: p}
		heap.Push(a, at)
		if a.Len() > k {
			heap.Pop(a)
		}
	}
	res := []int{}

	for a.Len() > 0 {
		res = append(res, (*a)[0].v)
		heap.Pop(a)
	}
	return res
}
