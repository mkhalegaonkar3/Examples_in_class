package main

import "fmt"

func main() {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("its a vowel")

	default:
		fmt.Println("not a vowel")
	}
}
