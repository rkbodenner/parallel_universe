package game

import (
  "fmt"
)

type Player string

type Game struct {
  Name string
  SetupRules []SetupRule
  Players []Player
  SetupAssignments map[Player][]SetupStep
  SetupSteps []SetupStep
}

func NewGame(name string, rules []SetupRule, playerCount uint) *Game {
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
    Name: name,
    SetupRules: rules,
    Players: players,
    SetupAssignments: make(map[Player][]SetupStep),
    SetupSteps: setupSteps,
  }
}

func (game *Game) PlayerCount() int {
  return len(game.Players)
}

func (game *Game) completeCurrentSetupStep(player Player) {
  for _,step := range game.SetupAssignments[player] {
    step.SetDone()
  }
}

func (game *Game) IsSetupStepAssigned(needle SetupStep) bool {
  for _,steps := range game.SetupAssignments {
    for _,step := range steps {
      if needle == step {
        return true
      }
    }
  }
  return false
}

func (game *Game) findNextUndoneSetupStep(player Player) (SetupStep, error) {
  for _,step := range game.SetupSteps {
    if !game.IsSetupStepAssigned(step) && step.CanBeOwnedBy(player) && !step.Done() {
      return step, nil
    }
  }
  return nil, fmt.Errorf("No undone steps available for %s", player)
}

func (game *Game) NextStep(player Player) (SetupStep, error) {
  found := false
  for _,p := range game.Players {
    if p == player {
      found = true
    }
  }
  if ! found {
    return nil, fmt.Errorf("No such player: %i", player)
  }

  game.completeCurrentSetupStep(player)
  step,err := game.findNextUndoneSetupStep(player)
  if nil != err {
    return step, err
  }
  game.SetupAssignments[player] = append(game.SetupAssignments[player], step)

  return step, nil
}

func (game *Game) AssignSteps() {
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
}

func (game *Game) PrintSetupSteps() {
  for _,step := range game.SetupSteps {
    fmt.Printf("%s\t%t\n", step.Rule().Description, step.Done())
  }
  fmt.Println("---")
}

func (game *Game) PrintStepAssignments() {
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
}

func (game *Game) PrintSetupRules() {
  for _,r := range game.SetupRules {
    if "Each player" == r.Arity {
      for _,p := range game.Players {
        fmt.Printf("%s\t%s\n", r.Description, p)
      }
    } else {
      fmt.Printf("%s\t%s\n", r.Description, r.Arity)
    }
  }
}

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
