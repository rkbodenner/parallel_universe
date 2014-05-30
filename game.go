package main

import (
  "fmt"
)

type Game struct {
  Setup *SetupRules
  PlayerCount int
}

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

var TicTacToe = Game{
  Setup: &TicTacToeSetup,
  PlayerCount: 2,
}

func main() {
  fmt.Println("game on")
  for _,r := range TicTacToe.Setup.Steps {
    if "Each player" == r.Arity {
      for i := 0; i < TicTacToe.PlayerCount; i++ {
        fmt.Printf("%s\tPlayer %d\n", r.Description, i)
      }
    } else {
      fmt.Printf("%s\t%s\n", r.Description, r.Arity)
    }
  }
}
