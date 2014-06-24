package session

import (
  "testing"
  "github.com/rkbodenner/parallel_universe/game"
)

func TestNewSessionEachPlayerAssignment(t *testing.T) {
  rules := []game.SetupRule{
    game.SetupRule{"Only step", "Each player"},
  }
  players := []*game.Player{
    &game.Player{1, "Alice"},
    &game.Player{2, "Bob"},
  }
  g := game.NewGame("war", rules)
  session := NewSession(g, players)

  var aliceStep game.SetupStep
  var bobStep game.SetupStep
  for _,step := range session.SetupSteps {
    if step.GetOwner() == players[0] {
      if aliceStep == nil {
        aliceStep = step
      } else {
        t.Fatal("One step per player")
      }
    }
    if step.GetOwner() == players[1] {
      if bobStep == nil {
        bobStep = step
      } else {
        t.Fatal("One step per player")
      }
    }
  }
  if aliceStep == nil || aliceStep.GetRule().Description != "Only step" {
    t.Fatal("Each player assigned a step")
  }
  if bobStep == nil || bobStep.GetRule().Description != "Only step" {
    t.Fatal("Each player assigned a step")
  }
}

func TestStepToLastStep(t *testing.T) {
  rules := []game.SetupRule{
    // First, last, and only step
    game.SetupRule{"Only step", "Each player"},
  }
  players := []*game.Player{
    &game.Player{1, "Alice"},
    &game.Player{2, "Bob"},
  }
  g := game.NewGame("war", rules)
  session := NewSession(g, players)

  var next game.SetupStep
  next = session.Step(players[0])
  if next == nil || next.GetRule().Description != "Only step" {
    t.Fatal("Should stay at last and only step")
  }

  next = session.Step(players[1])
  if next == nil || next.GetRule().Description != "Only step" {
    t.Fatal("Should stay at last and only step")
  }
}
