package mystring

import "errors"
import "fmt"

func Length(s string ) int {
  b := []byte(s)
  return len(b)
}

func Iseven (s string) (bool, error) {

  l := Length(s)

  if l == 0 {
    return false, errors.New("length is zero")
  }
  if l%2 == 0 {
    return  true, nil
  }
  return false, nil
}

type lenError struct {
    l  int
    prob string
}

func (e *lenError) Error() string {
    return fmt.Sprintf("string with length '%d' is %s", e.l, e.prob)
}

func Iseven2 (s string) (bool, error) {

  l := Length(s)

  if l == 0 {
    return false, &lenError{l, "neither even nor odd"}
  }
  if l%2 == 0 {
    return  true, nil
  }
  return false, nil
}


