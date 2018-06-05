package main

import (
  "fmt"
  "math"
)

type MyFloat float64

func (f MyFloat) Abs() flaot64 {
  if f < 0 {
    return float64(-f)
  }
  return float 64(f)
}

func main() {
  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())
}
