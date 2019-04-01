//GO Program to print full Pyramid using *
package main

import "fmt"

func main() {
	n := 5

	for i := 0; i < n; i++ {
		for k := n; k > i; k-- {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}

		fmt.Print("\n")
	}

}
