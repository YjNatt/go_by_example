package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declared and initialized:", b)

	var nestedArr [2][3]int
	for index1 := 0; index1 < 2; index1++ {
		for index2 := 0; index2 < 3; index2++ {
			nestedArr[index1][index2] = index1 + index2
		}
	}
	fmt.Println("nested array: ", nestedArr)
}