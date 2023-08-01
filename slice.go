package main

import (
  "fmt"
)

func main() {
  arr := []int{1, 5, 2, 4, 3, 7, 9, 8, 6, 0}

  fmt.Println("arr:", arr)
  fmt.Println("arr[0]:", arr[0])
  fmt.Println("arr[5]:", arr[5])
  fmt.Println("arr[0:4]:", arr[0:4])
  fmt.Println("arr[3:5]:", arr[3:5])
  fmt.Println("arr[2:]:", arr[2:])
  fmt.Println("arr[:3]:", arr[:3])
  fmt.Println("arr[2:-1]: invalid argument: index -1 (constant of type int) must not be negative")
}