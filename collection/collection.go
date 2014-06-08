package collection

import (
  "github.com/rkbodenner/parallel_universe/game"
)

func NewTicTacToe() *game.Game {
  var setup = []game.SetupRule{
    {"Draw 3x3 grid", "Once"},
    {"Choose X or O", "Each player"},
  }

  return game.NewGame("Tic-Tac-Toe", setup)
}

func NewForbiddenIsland() *game.Game {
  var setup = []game.SetupRule{
    {"Create Forbidden Island", "Once"},
    {"Place the treasures", "Once"},
    {"Divide the cards", "Once"},
    {"The island starts to sink", "Once"},
    {"Deal Adventurer cards", "Once"},
    {"Place Adventurer pawn", "Each player"},
    {"Hand out Treasure deck cards", "Once"},
    {"Set the water level", "Once"},
  }

  return game.NewGame("Forbidden Island", setup)
}

type Collection struct {
  Games []*game.Game
}

func NewCollection() *Collection {
  return &Collection{
    []*game.Game{
      NewTicTacToe(),
      NewForbiddenIsland(),
    },
  }
}
