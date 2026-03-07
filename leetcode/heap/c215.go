package heap

import (
	"container/heap"
)

/*
解题思路

使用最小堆：我们可以利用堆（heap）来优化这个过程。最小堆是一种完全二叉树，堆顶（根节点）存储的是当前堆中的最小元素。

堆的大小固定为 k：通过维护一个大小为 k 的最小堆，我们能够确保堆顶是数组中第 k 大的元素。因为：

在堆的元素数目小于 k 时，我们直接将元素加入堆。

当堆的元素数目达到 k 时，我们就开始比较堆顶的元素。如果当前遍历的元素大于堆顶元素，我们就替换堆顶元素。这样堆顶始终保留的是第 k 大的元素。

堆的操作：每次插入元素的时间复杂度是 O(log k)，因为堆的大小始终保持为 k。因此，我们遍历整个数组的时间复杂度是 O(n)，每次堆操作的时间复杂度是 O(log k)。由于堆的大小为 k，所以每次操作都只需要 O(log k) 的时间。

最后堆顶即为第 k 大的元素。
*/
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
