package main

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  str, err := reader.ReadString('\n')
  for err != io.EOF {
    if err != nil {
      log.Fatal(err)
    }
    fmt.Print(str)
    str, err = reader.ReadString('\n')
  }
}