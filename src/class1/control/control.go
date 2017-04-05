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

func TestC() {
  var t interface{}
  i := 10
  t = &i
 // t = 10 
 // t = true
 // t = "kaushik"
  switch t := t.(type) {
    default:
      theFmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
    case bool:
      theFmt.Printf("boolean %t\n", t)             // t has type bool
    case int:
      theFmt.Printf("integer %d\n", t)             // t has type int
    case *bool:
      theFmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
    case *int:
      theFmt.Printf("pointer to integer %d\n", *t) // t has type *int
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
  theFmt.Println()
  theFmt.Println("--------Test C-------------")
  TestC()

    
}
