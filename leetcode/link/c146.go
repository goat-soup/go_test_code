package link

import "fmt"

type Node struct {
	k    int
	v    int
	pre  *Node
	next *Node
}
type LRUCache struct {
	capacity int
	// map[k]Node
	cache map[int]*Node
	head  *Node
	tail  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	head.pre = nil
	tail.next = nil
	return LRUCache{capacity: capacity, cache: make(map[int]*Node), head: head, tail: tail}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}
	// break
	n.pre.next = n.next
	n.next.pre = n.pre

	// insert
	n.next = this.head.next
	n.pre = this.head
	this.head.next.pre = n
	this.head.next = n
	return this.cache[key].v
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		// no exist, insert
		t := &Node{k: key, v: value}
		this.addToHead(t)
		this.cache[key] = t
		// full cap, clean the last
		if len(this.cache) > this.capacity {
			r := this.tail.pre
			this.remove(r)
			delete(this.cache, r.k)
		}
	} else {
		//exist, update
		this.cache[key].v = value
		this.remove(this.cache[key])
		// insert
		this.addToHead(this.cache[key])
	}

}

func (this *LRUCache) remove(n *Node) {
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (this *LRUCache) addToHead(n *Node) {
	n.next = this.head.next
	n.pre = this.head
	this.head.next.pre = n
	this.head.next = n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRU() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))

	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
