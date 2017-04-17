package main

import (
  "fmt"
  "math"
)

type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}

const s string = "constant"

func main() {

 /// s = "variable" //not allowed
  fmt.Println("s: ", s)

  const n = 500000000
  
  const d = 3e20 / n
  fmt.Println(d)

  fmt.Println(int64(d))
  fmt.Println(math.Sin(n))

  fmt.Println("---------------------------")
  fmt.Println()

  fmt.Println("MB: ", MB, float64(MB))
  fmt.Println("KB: ", KB, float64(KB))
  fmt.Println("GB: ", GB, float64(GB))

  var bb ByteSize = 45000
  fmt.Println("bb: ", bb)
    
}
