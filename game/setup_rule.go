package game

import (
  "fmt"
)

type SetupRule struct {
  Description string
  Arity string
}

type SetupStep interface {
  Rule() SetupRule
  Owner() Player
  CanBeOwnedBy(Player) bool
  SetDone()
  Done() bool
}

type SinglePlayerSetupStep struct {
  rule SetupRule
  owner Player
  done bool
}

func NewSinglePlayerSetupStep(rule SetupRule, owner Player) (*SinglePlayerSetupStep, error) {
  if "Each player" != rule.Arity {
    return nil, fmt.Errorf("Setup rule must be for each player")
  }
  return &SinglePlayerSetupStep{rule: rule, owner: owner, done: false}, nil
}

func (step *SinglePlayerSetupStep) Rule() SetupRule {
  return step.rule
}

func (step *SinglePlayerSetupStep) Owner() Player {
  return step.owner
}

func (step *SinglePlayerSetupStep) CanBeOwnedBy(player Player) bool {
  return player == step.Owner()
}

func (step *SinglePlayerSetupStep) SetDone() {
  step.done = true
}

func (step *SinglePlayerSetupStep) Done() bool {
  return step.done
}

type GlobalSetupStep struct {
  rule SetupRule
  done bool
}

func NewGlobalSetupStep(rule SetupRule) (*GlobalSetupStep, error) {
  if "Once" != rule.Arity {
    return nil, fmt.Errorf("Setup rule must be done once")
  }
  return &GlobalSetupStep{rule: rule, done: false}, nil
}

func (step *GlobalSetupStep) Rule() SetupRule {
  return step.rule
}

func (step *GlobalSetupStep) Owner() Player {
  return Player("global")
}

func (step *GlobalSetupStep) CanBeOwnedBy(player Player) bool {
  return true
}

func (step *GlobalSetupStep) SetDone() {
  step.done = true
}

func (step *GlobalSetupStep) Done() bool {
  return step.done
}
