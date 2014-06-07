package main

import (
  "fmt"
  "github.com/rkbodenner/parallel_universe/collection"
)

func main() {
  fmt.Println("game on")
  game := collection.NewForbiddenIsland()

//  game.AssignSteps()
//  game.PrintStepAssignments()

  game.NextStep(game.Players[0])
  game.NextStep(game.Players[1])
  game.NextStep(game.Players[0])
  game.NextStep(game.Players[1])
  game.NextStep(game.Players[0])
  game.NextStep(game.Players[0])
  game.NextStep(game.Players[0])
  game.NextStep(game.Players[1])
  game.NextStep(game.Players[0])
  game.PrintStepAssignments()

  fmt.Println();
  game.PrintSetupRules()
}
