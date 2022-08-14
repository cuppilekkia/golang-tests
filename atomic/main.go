package main

import (
	"fmt"
	"sync"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *AtomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *AtomicInt) Value() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func main() {
	counter := AtomicInt{}
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			counter.Increase()
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	fmt.Println("Counter: ", counter.Value())
}
