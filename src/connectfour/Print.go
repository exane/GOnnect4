package main

import (
  "fmt"
)

func print(b Board) {
  for i, val := range b.fields {
    if i != 0 && i % SIZE == 0 {
      fmt.Print("\n")
    }
    fmt.Print(val)
  }
  fmt.Print("\n")
}