package main

import (
  "fmt"
  "github.com/rkbodenner/parallel_universe/collection"
  "github.com/rkbodenner/parallel_universe/game"
  "github.com/rkbodenner/parallel_universe/session"
)

func main() {
  var players = []*game.Player{
    &game.Player{1, "Player One"},
    &game.Player{2, "Player Two"},
  }
  game := collection.NewForbiddenIsland()
  session := session.NewSession(game, players)

  fmt.Println("game on")

  session.NextStep(session.Players[0])
  session.NextStep(session.Players[1])
  session.NextStep(session.Players[0])
  session.NextStep(session.Players[1])
  session.NextStep(session.Players[0])
  session.NextStep(session.Players[0])
  session.NextStep(session.Players[0])
  session.NextStep(session.Players[1])
  session.NextStep(session.Players[0])
  session.PrintStepAssignments()

  fmt.Println()
  game.PrintSetupRules()
}
