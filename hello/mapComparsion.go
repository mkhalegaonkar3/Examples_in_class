package main

import "fmt"

func main() {

	map1 := map[string]int{
		"india": 007,
	}
	map2 := map[string]int{
		"asd": 007,
	}
	fmt.Println(map1, map2)

	compareMap(map1, map2)
}

func compareMap(m1 map[string]int, m2 map[string]int) {

	var (
		key1,
		key2 string
		val1,
		val2 int
	)

	for key1, val1 = range m1 {

		for key2, val2 = range m2 {
			if key1 == key2 && val1 == val2 {
				fmt.Println("maps are same")
			} else {
				fmt.Println("maps are different")
			}

		}

	}

}
