package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("Empty array: ", a)

	a[3] = 2
	fmt.Println("Set array: ", a)
	fmt.Println("Item: ", a[3])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b: ", b)

	// You can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}

	b = [...]int{1, 3: 2, 5}
	fmt.Println("b: ", b)

	var twod [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twod[i][j] = 2*i + 3*j
		}
	}
	fmt.Println(twod)

	twod = [2][3]int{
		{1, 2, 3},
		{4, 7, 8},
	}
	fmt.Println(twod)
}
