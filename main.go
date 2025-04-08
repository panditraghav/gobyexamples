package main

import "fmt"

func sums(nums ...int) {
	fmt.Printf("Type of arg: %T\n", nums)
	fmt.Println("nums: ", nums)
	var total int = 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total: ", total)
}

func main() {
	sums(1, 2)
	sums(4, 5, 6, 7)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	nums := []int{8, 9, 10, 11}
	sums(nums...)
}
