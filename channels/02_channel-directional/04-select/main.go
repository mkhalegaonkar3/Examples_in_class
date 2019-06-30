package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quite := make(chan int)

	//send
	go send(even, odd, quite)
	rcv(even, odd, quite)
	fmt.Println("About to exit")
}
func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)

	q <- 0
	close(q)
	//if u r not closing channel  also fine when u r using select stmt
}
func rcv(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From eve", v)
		case v := <-o:
			fmt.Println("From odd", v)
		case v := <-q:
			fmt.Println("From quite", v)
			return
		}
	}
}
