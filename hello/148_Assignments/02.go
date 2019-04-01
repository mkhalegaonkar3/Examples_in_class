//GO Program to Check Whether a Number is Even or Odd
package main

import "fmt"

func main() {

	intList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := [10]int{}
	odd := [10]int{}
	fmt.Println(even, odd)
	for i, v := range intList {

		if v%2 == 0 {

			//fmt.Printf("%d is Even\n",v)
			even[i] = v
		} else {
			//fmt.Printf("%d is Odd\n",v)
			odd[i] = v
		}

	}
	fmt.Println("Odd slice = ", odd, "\nEven slice = ", even)
}
