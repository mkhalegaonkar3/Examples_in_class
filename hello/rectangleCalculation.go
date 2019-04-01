package main

import "fmt"

func reactCal(l, b int) (area int, peri float64) {
	area = l * b
	peri = float64(2 * (l + b))
	return area, peri
	//return l*b,float64(2*(l+b))
}

/*OR
func reactCal(l,b int)(int,float64){
	var area =l*b
	var peri =float64(2*(l+b))
	return area,peri
	//return l*b,float64(2*(l+b))
}

OR
func reactCal(l,b int)(int,float64){

	return l*b,float64(2*(l+b))
}


*/

func main() {

	length := 20
	breadth := 5

	area, per := reactCal(length, breadth)

	fmt.Println("Area=", area)
	fmt.Println("Permieter=", per)
	fmt.Printf("%T", per)
}
