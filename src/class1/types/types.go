package main

import fmt "fmt"

type Employee struct {
  id int
  name string
}

type Department struct {
   int
   string
}


func main(){

  var e1 Employee = Employee{1, "kaushik"}
  fmt.Println("e1: ", e1)
  fmt.Println("e1.id: ", e1.id)
  fmt.Println("e1.name: ", e1.name)

  e2 := Employee{2, "Aditya"}
  fmt.Println("e2: ", e2)

  e3 := & Employee{3, "Aditi"}
  fmt.Println("e3: ", e3)
  fmt.Println("e3: ", *e3)

  d1 := Department {1, "Engineering"}
  fmt.Println("d1: ", d1)
  fmt.Println("d1.int: ", d1.int)
  fmt.Println("d1.string: ", d1.string)
}
