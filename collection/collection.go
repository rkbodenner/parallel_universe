package collection

import (
  "github.com/rkbodenner/parallel_universe/game"
)

func NewTicTacToe() *game.Game {
  var setup = []*game.SetupRule{
    game.NewSetupRule("Draw 3x3 grid", "Once"),
    game.NewSetupRule("Choose X or O", "Each player"),
  }

  return game.NewGame("Tic-Tac-Toe", setup, 2, 2)
}

func NewForbiddenIsland() *game.Game {
  var setup = []*game.SetupRule{
    game.NewSetupRule("Create Forbidden Island", "Once"),     //0
    game.NewSetupRule("Place the treasures", "Once"),
    game.NewSetupRule("Divide the cards", "Once"),            //2
    game.NewSetupRule("The island starts to sink", "Once"),   //3
    game.NewSetupRule("Deal Adventurer cards", "Once"),       //4
    game.NewSetupRule("Place Adventurer pawn", "Each player"),//5
    game.NewSetupRule("Hand out Treasure deck cards", "Once"),//6
    game.NewSetupRule("Set the water level", "Once"),
  }
  setup[3].Dependencies = []*game.SetupRule{setup[0], setup[2]}
  setup[4].Dependencies = []*game.SetupRule{setup[2]}
  setup[5].Dependencies = []*game.SetupRule{setup[4]}
  setup[6].Dependencies = []*game.SetupRule{setup[2]}

  return game.NewGame("Forbidden Island", setup, 2, 4)
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
