package game

import (
  "fmt"
)

type SetupRule struct {
  Id int
  Description string
  Arity string
  Dependencies []*SetupRule
}

func NewSetupRule(desc string, arity string, deps ...*SetupRule) *SetupRule {
  return &SetupRule{0, desc, arity, deps}
}

type SetupStep struct {
  Rule *SetupRule
  Owner *Player
  Done bool
}

func NewGlobalSetupStep(rule *SetupRule) (*SetupStep, error) {
  if "Once" != rule.Arity {
    return nil, fmt.Errorf("Setup rule must be done once")
  }
  return &SetupStep{rule, nil, false}, nil
}

func NewSinglePlayerSetupStep(rule *SetupRule, owner *Player) (*SetupStep, error) {
  if "Each player" != rule.Arity {
    return nil, fmt.Errorf("Setup rule must be for each player")
  }
  return &SetupStep{rule, owner, false}, nil
}

func (step *SetupStep) CanBeOwnedBy(player *Player) bool {
  return step.Owner == nil || player.Id == step.Owner.Id
}

func (step *SetupStep) Finish() {
  step.Done = true
}

func (step *SetupStep) String() string {
  return fmt.Sprintf("%s", step.Rule.Description)
}
