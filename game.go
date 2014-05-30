package main

import (
  "fmt"
)

type SetupRules struct {
  Steps []string
}

var TicTacToeSetup = SetupRules{
  []string{
    "Draw 3x3 grid",
    "Choose X or O",
  },
}

func main() {
  fmt.Println("game on")
  for _,r := range TicTacToeSetup.Steps {
    fmt.Println(r)
  }
}
