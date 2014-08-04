package game

import (
  "fmt"
)

type SetupRule struct {
  Id int
  Description string
  Details string
  Arity string
  Dependencies []*SetupRule
}

func NewSetupRule(desc string, arity string, deps ...*SetupRule) *SetupRule {
  return &SetupRule{0, desc, "", arity, deps}
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

func (a *SetupStep) Equal(b *SetupStep) bool {
  if a.Rule.Id == b.Rule.Id {
    if a.Owner == nil && b.Owner == nil {
      return true
    }
    if a.Owner != nil && b.Owner != nil {
      if a.Owner.Id == b.Owner.Id {
        return true
      }
    }
  }
  return false
}

func (step *SetupStep) String() string {
  return fmt.Sprintf("%s", step.Rule.Description)
}
