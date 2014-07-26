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

const IPSUM string = "You think water moves fast? You should see ice. It moves like it has a mind. Like it knows it killed the world once and got a taste for murder. After the avalanche, it took us a week to climb out. Now, I don't know exactly when we turned on each other, but I know that seven of us survived the slide... and only five made it out. Now we took an oath, that I'm breaking now. We said we'd say it was the snow that killed the other two, but it wasn't. Nature is lethal but it doesn't hold a candle to man."

func NewSetupRule(desc string, arity string, deps ...*SetupRule) *SetupRule {
  return &SetupRule{0, desc, IPSUM, arity, deps}
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
