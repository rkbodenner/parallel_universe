package main

import (
  "fmt"
)

type SetupRules struct {
  Steps []SetupStep
}

type SetupStep struct {
  Description string
  Arity string
}

var TicTacToeSetup = SetupRules{
  []SetupStep{
    {"Draw 3x3 grid", "Once"},
    {"Choose X or O", "Each player"},
  },
}

func main() {
  fmt.Println("game on")
  for _,r := range TicTacToeSetup.Steps {
    fmt.Printf("%s\t%s\n", r.Description, r.Arity)
  }
}
