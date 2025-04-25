package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var (
	ErrOutOfTea = fmt.Errorf("no more tea available")
	ErrPower    = fmt.Errorf("can't boil water")
)

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		/*
			We can wrap errors with higher-level errors to add context.
			The simplest way to do this is with the %w verb in fmt.Errorf.
			Wrapped errors create a logical chain (A wraps B, which wraps C, etc.)
			that can be queried with functions like errors.Is and errors.As.
		*/
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{32, 42} {
		if val, err := f(i); err != nil {
			fmt.Printf("Error: %v, return: %v\n", err, val)
		} else {
			fmt.Println("No error return: ", val)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			fmt.Printf("Error: %v\n", err)
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Println("Unknown error!")
			}
			continue
		}
		fmt.Printf("Tea is ready! (i = %v)\n", i)
	}
}
