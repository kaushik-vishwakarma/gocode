package main

import (
 "fmt"
)

func TestForA () {
  for i:=0; i<10; i++ {
    fmt.Print("i=", i, ", ")
  }
  fmt.Println()
}

func TestForB () {
  i, n := 0, 10
  for i < n {
    fmt.Print("i=", i, ", ")
    i++
  }
  fmt.Println()
}

func TestForC () {
  i, n := 0, 10
  for  {
    if i >= n {
      break
    }
    fmt.Print("i=", i, ", ")
    i++
  }
  fmt.Println()
}


func main() {

  fmt.Println("--------Test for A-------------")
  TestForA()
  fmt.Println()
  fmt.Println("--------Test for B-------------")
  TestForB()
  fmt.Println()
  fmt.Println("--------Test for C-------------")
  TestForC()
  fmt.Println()
}
