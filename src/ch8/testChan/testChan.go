package main

import "fmt"

func main() {
	case0()
	//case1()
	//case2()
}

// this case test range list allow modify the list or not
// the result is not, list is just a 'capture'
func case0() {
	list := []int{0, 1, 2, 3}
	for _, i := range list {
		fmt.Println(i)
		list = append(list, i)
	}
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
