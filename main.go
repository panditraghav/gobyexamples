package main

import (
  "fmt"
)

func main() {
  var a int = 10
  fmt.Println("a: ", a)
  fmt.Printf("Typeof a: %T\n", a)

  var b, c int = 10, 20
  fmt.Println("b: ", b, ", c: ", c)
  fmt.Printf("Typeof b: %T\n", b)

  var d int
  fmt.Println("d: ", d)

  f := "apple"
  fmt.Println(f)
}
