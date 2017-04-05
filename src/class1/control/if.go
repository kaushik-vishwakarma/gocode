package main

import (
 theFmt "fmt"
)

func TestIfA() {
  var i, j int = 0, 10
  if i < j {
    theFmt.Println("i less than j", i, j)
  } else if i == 10 {
    theFmt.Println("i equals j", i, j)
  } else {
    theFmt.Println("i greater j", i, j)
  }
}

func TestIfB() {
  var isTrue = false
  if isTrue {
    theFmt.Println("YAY!!!!", isTrue)
  } else {
    theFmt.Println("OH NO!!!!", isTrue)
  }
}


func main() {

  theFmt.Println("--------Test if A-------------")
  TestIfA()
  theFmt.Println()
  theFmt.Println("--------Test if B-------------")
  TestIfB()
  theFmt.Println()
}
