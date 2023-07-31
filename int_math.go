package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  str, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
  nums := strings.Split(strings.Trim(str, " \t\n"), " ")
  if len(nums) < 2 {
    fmt.Println("Please input at least 2 numbers")
  }
  a, err := strconv.Atoi(nums[0])
  if err != nil {
    log.Fatal(err)
  }

  b, err := strconv.Atoi(nums[1])
  if err != nil {
    log.Fatal(err)
  }

  if b == 0 {
    fmt.Println("Second number cannot be 0")
  }

  fmt.Printf("%d + %d = %d\n", a, b, a + b)
  fmt.Printf("%d - %d = %d\n", a, b, a - b)
  fmt.Printf("%d * %d = %d\n", a, b, a * b)
  fmt.Printf("%d div %d = %d\n", a, b, a / b)
  fmt.Printf("%d mod %d = %d\n", a, b, a % b)
}