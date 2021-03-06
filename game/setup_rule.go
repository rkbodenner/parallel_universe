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

func (a *SetupRule) Equal(b *SetupRule) bool {
  return a.Id == b.Id &&
    a.Description == b.Description &&
    a.Details == b.Details &&
    a.Arity == b.Arity
}

type SetupStep struct {
  Rule *SetupRule
  Owner *Player
  Done bool
}

func NewGlobalSetupStep(rule *SetupRule) (*SetupStep, error) {
  if "Once" != rule.Arity {
    return nil, fmt.Errorf("Global setup step cannot be created from a rule with different arity")
  }
  return &SetupStep{rule, nil, false}, nil
}

func NewPerPlayerSetupStep(rule *SetupRule, owner *Player) (*SetupStep, error) {
  if "Each player" != rule.Arity {
    return nil, fmt.Errorf("Per-player setup step cannot be created from a rule with different arity")
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
  if a.Rule.Equal(b.Rule) {
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
  ownerName := ""
  if nil != step.Owner {
    ownerName = step.Owner.Name
  }

  return fmt.Sprintf("%s (%s)", step.Rule.Description, ownerName)
}
