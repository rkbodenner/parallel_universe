package game

import (
  "fmt"
)

type Player string

type Game struct {
  SetupRules []SetupRule
  Players []Player
  SetupAssignments map[Player][]SetupStep
  SetupSteps []SetupStep
}

func NewGame(rules []SetupRule, playerCount uint) *Game {
  players := make([]Player, playerCount)
  for i := range players {
    players[i] = (Player)(fmt.Sprintf("Player %d", i+1))
  }

  setupSteps := make([]SetupStep, 0)
  for _,rule := range rules {
    if "Once" == rule.Arity {
      step, err := NewGlobalSetupStep(rule)
      if nil != err {
        fmt.Println(err)
      }
      setupSteps = append(setupSteps, step)
    } else if "Each player" == rule.Arity {
      for _,p := range players {
        step, err := NewSinglePlayerSetupStep(rule, p)
        if nil != err {
          fmt.Println(err)
        }
        setupSteps = append(setupSteps, step)
      }
    }
  }

  return &Game{
    SetupRules: rules,
    Players: players,
    SetupSteps: setupSteps,
  }
}

func (game *Game) PlayerCount() int {
  return len(game.Players)
}

func (game *Game) AssignSteps() error {
  game.SetupAssignments = make(map[Player][]SetupStep)

  for _,step := range game.SetupSteps {
    // Round-robin the one-time steps amongst all players
    if "Once" == step.Rule().Arity {
      player := game.Players[0]
      for _,p := range game.Players {
        if len(game.SetupAssignments[p]) < len(game.SetupAssignments[player]) {
          player = p
        }
      }
      game.SetupAssignments[player] = append(game.SetupAssignments[player], step)
    } else if "Each player" == step.Rule().Arity {
      // Assign the step to the player it's associated with
      player := step.Owner()
      game.SetupAssignments[player] = append(game.SetupAssignments[player], step)
    }
  }
  return nil
}

func (game *Game) PrintStepAssignments() error {
  for _,player := range game.Players {
    fmt.Printf("-- %d steps for %s\n", len(game.SetupAssignments[player]), player)
    for _,step := range game.SetupAssignments[player] {
      fmt.Printf("%s", step.Rule().Description)
      if "Each player" == step.Rule().Arity {
        fmt.Println(" *")
      } else {
        fmt.Println()
      }
    }
  }
  return nil
}

func (game *Game) PrintSetupRules() error {
  for _,r := range game.SetupRules {
    if "Each player" == r.Arity {
      for _,p := range game.Players {
        fmt.Printf("%s\t%s\n", r.Description, p)
      }
    } else {
      fmt.Printf("%s\t%s\n", r.Description, r.Arity)
    }
  }
  return nil
}

type SetupRule struct {
  Description string
  Arity string
}

type SetupStep interface {
  Rule() SetupRule
  Owner() Player
}

type SinglePlayerSetupStep struct {
  rule SetupRule
  owner Player
  Done bool
}

func NewSinglePlayerSetupStep(rule SetupRule, owner Player) (*SinglePlayerSetupStep, error) {
  if "Each player" != rule.Arity {
    return nil, fmt.Errorf("Setup rule must be for each player")
  }
  return &SinglePlayerSetupStep{rule: rule, owner: owner, Done: false}, nil
}

func (step *SinglePlayerSetupStep) Rule() SetupRule {
  return step.rule
}

func (step *SinglePlayerSetupStep) Owner() Player {
  return step.owner
}

type GlobalSetupStep struct {
  rule SetupRule
  Done bool
}

func NewGlobalSetupStep(rule SetupRule) (*GlobalSetupStep, error) {
  if "Once" != rule.Arity {
    return nil, fmt.Errorf("Setup rule must be done once")
  }
  return &GlobalSetupStep{rule: rule, Done: false}, nil
}

func (step *GlobalSetupStep) Rule() SetupRule {
  return step.rule
}

func (step *GlobalSetupStep) Owner() Player {
  return Player("global")
}
