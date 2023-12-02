package syncmutex

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, mu *sync.WaitGroup, lock *sync.RWMutex) {
	defer mu.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b

}

func SyncMutexDeposit(amount int) {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 1; i <= 15; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println("Balance after depositing: ", Balance(&lock))
}
