package main

import (
  "fmt"
)

type Player string

type Game struct {
  Setup *SetupRules
  Players []Player
}

func NewGame(rules *SetupRules, playerCount uint) *Game {
  players := make([]Player, playerCount)
  for i := range players {
    players[i] = (Player)(fmt.Sprintf("Player %d", i+1))
  }
  return &Game{
    Setup: rules,
    Players: players,
  }
}

func (game *Game) PlayerCount() int {
  return len(game.Players)
}

func (game *Game) PrintSteps() error {
  for _,r := range game.Setup.Steps {
    if "Each player" == r.Arity {
      for _,p := range game.Players {
        fmt.Printf("%s\t%s\n", r.Description, p)
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
