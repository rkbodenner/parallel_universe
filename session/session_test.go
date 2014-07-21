package session

import (
  "reflect"
  "runtime"
  "strings"
  "testing"
  "github.com/rkbodenner/parallel_universe/game"
)

func TestNewSessionEnforcePlayerLimits(t *testing.T) {
  players := []*game.Player{
    &game.Player{1, "Alice"},
    &game.Player{2, "Bob"},
    &game.Player{3, "Carol"},
  }
  g := game.NewGame("war", nil, 1, len(players) - 1)

  var session *Session
  var err error

  session, err = NewSession(g, []*game.Player{})
  if err == nil || session != nil {
    t.Fatal("Cannot create a session with less than the minimum number of players")
  }

  session, err = NewSession(g, players[:0])
  if err != nil && session != nil {
    t.Fatal("Unexpected error creating session")
  }

  session, err = NewSession(g, players[0:])
  if err == nil || session != nil {
    t.Fatal("Cannot create a session with more than the minimum number of players")
  }
}

func TestNewSessionEachPlayerAssignment(t *testing.T) {
  rules := []*game.SetupRule{
    game.NewSetupRule("Only step", "Each player"),
  }
  players := []*game.Player{
    &game.Player{1, "Alice"},
    &game.Player{2, "Bob"},
  }
  g := game.NewGame("war", rules, len(players), len(players))
  session, err := NewSession(g, players)
  if nil != err {
    t.Fatal("Error creating session")
  }

  var aliceStep *game.SetupStep
  var bobStep *game.SetupStep
  for _,step := range session.SetupSteps {
    if step.Owner == players[0] {
      if aliceStep == nil {
        aliceStep = step
      } else {
        t.Fatal("One step per player")
      }
    }
    if step.Owner == players[1] {
      if bobStep == nil {
        bobStep = step
      } else {
        t.Fatal("One step per player")
      }
    }
  }
  if aliceStep == nil || aliceStep.Rule.Description != "Only step" {
    t.Fatal("Each player assigned a step")
  }
  if bobStep == nil || bobStep.Rule.Description != "Only step" {
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
  rule0 := game.NewSetupRule("0", "Once")
  rule1 := game.NewSetupRule("1", "Once", rule0)
  rule2 := game.NewSetupRule("2", "Once", rule1)

  players := []*game.Player{
    &game.Player{1, "Alice"},
    &game.Player{2, "Bob"},
  }
  g := game.NewGame("war", []*game.SetupRule{rule0, rule1, rule2}, len(players), len(players))
  session, err := NewSession(g, players)
  if nil != err {
    t.Fatal("Error creating session")
  }

  states := make(map[*game.SetupRule][]stepState)
  var step *game.SetupStep

  states[rule0] = []stepState{undone, unassigned}
  states[rule1] = []stepState{undone, unassigned}
  states[rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step = session.Step(players[0])
  states[rule0] = []stepState{undone, assigned}
  states[rule1] = []stepState{undone, unassigned}
  states[rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  // No change
  session.Step(players[0])
  states[rule0] = []stepState{undone, assigned}
  states[rule1] = []stepState{undone, unassigned}
  states[rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  // No change
  session.Step(players[1])
  states[rule0] = []stepState{undone, assigned}
  states[rule1] = []stepState{undone, unassigned}
  states[rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step.Finish()
  states[rule0] = []stepState{done, assigned}
  states[rule1] = []stepState{undone, unassigned}
  states[rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step = session.Step(players[0])
  states[rule0] = []stepState{done, unassigned}
  states[rule1] = []stepState{undone, assigned}
  states[rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step.Finish()
  states[rule0] = []stepState{done, unassigned}
  states[rule1] = []stepState{done, assigned}
  states[rule2] = []stepState{undone, unassigned}
  verifyStates(t, session, states)

  step = session.Step(players[1])
  states[rule0] = []stepState{done, unassigned}
  states[rule1] = []stepState{done, assigned}
  states[rule2] = []stepState{undone, assigned}
  verifyStates(t, session, states)

  // No change
  session.Step(players[0])
  states[rule0] = []stepState{done, unassigned}
  states[rule1] = []stepState{done, assigned}
  states[rule2] = []stepState{undone, assigned}
  verifyStates(t, session, states)

  step.Finish()
  states[rule0] = []stepState{done, unassigned}
  states[rule1] = []stepState{done, assigned}
  states[rule2] = []stepState{done, assigned}
  verifyStates(t, session, states)
}

func TestStepToLastStep(t *testing.T) {
  rules := []*game.SetupRule{
    // First, last, and only step
    game.NewSetupRule("Only step", "Each player"),
  }
  players := []*game.Player{
    &game.Player{1, "Alice"},
    &game.Player{2, "Bob"},
  }
  g := game.NewGame("war", rules, len(players), len(players))
  session, err := NewSession(g, players)
  if nil != err {
    t.Fatal("Error creating session")
  }

  var next *game.SetupStep
  next = session.Step(players[0])
  if next == nil || next.Rule.Description != "Only step" {
    t.Fatal("Should stay at last and only step")
  }

  next = session.Step(players[1])
  if next == nil || next.Rule.Description != "Only step" {
    t.Fatal("Should stay at last and only step")
  }
}
