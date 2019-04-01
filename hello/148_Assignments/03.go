//GO Program to Generate Fibonacci Sequence Up to a Certain Number
package main

import "fmt"

func main() {
	n := 3
	f1 := 1
	f2 := 1

	var f3 int
	//fmt.Println(i)

	fmt.Print(f1)
	fmt.Print("\t", f2)
	for i := 0; i < n; i++ {

		f3 = f1 + f2
		fmt.Print("\t", f3)

		f1 = f2
		f2 = f3
	}
	//fmt.Print("\t",f3)
}
