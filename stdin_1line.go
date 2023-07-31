package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  str, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
  fmt.Print(str)
}