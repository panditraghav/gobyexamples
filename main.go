package main

import "fmt"

func main() {
  i := 1
  // Only condition
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  // Standard loop
  for j := 0; j < 3; j++ {
    fmt.Println(j)
  }

  // 0 1 2
  for i := range 3 {
    fmt.Println("range", i)
  }

  // Infinite loop
  for {
    fmt.Println("loop")
    break
  }

  for n := range 6 {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}
