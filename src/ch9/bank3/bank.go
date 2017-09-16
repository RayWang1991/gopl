package main

import (
	"fmt"
	"sync"
)

// bank3 add withdraw func that has the same effect of ex9_1, but using mutext
// this package indicate an account
// the deposits, and balance channel holds the send, receive communications to that account, respectively
// all read, write is confined to the monitor func

var account int

var mu sync.Mutex

func Deposit(money int) {
	mu.Lock()
	account += money
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
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

func main() {
	// simple tests
	fmt.Println(Balance()) // should be 0
	Deposit(10)
	fmt.Println(Balance()) // should be 10
	Deposit(20)
	fmt.Println(Balance())                         // should be 30
	fmt.Printf("%t,%d\n", Withdraw(20), Balance()) // should be true, 10
	fmt.Printf("%t,%d\n", Withdraw(20), Balance()) // should be false, 10
}
