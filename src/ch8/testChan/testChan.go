package main

import "fmt"

func main() {
	case1()
	case2()
}

func case2() {
	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Println("send1")
		ch <- 2
		fmt.Println("send2")
		ch <- 3
		fmt.Println("send3")
		ch <- 4
		fmt.Println("send4")
	}()
	<-ch
	<-ch
	<-ch
	<-ch
}

func case1() {
	ch := make(chan int, 8) // buffered channel with capacity 8
	for i := 0; i < 5; i++ {
		ch <- i
	}
	for i := 0; i < 3; i ++ {
		x := <-ch
		fmt.Println(x)
	}
	close(ch)
	// sending values to closed channel cause panic
	// ch <- 9348
	// receiving values from closed channel is allowed
	// when there are no values left, empty values of that type will be receiving
	for i := 0; i < 3; i ++ {
		x := <-ch
		fmt.Println(x)
	}
}
