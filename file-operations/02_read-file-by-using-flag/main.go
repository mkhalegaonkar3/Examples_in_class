package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	fptr := flag.String("fpath", "test1.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)

	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("Read file Error", err)
	}
	fmt.Println("File content is :", string(data))
}
