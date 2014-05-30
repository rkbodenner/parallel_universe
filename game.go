package main

import (
  "fmt"
)

type Game struct {
  Setup *SetupRules
  PlayerCount int
}

func NewGame(rules *SetupRules, playerCount int) *Game {
  return &Game{
    Setup: rules,
    PlayerCount: playerCount,
  }
}

func (game *Game) PrintSteps() (error) {
  for _,r := range game.Setup.Steps {
    if "Each player" == r.Arity {
      for i := 0; i < game.PlayerCount; i++ {
        fmt.Printf("%s\tPlayer %d\n", r.Description, i)
      }
    } else {
      fmt.Printf("%s\t%s\n", r.Description, r.Arity)
    }
  }
  return nil
}

type SetupRules struct {
  Steps []SetupStep
}

type SetupStep struct {
  Description string
  Arity string
}

func NewTicTacToe() *Game {
  var setup = SetupRules{
    []SetupStep{
      {"Draw 3x3 grid", "Once"},
      {"Choose X or O", "Each player"},
    },
  }

  return NewGame(&setup, 2)
}

func NewForbiddenIsland() *Game {
  var setup = SetupRules{
    []SetupStep{
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

  return NewGame(&setup, 2)
}

func main() {
  fmt.Println("game on")
  game := NewForbiddenIsland()
  game.PrintSteps()
}
