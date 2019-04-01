package main

import "fmt"

func main() {
	arr1 := [3]int{44}
	//OR
	arr2 := [3]int{12, 22, 21}

	//OR

	c := [...]int{12, 34, 21}
	//c:=[]int{12,34,21}

	//[5]int is distinct than [25]int
	// we can called it as int of size 5 & int of size 25

	fmt.Println(c)

	fmt.Println(arr1)
	fmt.Println(arr2)

	arr2 = arr1
	fmt.Println(arr2)

	arr2[2] = 50
	fmt.Println(arr2)
	fmt.Println(arr1)

	num := [...]int{12, 23, 12, 32, 23}
	fmt.Println(num)
	ar := x(num)
	fmt.Println(ar)
	fmt.Println(num)

}

func x(num [5]int) [5]int {
	num[0] = 55
	fmt.Println(len(num))
	return num
}
