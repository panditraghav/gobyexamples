package main

import (
	"fmt"
	"slices"
)

// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.
func main() {
	var s []string

	fmt.Println("Un initialized: ", s, ", len: ", len(s), ", is_nil: ", s == nil)

	// type, length, capacity
	s = make([]string, 3, 8)
	fmt.Println("val: ", s, ", len: ", len(s), ", cap: ", cap(s))

	s[0] = "Hello"
	s[1] = "How"
	s[2] = "Are"

	fmt.Println("s:", s)
	s = append(s, "You")
	fmt.Println("s:", s)

	// Copy slice
	var c []string = make([]string, len(s))
	// Or like this:
	// c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)

	l := s[1:3]
	fmt.Println("l:", l)
	fmt.Println("s[2:]:", s[2:])
	fmt.Println("s[:3]:", s[:3])

	// We can declare and initialize a variable for slice in a single line as well.
	t1 := []string{"g", "h", "i"}
	fmt.Println("t: ", t1)

	t2 := []string{"g", "h", "i"}

	if slices.Equal(t1, t2) {
		fmt.Printf("%v is equal to %v\n", t1, t2)
	}

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
}
