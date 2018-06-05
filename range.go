package main

import "fmt"

func main() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << uint(i) // == 2**i
  }
  // skip the index value
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
  //with index value
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}
