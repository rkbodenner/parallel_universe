package main

import (
  "fmt"
  "github.com/rkbodenner/parallel_universe/collection"
  "github.com/rkbodenner/parallel_universe/session"
)

func main() {
  fmt.Println("game on")
  game := collection.NewForbiddenIsland()
  session := session.NewSession(game, 2)

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

  fmt.Println();
  game.PrintSetupRules()
}
