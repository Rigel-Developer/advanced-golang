package syncmutex

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, mu *sync.WaitGroup, lock *sync.Mutex) {
	defer mu.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance() int {
	b := balance
	return b

}

func SyncMutexDeposit(amount int) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println("Balance after depositing: ", Balance())
}
