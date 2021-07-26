package main

import "fmt"

func main() {
	count := 1
	for count <= 3 {
		fmt.Println(count)
		count = count + 1
	}

	for count2 := 7; count2 <= 9; count2++ {
		fmt.Println(count2)
	}

	for {
		fmt.Println("loop")
		break
	}

	for number := 0; number <= 5; number++ {
		if number % 2 == 0 {
			continue
		}
		fmt.Println(number)
	}
}