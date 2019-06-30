package main

import (
	"fmt"
)

func hello1(done1 chan bool) {
	fmt.Println("Hello world goroutine1")

	done1 <- false
	fmt.Println("after done1")

}

func hello2(done chan bool) {
	fmt.Println("Hello world goroutine2")
	done <- true
	fmt.Println("after done2")

}
func hello3(done1 chan bool) {
	fmt.Println("Hello world goroutine3")
	done1 <- false
	fmt.Println("after done3")

}
func hello4(done chan bool) {
	fmt.Println("Hello world goroutine4")
	done <- true
	fmt.Println("after done4")

}

func main() {
	done := make(chan bool)
	done1 := make(chan bool)

	go hello1(done)
	go hello2(done1)
	go hello3(done)
	go hello4(done1)
	fmt.Println(<-done1)
	fmt.Println(<-done)

	fmt.Println("main function")

}
