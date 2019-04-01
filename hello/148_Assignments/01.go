//Program in Go language to Program to Add Two Matrix Using Multi-dimensional Arrays
package main

import "fmt"

func main() {

	var m1 = [2][2]int{}
	var m2 = [2][2]int{}
	m1[0][0] = 12
	m1[0][1] = 13
	m1[1][0] = 41
	m1[1][1] = 42

	m2[0][0] = 22
	m2[0][1] = 23
	m2[1][0] = 31
	m2[1][1] = 32

	for i, v := range m1 {

		fmt.Println(i, v)

	}

	for i1, v1 := range m2 {
		fmt.Println(i1, v1)

	}
}
