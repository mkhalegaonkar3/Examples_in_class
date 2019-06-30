package main

import (
	"fmt"
	"os"
)

func main() {
	//f, err := os.Create("lines.txt")
	f, err := os.OpenFile("lines.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	//var l int
	//l, err = f.WriteString("newLine")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	//fmt.Println(l)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
