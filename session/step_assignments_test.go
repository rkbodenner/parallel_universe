package session

import (
  "testing"
  "github.com/rkbodenner/parallel_universe/game"
)

func TestGetAssignees(t *testing.T) {
  p1 := game.Player{1, "1"}
  p2 := game.Player{2, "2"}
  r1 := game.SetupRule{"1", "Once", nil}
  r2 := game.SetupRule{"2", "Each player", nil}

  assignments := NewStepMap()
  var step game.SetupStep
  step,_ = game.NewGlobalSetupStep(r1)
  assignments.Set(&p1, step)

  var assignees []*game.Player

  assignees = assignments.GetAssignees(&r1)
  if len(assignees) != 1 {
    t.Fatal("Should have one assigned player")
  }
  if *assignees[0] != p1 {
    t.Fatal("Player 1 should be the assignee")
  }

  step,_ = game.NewSinglePlayerSetupStep(r2, &p1)
  assignments.Set(&p1, step)
  step,_ = game.NewSinglePlayerSetupStep(r2, &p2)
  assignments.Set(&p2, step)

  assignees = assignments.GetAssignees(&r2)
  if len(assignees) != 2 {
    t.Fatal("Should have two assigned players")
  }
  foundP1, foundP2 := false, false
  for _,p := range assignees {
    if *p == p1 {
      foundP1 = true
    }
    if *p == p2 {
      foundP2 = true
    }
  }
  if !(foundP1 && foundP2) {
    t.Fatal("Both players should be assigned")
  }
}
