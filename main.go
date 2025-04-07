package main

import (
	"fmt"
	"time"
)

func main() {
  i := 3
  switch i {
  case 1:
    fmt.Println("One")
  case 2:
    fmt.Println("Two")
  case 3:
    fmt.Println("Three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's weekend!!!!!")
  default:
    fmt.Println("Go to work!")
  }

  // Switch without an expression works as if/else logic
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Before noon")
  default:
    fmt.Println("After noon")
  }

  // Type switch compares types instead of values
  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("Bool")
    case int:
      fmt.Println("Int")
    case int32:
      fmt.Println("Int32")
    default:
      fmt.Printf("Don't know type: %T\n", t)
    }
  }

  whatAmI(1)
  whatAmI(true)
  whatAmI('o')
  whatAmI("Hello")
}
