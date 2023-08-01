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
    if len(nums) < 3 {
      fmt.Println("Please input at least 2 integers and one of + - * / % between them separated by a space")
      continue
    }
    a, err1 := strconv.Atoi(nums[0])
    if err1 != nil {
      fmt.Println("First token must be an integer")
    }

    sym := nums[1]
    okSym := sym == "+" || sym == "-" || sym == "*" || sym == "/" || sym == "%"
    if !okSym {
      fmt.Println("Second token must be an operator, one of + - * / %")
    }

    b, err3 := strconv.Atoi(nums[2])
    if err3 != nil {
      fmt.Println("Third token must be an integer")
    } else if b == 0 && (sym == "/" || sym == "%") {
      fmt.Println("Third token, the second integer, cannot be 0 if operator is either / or %")
    }

    if err1 != nil || err3 != nil || b == 0 {
      continue
    }

    if sym == "+" {
      fmt.Printf("%d + %d = %d\n", a, b, a + b)
    } else if sym == "-" {
      fmt.Printf("%d - %d = %d\n", a, b, a - b)
    } else if sym == "*" {
      fmt.Printf("%d * %d = %d\n", a, b, a * b)
    } else if sym == "/" {
      fmt.Printf("%d div %d = %d\n", a, b, a / b)
    } else {
      fmt.Printf("%d mod %d = %d\n", a, b, a % b)
    }
  }
}