package main

import (
  "fmt"
  "github.com/rkbodenner/parallel_universe/game"
)

func NewTicTacToe() *game.Game {
  var setup = game.SetupRules{
    []game.SetupStep{
      {"Draw 3x3 grid", "Once"},
      {"Choose X or O", "Each player"},
    },
  }

  return game.New(&setup, 2)
}

func NewForbiddenIsland() *game.Game {
  var setup = game.SetupRules{
    []game.SetupStep{
      {"Create Forbidden Island", "Once"},
      {"Place the treasures", "Once"},
      {"Divide the cards", "Once"},
      {"The island starts to sink", "Once"},
      {"Deal Adventurer cards", "Once"},
      {"Place Adventurer pawn", "Each player"},
      {"Hand out Treasure deck cards", "Once"},
      {"Set the water level", "Once"},
    },
  }

  return game.New(&setup, 2)
}

func main() {
  fmt.Println("game on")
  game := NewForbiddenIsland()
  game.AssignSteps()
  game.PrintStepAssignments()
}
