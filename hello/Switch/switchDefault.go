package main

import "fmt"

func main() {

	/*switch num := number(); {
		case num <50:
			fmt.Println("FALSE")

		case num < 100:
			fmt.Println("TRUE")
			fallthrough
		default:
			fmt.Println("DEFAULT")
		}
	}*/
	switch {
	case number() < 50:
		fmt.Println("FALSE")

	case number() < 100:
		fmt.Println("TRUE")
		fallthrough
	default:
		fmt.Println("DEFAULT")
	}
}
func number() int {
	num := 15 * 5
	return num
}
