package main

import (
	"fmt"
	"sync"
	"time"
)

// RWMutex provide a shared lock RMutex and a exclusive lock Mutext
// bank3 add withdraw func that has the same effect of ex9_1, but using mutext
// this package indicate an account
// the deposits, and balance channel holds the send, receive communications to that account, respectively
// all read, write is confined to the monitor func

var account int

var mu sync.RWMutex

func Deposit(money int) {
	mu.Lock()
	account += money
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(1 * time.Second)
	return account
}

func Withdraw(money int) bool {
	mu.Lock()
	defer mu.Unlock()
	if money <= account {
		account -= money
		return true
	}
	return false
}

var wg = sync.WaitGroup{}

func main() {
	// simple tests
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			fmt.Printf("%s time %s\n", Balance(), time.Now())
			wg.Done()
		}()
	}
	wg.Wait()
}
