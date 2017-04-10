package main

import "fmt"

func fiboFact() func() int {

  i := 0
  j := 0

  return func() int {
    switch {
      case i == 0 :
        i++
        return i
      case j == 0 :
        j++
        return j
      default :
        ti := i 
        i += j
        j = ti
        return i
    }
  }
}

func main() {

  myfibo := fiboFact()
  for i:=0; i<20; i++ {
    fmt.Print(myfibo(), " ");
  }
  fmt.Println();

  newfibo := fiboFact()
  for i:=0; i<25; i++ {
    fmt.Print(newfibo(), " ");
  }
  fmt.Println();
}
