package main

import (
 "fmt"
)


func TestContA() {
  guestList := []string{"bill", "jill", "joan"}
  arrived := []string{"sally", "jill", "joan"} 

  for _, guest := range guestList {
    fmt.Println("Guest ", guest)
    for _, person := range arrived {
      fmt.Println("Person ", person)

      if person == guest {
        fmt.Printf("Let %s In\n", person)
        continue 
      }
    }
  }
}



func TestContB() {
  guestList := []string{"bill", "jill", "joan"}
  arrived := []string{"sally", "jill", "joan"} 

CheckList:
  for _, guest := range guestList {
    fmt.Println("Guest ", guest)
    for _, person := range arrived {
      fmt.Println("Person ", person)

      if person == guest {
        fmt.Printf("Let %s In\n", person)
        continue CheckList
      }
    }
  }
}

func DummySender() bool {
    return true
}

func TestBreakA() {
  guestList := []string{"bill", "jill", "joan"}
  arrived := []string{"sally", "jill", "joan"} 

Loop:
  for _, guest := range guestList {
    fmt.Println("Guest ", guest)
    for _, person := range arrived {
      fmt.Println("Person ", person)

      if person == guest {
        err := DummySender()
        // notice the use of break <label> to stop the processing
        // in case of error. If break does not have label then
        // we will need to keep a flag to break from both loops.
        if err {
          break Loop
        }
        fmt.Printf("Let %s In\n", person)
      }
    }
  }
}



func main() {

  //Test A and Test B does the same thing but Test A is more efficient
  //Check the same and realize the importance of continue <label>
  fmt.Println("--------Test Continue A-------------")
  TestContA()
  fmt.Println()
  fmt.Println("--------Test Continue B-------------")
  TestContB()
  fmt.Println()
  fmt.Println()
  fmt.Println("--------Test Break A-------------")
  TestBreakA()
  fmt.Println()
}

