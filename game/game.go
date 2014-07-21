package game

import (
  "fmt"
)

type Game struct {
  Id uint `json:"id"`
  Name string `json:"name"`
  SetupRules []*SetupRule
  MinPlayers int `json:"min_players"`
  MaxPlayers int `json:"max_players"`
}

func NewGame(name string, rules []*SetupRule, minPlayers int, maxPlayers int) *Game {
  return &Game{
    Name: name,
    SetupRules: rules,
    MinPlayers: minPlayers,
    MaxPlayers: maxPlayers,
  }
}

func (game *Game) PrintSetupRules() {
  for _,r := range game.SetupRules {
    fmt.Printf("%s\t%s\n", r.Description, r.Arity)
  }
}
