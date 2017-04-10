package main

import (
 "fmt"
)

func FuncA() {
  fmt.Println("function A called")
}

func FuncB() int {
  i := 10
  fmt.Println("function B returning ", i)
  return i
}

func FuncC() (i int) {
  //named return 
  i = 100
  fmt.Println("function C returning ", i)
  return //we do not need to return any variable
  //i will be returned automatically
}

func AddNumbers(nums... int) (sum int) {
  for _, v := range nums {
    sum += v
  }
  return
}

//functions can return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

func main() {

  fmt.Println("--------Test func A-------------")
  FuncA()
  fmt.Println()
  fmt.Println("--------Test func B-------------")
  i := FuncB()
  fmt.Println("FuncB returned: ", i)
  fmt.Println()
  fmt.Println("--------Test func C-------------")
  j := FuncC()
  fmt.Println("FuncC returned: ", j)
  fmt.Println()
  fmt.Println("--------Test Add Numbers-------------")
  sum := AddNumbers(5, 5, 29, 30)
  fmt.Println("AddNumbers returned: ", sum)
  fmt.Println()
  fmt.Println("--------Test Add Numbers-------------")
  a, b := "A", "B"
  fmt.Println("Before Swap a:", a, ", b:", b)
  a, b = swap(a, b)
  fmt.Println("After Swap a:", a, ", b:", b)
  fmt.Println()
}

