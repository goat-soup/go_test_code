package pool

import (
	"fmt"
	"sync"
	"time"
)

// 连接池管理一组数据库连接或者网络连接，可复用

type Connection struct {
	ID int
}

type ConnectionPool struct {
	m        sync.Mutex
	pool     chan *Connection
	nextID   int
	maxAlive int
}

func NewConnectionPool(alive int) *ConnectionPool {
	return &ConnectionPool{
		pool:     make(chan *Connection, alive),
		maxAlive: alive,
	}
}

func (cp *ConnectionPool) Put(conn *Connection) {
	cp.pool <- conn
}

func (cp *ConnectionPool) Get() *Connection {
	cp.m.Lock()
	defer cp.m.Unlock()

	select {
	case conn := <-cp.pool:
		fmt.Println("Resuing connection: ", conn.ID)
		return conn
	default:
		if cp.nextID < cp.maxAlive {
			cp.nextID++
			conn := &Connection{ID: cp.nextID}
			fmt.Println("New connection: ", conn.ID)
			return conn
		}
		fmt.Println("Pool is full")
		conn := <-cp.pool
		return conn
	}
}

func ConnectionPoolTest() {
	cp := NewConnectionPool(3)
	for i := 0; i < 5; i++ {
		go func(id int) {
			conn := cp.Get()
			fmt.Printf("i am worker: %d,get connection %d\n", id, conn.ID)
			time.Sleep(2 * time.Second)
			cp.Put(conn)
		}(i)
	}
	time.Sleep(10 * time.Second)
}
