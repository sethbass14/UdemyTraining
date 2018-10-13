package main

import "fmt"

func factorial(x int) int {
  if x == 0 {
    return 1
  }
  return x * factorail(x - 1)
}

func main() {
  fmt.Println(factorial(4))
}
