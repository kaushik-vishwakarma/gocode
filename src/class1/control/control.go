package main

import (
 theFmt "fmt"
)

func TestA() {
  var i, j int = 0, 10
  if i < j {
    theFmt.Println("i less than j", i, j)
  } else if i == 10 {
    theFmt.Println("i equals j", i, j)
  } else {
    theFmt.Println("i greater j", i, j)
  }
}

func TestA2() {
  i, j :=  11, 10
  switch {
    case i < j:
      theFmt.Println("i less than j", i, j)
    case i == j:
      theFmt.Println("i equals j", i, j)
   default :
    theFmt.Println("i greater j", i, j)
  }
}

func TestB() {
  var isTrue = false
  if isTrue {
    theFmt.Println("YAY!!!!", isTrue)
  } else {
    theFmt.Println("OH NO!!!!", isTrue)
  }
}

func main() {

  theFmt.Println("--------Test A-------------")
  TestA()
  theFmt.Println()
  theFmt.Println("--------Test B-------------")
  TestB()
  theFmt.Println()
  theFmt.Println("--------Test A2-------------")
  TestA2()

    
}
