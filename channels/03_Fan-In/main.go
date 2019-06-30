package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)
	go rcv(even, odd, fanIn)
	for v := range fanIn {
		fmt.Println(v)
	}
	fmt.Println("about to exit")

}
func send(e chan int, o chan int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}
func rcv(e chan int, o chan int, f chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			f <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			f <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(f)
}
