package main

import (
	"fmt"

	"github.com/mkhalegaonkar3/Examples_in_class/structs/computer"
)

func main() {
	var spec computer.Spec

	spec.Maker = " Apple"

	spec.Price = 1000

	fmt.Println("Spec:=", spec)

}

/* OUTPUT

$ go run main.go
Spec:= { Apple  1000}

*/
