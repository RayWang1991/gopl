package main

import "fmt"

// ex9_1 adds func withdraw(amount int)bool to that account
// this package indicate an account
// the deposits, and balance channel holds the send, receive communications to that account, respectively
// all read, write is confined to the monitor func

type withdrawReq struct {
	money  int
	result chan<- bool
}

var deposits = make(chan int)
var withdraw = make(chan withdrawReq)
var balances = make(chan int)

func Deposit(money int) { deposits <- money }

func Balance() int { return <-balances }

func Withdraw(money int) bool {
	res := make(chan bool)
	withdraw <- withdrawReq{money, res}
	return <-res
}

func monitor() {
	account := 0
	for {
		select {
		case money := <-deposits:
			account += money
		case balances <- account:
		case wr := <-withdraw:
			var res bool
			if wr.money > account {
				res = false
			} else {
				account -= wr.money
				res = true
			}
			go func() {
				wr.result <- res
			}()
		}
	}
}

func main() {
	go monitor()
	// simple tests
	fmt.Println(Balance()) // should be 0
	Deposit(10)
	fmt.Println(Balance()) // should be 10
	Deposit(20)
	fmt.Println(Balance())                         // should be 30
	fmt.Printf("%t,%d\n", Withdraw(20), Balance()) // should be true, 10
	fmt.Printf("%t,%d\n", Withdraw(20), Balance()) // should be false, 10
}
