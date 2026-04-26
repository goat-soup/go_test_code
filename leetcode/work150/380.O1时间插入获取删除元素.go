package work150

import "math/rand"

type RandomizedSet struct {
	nums []int
	idx  map[int]int // val -> index
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		idx:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.idx[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.idx[val]
	if !ok {
		return false
	}
	// 将最后一个元素记录到要删除的位置
	lastVal := this.nums[len(this.nums)-1]
	this.nums[index] = lastVal
	this.idx[lastVal] = index
	// 清理
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.idx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
