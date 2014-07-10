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

type SetupStep interface {
  GetRule() *SetupRule
  GetOwner() *Player
  CanBeOwnedBy(*Player) bool
  Finish()
  IsDone() bool

  String() string
}

type SinglePlayerSetupStep struct {
  Rule *SetupRule
  Owner *Player
  Done bool
}

func NewSinglePlayerSetupStep(rule *SetupRule, owner *Player) (*SinglePlayerSetupStep, error) {
  if "Each player" != rule.Arity {
    return nil, fmt.Errorf("Setup rule must be for each player")
  }
  return &SinglePlayerSetupStep{rule, owner, false}, nil
}

func (step *SinglePlayerSetupStep) GetRule() *SetupRule {
  return step.Rule
}

func (step *SinglePlayerSetupStep) GetOwner() *Player {
  return step.Owner
}

func (step *SinglePlayerSetupStep) CanBeOwnedBy(player *Player) bool {
  return player == step.Owner
}

func (step *SinglePlayerSetupStep) Finish() {
  step.Done = true
}

func (step *SinglePlayerSetupStep) IsDone() bool {
  return step.Done
}

func (step *SinglePlayerSetupStep) String() string {
  return fmt.Sprintf("%s", step.Rule.Description)
}


type GlobalSetupStep struct {
  Rule *SetupRule
  Done bool
}

func NewGlobalSetupStep(rule *SetupRule) (*GlobalSetupStep, error) {
  if "Once" != rule.Arity {
    return nil, fmt.Errorf("Setup rule must be done once")
  }
  return &GlobalSetupStep{rule, false}, nil
}

func (step *GlobalSetupStep) GetRule() *SetupRule {
  return step.Rule
}

func (step *GlobalSetupStep) GetOwner() *Player {
  return nil
}

func (step *GlobalSetupStep) CanBeOwnedBy(player *Player) bool {
  return true
}

func (step *GlobalSetupStep) Finish() {
  step.Done = true
}

func (step *GlobalSetupStep) IsDone() bool {
  return step.Done
}

func (step *GlobalSetupStep) String() string {
  return fmt.Sprintf("%s", step.Rule.Description)
}
