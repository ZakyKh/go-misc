package main

import (
  "bufio"
  "errors"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func getErrorMessage(errType string, idx int) string {
  if errType == "invalid_query" {
    return `Invalid query format.
                 Accepted format:
                     'EXECUTE <expression>' where expression is a 2-operand integer operation,
                                            the operator is one of + - * / %, and the second
                                            integer is non-zero if the operator is / or %.
                     or
                     'EVALUATE <equation>'  where equation is a 2-side 2-operand integer
                                            operation equality check where the check operator
                                            is one of < > <= >= = # (inequal), the operator
                                            in each operation is one of + - * / %, and the
                                            second integer in each operation is non-zero if
                                            the operator for that expression is / or %.`
  } else if errType == "invalid_integer" {
    return fmt.Sprintf("Token %d must be an integer", idx)
  } else if errType == "invalid_operator" {
    return fmt.Sprintf("Token %d must be an operator, one of + - * / \%", idx)
  } else if errType == "invalid_div_0" {
    return fmt.Sprintf("Token %d cannot be 0 if operator is either / or \%", idx)
  } else if errType == "invalid_eq_operator" {
    return fmt.Sprintf("Token %d be an equality/inequality operator, one of < > <= >= = # (inequal)")
  } else {
    return "Unknown error"
  }
}

func execute(a, b int, sym string) int {
  if sym == "+" {
    return a + b
  } else if sym == "-" {
    return a - b
  } else if sym == "*" {
    return a * b
  } else if sym == "/" {
    return a / b
  } else {
    return a % b
  }
}

func evaluate(a, b int, sym string) bool {
  if sym == "<" {
    return a < b
  } else if sym == ">" {
    return a > b
  } else if sym == "<=" {
    return a <= b
  } else if sym == ">=" {
    return a >= b
  } else if sym == "=" {
    return a == b
  } else {
    return a != b
  }
}

func parseExpression(args []string, offset int) (int, int, string, []error) {
  var errs []error
  a, err1 := strconv.Atoi(args[0])
  if err1 != nil {
    errs = append(errs, errors.New(getErrorMessage("invalid_integer", 0 + offset)))
  }

  sym := args[1]
  okSym := sym == "+" || sym == "-" || sym == "*" || sym == "/" || sym == "%"
  if !okSym {
    errs = append(errs, errors.New(getErrorMessage("invalid_operator", 1 + offset)))
  }

  b, err3 := strconv.Atoi(args[2])
  if err3 != nil {
    errs = append(errs, errors.New(getErrorMessage("invalid_integer", 2 + offset)))
  } else if b == 0 && (sym == "/" || sym == "%") {
    errs = append(errs, errors.New(getErrorMessage("invalid_div_0", 2 + offset)))
  }

  return a, b, sym, errs
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  var query string
  var err error

  for (true) {
    query, err = reader.ReadString('\n')
    if err != nil {
      log.Fatal(err)
    }
    tokens := strings.Split(strings.Trim(query, " \t\n"), " ")
    if len(tokens) < 1 {
      fmt.Println(getErrorMessage("invalid_query", -1))
      continue
    }

    directive := tokens[0]
    if directive == "EXECUTE" {
      if len(tokens) < 4 {
        fmt.Println(getErrorMessage("invalid_query", -1))
        continue
      }

      a, b, sym, errs := parseExpression(tokens[1:], 1)
      if len(errs) > 0 {
        for _, err := range errs {
          fmt.Println(err.Error())
        }
        continue
      }
      fmt.Printf("%d %s %d = %d\n", a, sym, b, execute(a, b, sym))
    } else if directive == "EVALUATE" {
      if len(tokens) < 8 {
        fmt.Println(getErrorMessage("invalid_query", -1))
        continue
      }

      eqSym := tokens[4]
      okEqSym := eqSym == "<" || eqSym == ">" || eqSym == "<=" || eqSym == ">=" || eqSym == "=" || eqSym == "#"
      if !okEqSym {
        fmt.Println(getErrorMessage("invalid_eq_operator", 4))
        continue
      }

      a1, b1, sym1, errs1 := parseExpression(tokens[1:4], 1)
      a2, b2, sym2, errs2 := parseExpression(tokens[5:8], 5)
      if len(errs1) > 0 {
        for _, err := range errs1 {
          fmt.Println(err.Error())
        }
      }
      if len(errs2) > 0 {
        for _, err := range errs2 {
          fmt.Println(err.Error())
        }
      }

      if len(errs1) + len(errs2) > 0 {
        continue
      }

      fmt.Printf("%d %s %d %s %d %s %d is %t\n", a1, sym1, b1, eqSym, a2, sym2, b2, evaluate(execute(a1, b1, sym1), execute(a2, b2, sym2), eqSym))
    } else {
      fmt.Println(getErrorMessage("invalid_query", -1))
      continue
    }
  }
}