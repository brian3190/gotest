package main

import "fmt"

func main() {
  fmt.Println("Hello World")

  // uint8
  // uint16
  // uint32
  // uint64
  // int8, int 16, int32, int 64

  var age int = 40

  var favNum float64 = 1.6180339

  // randNum := 1

  fmt.Println(age, favNum)

  fmt.Println("i like \n \n")
  fmt.Println("Newlines")

  // prints integer
  fmt.Printf("%d \n", 100)
  // prints binary representation of 100
  fmt.Printf("%b \n", 100)
  // prints character that has charcode of 44
  fmt.Printf("%c \n", 44)
  // prints hexcodes
  fmt.Printf("%x \n", 17)
  // prints value of scientific notation
  // fmt.Printf("%e \n", pi)

  i := 1

  for i <= 10 {
    fmt.Println(i)

    i++
  }

  for j :=0; j < 5; j++ {
    fmt.Println(j)
  }

  yourAge := 18

  if yourAge >= 16 {
    fmt.Println("You Can Drive")
  } else if yourAge >= 18 {
    fmt.Println("You Can Vote")
  } else {
    fmt.Println("You Can Have Fun")
  }

  switch goAge {
    case 16 : fmt.Println("Go Drive")
    case 18 : fmt.Println("Go Vote")
    default : fmt.Println("Go Have Fun")
  }
}
