package session

import (
  "reflect"
  "runtime"
  "strings"
  "testing"
  "github.com/rkbodenner/parallel_universe/game"
)

func TestNewSessionEachPlayerAssignment(t *testing.T) {
  rules := []game.SetupRule{
    game.SetupRule{"Only step", "Each player", nil},
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

type stepState func(rule *game.SetupRule, session *Session) (inThisState bool)

func done(rule *game.SetupRule, session *Session) bool {
  return session.IsRuleDone(rule)
}

func undone(rule *game.SetupRule, session *Session) bool {
  return !done(rule, session)
}

func assigned(rule *game.SetupRule, session *Session) bool {
  return len(session.SetupAssignments.GetAssignees(rule)) != 0
}

func unassigned(rule *game.SetupRule, session *Session) bool {
  return !assigned(rule, session)
}

func verifyStates(t *testing.T, session *Session, stepStates map[*game.SetupRule][]stepState) {
  for rule,states := range stepStates {
    for _,state := range states {
      if !state(rule, session) {
        fnNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(state).Pointer()).Name(), ".")
        t.Fatalf("Rule %s expected to be %s", rule.Description, fnNameParts[len(fnNameParts)-1])
      }
    }
  }
}

func TestStepHonorsDependencies(t *testing.T) {
  rule0 := game.SetupRule{"0", "Once", nil}
  rule1 := game.SetupRule{"1", "Once", []*game.SetupRule{&rule0}}
  rule2 := game.SetupRule{"2", "Once", []*game.SetupRule{&rule1}}

  players := []*game.Player{
    &game.Player{1, "Alice"},
    &game.Player{2, "Bob"},
  }
  g := game.NewGame("war", []game.SetupRule{rule0, rule1, rule2})
  session := NewSession(g, players)

  states := make(map[*game.SetupRule][]stepState)
  var step game.SetupStep

  states[&rule0] = []stepState{undone, unassigned}
  states[&rule1] = []stepState{undone, unassigned}
  states[&rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step = session.Step(players[0])
  states[&rule0] = []stepState{undone, assigned}
  states[&rule1] = []stepState{undone, unassigned}
  states[&rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  // No change
  session.Step(players[0])
  states[&rule0] = []stepState{undone, assigned}
  states[&rule1] = []stepState{undone, unassigned}
  states[&rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  // No change
  session.Step(players[1])
  states[&rule0] = []stepState{undone, assigned}
  states[&rule1] = []stepState{undone, unassigned}
  states[&rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step.Finish()
  states[&rule0] = []stepState{done, assigned}
  states[&rule1] = []stepState{undone, unassigned}
  states[&rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step = session.Step(players[0])
  states[&rule0] = []stepState{done, unassigned}
  states[&rule1] = []stepState{undone, assigned}
  states[&rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step.Finish()
  states[&rule0] = []stepState{done, unassigned}
  states[&rule1] = []stepState{done, assigned}
  states[&rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step = session.Step(players[1])
  states[&rule0] = []stepState{done, unassigned}
  states[&rule1] = []stepState{done, assigned}
  states[&rule2] = []stepState{undone, assigned}
  verifyStates(t, session, states)

  // No change
  session.Step(players[0])
  states[&rule0] = []stepState{done, unassigned}
  states[&rule1] = []stepState{done, assigned}
  states[&rule2] = []stepState{undone, assigned}
  verifyStates(t, session, states)

  step.Finish()
  states[&rule0] = []stepState{done, unassigned}
  states[&rule1] = []stepState{done, assigned}
  states[&rule2] = []stepState{done, assigned}
  verifyStates(t, session, states)
}

func TestStepToLastStep(t *testing.T) {
  rules := []game.SetupRule{
    // First, last, and only step
    game.SetupRule{"Only step", "Each player", nil},
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
