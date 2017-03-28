package mystring

import "errors"

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
