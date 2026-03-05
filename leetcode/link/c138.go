package link

type Node1 struct {
	Val    int
	Next   *Node1
	Random *Node1
}

func CopyRandomList(head *Node1) *Node1 {
	if head == nil {
		return nil
	}
	m := make(map[*Node1]*Node1)
	pHead := head
	for pHead != nil {
		m[pHead] = &Node1{Val: pHead.Val}
		pHead = pHead.Next
	}
	pHead = head
	for pHead != nil {
		m[pHead].Next = m[pHead.Next]
		m[pHead].Random = m[pHead.Random]
		pHead = pHead.Next
	}
	return m[head]
}
