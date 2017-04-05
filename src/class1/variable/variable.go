package main

import (
  "fmt"
)

func main() {
  var i,j int = 10, 11
  k := 10
  p := new(int)
  *p = 101;


  fmt.Printf("%T, %v\n", i, i)
  fmt.Printf("%T, %v\n", j, j)
  fmt.Printf("%T, %v\n", k, k)
  fmt.Printf("%T, %v\n", p, p)
  fmt.Printf("%T, %v\n", *p, *p)

    
}
