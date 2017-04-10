package main

import "fmt"

func intSeq() func() int {

  i := 0
  return func() int {
    i++
    return i
  }
}


func main() {

  nextint := intSeq()
  fmt.Println(nextint())
  fmt.Println(nextint())
  fmt.Println(nextint())

  fmt.Println("------------------------")
  newint := intSeq()
  fmt.Println(newint())
  fmt.Println(newint())
}
