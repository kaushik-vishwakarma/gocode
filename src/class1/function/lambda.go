package main

import fmt "fmt"

type Stringy func() string

func foo() string{
  return "Stringy function"
}

func takesAFunction(foo Stringy){
  fmt.Printf("takesAFunction: %v\n", foo())
}

func returnsAFunction()Stringy{
  return func()string{
    fmt.Printf("Inner stringy function\n");
    return "bar" // have to return a string to be stringy
  }
}

func main(){
  takesAFunction(foo);
  var f Stringy = returnsAFunction();
  f();
  baz := func()string{
    return "anonymous stringy\n"
  };
  fmt.Printf(baz());
}
