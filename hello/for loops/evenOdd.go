package main

import "fmt"

func main() {
	//n:=10

	/*for i:=0;i<n;i++{

	if i%2==0{
		fmt.Println(i,"even")
		}else{
		fmt.Println(i,"odd")
		}
	}*/
	//same line if consition and initialisation
	if num := 0; num == 0 {
		fmt.Println(num, "ZERO")
	} else if num%2 == 0 {
		fmt.Println(num, "even")
	} else {
		fmt.Println(num, "odd")
	}

}
