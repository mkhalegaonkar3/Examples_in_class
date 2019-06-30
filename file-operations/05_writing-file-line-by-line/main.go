package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("lines")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
	s := "Hi i am Mayuresh"
	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Fprintln(f, s)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
