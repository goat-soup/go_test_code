package heap

import (
	"container/heap"
)

// 定义一个最小堆结构体，用来维护堆
type MinHeap []int

// 实现 heap.Interface 的方法

// Len 是堆中元素的个数
func (h MinHeap) Len() int { return len(h) }

// Less 比较堆中两个元素的大小，返回值 true 表示第一个元素小于第二个元素
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap 交换堆中两个元素的位置
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push 插入一个新的元素到堆中
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出堆顶元素
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindKthLargest(nums []int, k int) int {
	// 初始化一个空的最小堆
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for i := 0; i < len(nums); i++ {
		// 将当前元素加入堆
		heap.Push(minHeap, nums[i])

		// 当堆的大小超过 k 时，弹出堆顶元素
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	// 堆顶元素即为第 k 个最大元素
	return (*minHeap)[0]
}
