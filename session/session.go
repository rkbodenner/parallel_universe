package session

import (
  "fmt"
  "github.com/rkbodenner/parallel_universe/game"
)

type Session struct {
  Game *game.Game
  Players []game.Player
  SetupAssignments map[game.Player][]game.SetupStep
  SetupSteps []game.SetupStep
}

func NewSession(g *game.Game, playerCount uint) *Session {
  players := make([]game.Player, playerCount)
  for i := range players {
    players[i] = (game.Player)(fmt.Sprintf("Player %d", i+1))
  }

  setupSteps := make([]game.SetupStep, 0)
  for _,rule := range g.SetupRules {
    if "Once" == rule.Arity {
      step, err := game.NewGlobalSetupStep(rule)
      if nil != err {
        fmt.Println(err)
      }
      setupSteps = append(setupSteps, step)
    } else if "Each player" == rule.Arity {
      for _,p := range players {
        step, err := game.NewSinglePlayerSetupStep(rule, p)
        if nil != err {
          fmt.Println(err)
        }
        setupSteps = append(setupSteps, step)
      }
    }
  }
  return &Session{
    Game: g,
    Players: players,
    SetupAssignments: make(map[game.Player][]game.SetupStep),
    SetupSteps: setupSteps,
  }
}

func (session *Session) PlayerCount() int {
  return len(session.Players)
}

func (session *Session) completeCurrentSetupStep(player game.Player) {
  for _,step := range session.SetupAssignments[player] {
    step.SetDone()
  }
}

func (session *Session) IsSetupStepAssigned(needle game.SetupStep) bool {
  for _,steps := range session.SetupAssignments {
    for _,step := range steps {
      if needle == step {
        return true
      }
    }
  }
  return false
}

func (session *Session) findNextUndoneSetupStep(player game.Player) (game.SetupStep, error) {
  for _,step := range session.SetupSteps {
    if !session.IsSetupStepAssigned(step) && step.CanBeOwnedBy(player) && !step.Done() {
      return step, nil
    }
  }
  return nil, fmt.Errorf("No undone steps available for %s", player)
}

func (session *Session) NextStep(player game.Player) (game.SetupStep, error) {
  found := false
  for _,p := range session.Players {
    if p == player {
      found = true
    }
  }
  if ! found {
    return nil, fmt.Errorf("No such player: %i", player)
  }

  session.completeCurrentSetupStep(player)
  step,err := session.findNextUndoneSetupStep(player)
  if nil != err {
    return step, err
  }
  session.SetupAssignments[player] = append(session.SetupAssignments[player], step)

  return step, nil
}

func (session *Session) PrintSetupSteps() {
  for _,step := range session.SetupSteps {
    fmt.Printf("%s\t%t\n", step.Rule().Description, step.Done())
  }
  fmt.Println("---")
}

func (session *Session) PrintStepAssignments() {
  for _,player := range session.Players {
    fmt.Printf("-- %d steps for %s\n", len(session.SetupAssignments[player]), player)
    for _,step := range session.SetupAssignments[player] {
      fmt.Printf("%s", step.Rule().Description)
      if "Each player" == step.Rule().Arity {
        fmt.Println(" *")
      } else {
        fmt.Println()
      }
    }
  }
}

