package main

import (
  "fmt"
  "os"
)

var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
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

  fmt.Println()
  fmt.Println()
  fmt.Println("home:", home)
  fmt.Println("user:", user)
  fmt.Println("gopath:", gopath)

    
}
