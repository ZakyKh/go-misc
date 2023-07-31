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
  var str string
  var err error

  for (true) {
    str, err = reader.ReadString('\n')
    if err != nil {
      log.Fatal(err)
    }
    nums := strings.Split(strings.Trim(str, " \t\n"), " ")
    if len(nums) < 2 {
      fmt.Println("Please input at least 2 integers")
      continue
    }
    a, err1 := strconv.Atoi(nums[0])
    if err1 != nil {
      fmt.Println("First token must be an integer")
    }

    b, err2 := strconv.Atoi(nums[1])
    if err2 != nil {
      fmt.Println("Second token must be an integer")
    } else if b == 0 {
      fmt.Println("Second integer cannot be 0")
    }

    if err1 != nil || err2 != nil || b == 0 {
      continue
    }

    fmt.Printf("%d + %d = %d\n", a, b, a + b)
    fmt.Printf("%d - %d = %d\n", a, b, a - b)
    fmt.Printf("%d * %d = %d\n", a, b, a * b)
    fmt.Printf("%d div %d = %d\n", a, b, a / b)
    fmt.Printf("%d mod %d = %d\n", a, b, a % b)
  }
}