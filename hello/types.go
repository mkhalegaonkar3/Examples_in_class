package main

import "fmt"

func main() {

	c := 5 + 2i
	f := 23.34
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	i := 55

	x := float64(i) + f
	fmt.Println(x)
}
