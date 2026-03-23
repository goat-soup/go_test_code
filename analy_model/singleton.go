package analymodel

import "sync"

type Singleton struct {
	data int
}

var instance *Singleton

var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: 42}
	})
	return instance
}
