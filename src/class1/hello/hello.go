package main

import "fmt"
import "class1/mystring"

func printit(s string) {

  if oe, e := mystring.Iseven(s); e == nil {
    if oe == true {
      fmt.Println ("Hello", s, "- Even fellow")
    } else {
      fmt.Println ("Hello", s, "- Odd fellow")
    }
  } else {
    fmt.Println("Error ", e)
  }
}

func main() {

  tests := []string {"Vijay", "Kavi", ""}

  for _, s := range tests {
    printit(s)
  }
}
