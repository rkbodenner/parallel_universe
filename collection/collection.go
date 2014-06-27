package collection

import (
  "github.com/rkbodenner/parallel_universe/game"
)

func NewTicTacToe() *game.Game {
  var setup = []*game.SetupRule{
    game.NewSetupRule("Draw 3x3 grid", "Once"),
    game.NewSetupRule("Choose X or O", "Each player"),
  }

  return game.NewGame("Tic-Tac-Toe", setup)
}

func NewForbiddenIsland() *game.Game {
  var setup = []*game.SetupRule{
    game.NewSetupRule("Create Forbidden Island", "Once"),
    game.NewSetupRule("Place the treasures", "Once"),
    game.NewSetupRule("Divide the cards", "Once"),
    game.NewSetupRule("The island starts to sink", "Once"),
    game.NewSetupRule("Deal Adventurer cards", "Once"),
    game.NewSetupRule("Place Adventurer pawn", "Each player"),
    game.NewSetupRule("Hand out Treasure deck cards", "Once"),
    game.NewSetupRule("Set the water level", "Once"),
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
