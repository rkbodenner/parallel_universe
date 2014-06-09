package game

import (
  "fmt"
)

type Game struct {
  Id uint `json:"id"`
  Name string `json:"name"`
  SetupRules []SetupRule
}

func NewGame(name string, rules []SetupRule) *Game {
  return &Game{
    Name: name,
    SetupRules: rules,
  }
}

func (game *Game) PrintSetupRules() {
  for _,r := range game.SetupRules {
    fmt.Printf("%s\t%s\n", r.Description, r.Arity)
  }
}
