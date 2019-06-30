package main

import "fmt"

func main() {

	ch := make(chan int)
	go foo(ch)
	//bar(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
func foo(c chan<- int) {

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
	//if we are not closing channel then will get "fatal error: all goroutines are asleep - deadlock!",
	//bcz for.. range loop will continue(read value) upto closing of channel
}
