package main

import "fmt"

func main() {

	a := [...]int{2, 3, 5, 4}
	sum := 0
	for i, v := range a {

		fmt.Println(i, v)
		sum += v
	}
	fmt.Println("Total", sum)
}
