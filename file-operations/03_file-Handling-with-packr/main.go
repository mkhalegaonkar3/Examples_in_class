package main

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("/data")
	data := box.String("demo.txt")
	data1 := box.String("demo1.txt")

	fmt.Println("Contents of file:", data)
	fmt.Println("Contents of file1:", data1)

}
