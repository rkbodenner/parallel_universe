package collection

import (
  "github.com/rkbodenner/parallel_universe/game"
)

func NewTicTacToe() *game.Game {
  var setup = []game.SetupRule{
    {"Draw 3x3 grid", "Once", []*game.SetupRule{}},
    {"Choose X or O", "Each player", []*game.SetupRule{}},
  }

  return game.NewGame("Tic-Tac-Toe", setup)
}

func NewForbiddenIsland() *game.Game {
  var setup = []game.SetupRule{
    {"Create Forbidden Island", "Once", []*game.SetupRule{}},
    {"Place the treasures", "Once", []*game.SetupRule{}},
    {"Divide the cards", "Once", []*game.SetupRule{}},
    {"The island starts to sink", "Once", []*game.SetupRule{}},
    {"Deal Adventurer cards", "Once", []*game.SetupRule{}},
    {"Place Adventurer pawn", "Each player", []*game.SetupRule{}},
    {"Hand out Treasure deck cards", "Once", []*game.SetupRule{}},
    {"Set the water level", "Once", []*game.SetupRule{}},
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
