package main

import "fmt"

func main() {

  name := "Todd"
  fmt.Println(&name)

  changeMe(&name)

  fmt.Println(&name)
  fmt.Println(name)
}

func changeMe(z *string) {
  fmt.Println(z)
  fmt.Prinln(*z)
  *z = "Rocky"
  fmt.Println(z)
  fmt.Print(*z)
}
