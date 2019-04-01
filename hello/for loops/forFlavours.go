package main

import "fmt"

func main() {

	/*i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	} //this is like a while loop in other programming laguages
	*/
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Println(no, i, no*i)
	}
	/*output10 1 10
	11 2 22
	12 3 36
	13 4 52
	14 5 70
	15 6 90
	16 7 112
	17 8 136
	18 9 162
	19 10 190*/
}
