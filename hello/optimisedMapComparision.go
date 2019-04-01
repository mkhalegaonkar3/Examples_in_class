package main

import "fmt"

func main() {

	m1 := map[string]int{
		"india": 001,
	}
	m2 := map[string]int{
		"pak": 234,
	}

	compareMaps(m1, m2)
}
func compareMaps(m1 map[string]int, m2 map[string]int) {

	val1, ok1 := m1["india"]
	val2, _ := m2["pak"]
	if ok1 == true {
		if val1 == val2 {
			fmt.Println("values are same")
		} else {
			fmt.Println("not present")
		}

	}
}
