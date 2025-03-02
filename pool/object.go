package pool

import (
	"fmt"
	"sync"
)

type ExpensiveObject struct {
	ID int
}

type ExpensiveObjectPool struct {
	m        sync.Mutex
	pool     chan *ExpensiveObject
	maxAlive int
}

func (eop *ExpensiveObjectPool) Put(conn *ExpensiveObject) {
	eop.pool <- conn
}

func (eop *ExpensiveObjectPool) Get() *ExpensiveObject {
	select {
	case eo := <-eop.pool:
		fmt.Println("Resuing Pool :", eo.ID)
		return eo
	default:
		eop.m.Lock()
		defer eop.m.Unlock()
		eop.maxAlive++
		eo := &ExpensiveObject{ID: eop.maxAlive}
		fmt.Println("New object:", eo.ID)
		return eo
	}

}
func NewObjectPool(maxAlive int) *ExpensiveObjectPool {
	return &ExpensiveObjectPool{
		// maxAlive: maxAlive,
		pool: make(chan *ExpensiveObject, maxAlive),
	}
}

func ObjectPoolTest() {
	eop := NewObjectPool(3)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			eo := eop.Get()
			fmt.Printf("i am %d, i use %d\n", id, eo.ID)
			eop.Put(eo)
		}(i)
	}
	wg.Wait()
}
