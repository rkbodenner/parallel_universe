package session

import (
  "fmt"
  "github.com/rkbodenner/parallel_universe/game"
)

type Session struct {
  Id uint `json:"id"`
  Game *game.Game
  Players []*game.Player
  SetupAssignments StepAssignments
  SetupSteps []game.SetupStep
  freeSetupSteps map[game.SetupStep]bool
}

func NewSession(g *game.Game, players []*game.Player) *Session {
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

  freeSetupSteps := make(map[game.SetupStep]bool)
  for _,step := range setupSteps {
    freeSetupSteps[step] = true
  }

  return &Session{
    Game: g,
    Players: players,
    SetupAssignments: NewStepMap(),
    SetupSteps: setupSteps,
    freeSetupSteps: freeSetupSteps,
  }
}

func (session *Session) PlayerCount() int {
  return len(session.Players)
}

func (session *Session) findNextUndoneSetupStep(player *game.Player) (game.SetupStep, error) {
  for step,_ := range session.freeSetupSteps {
    if step.CanBeOwnedBy(player) && !step.IsDone() {
      return step, nil
    }
  }
  return nil, fmt.Errorf("No undone steps available for %s", player.Name)
}

func (session *Session) Step(player *game.Player) game.SetupStep {
  step,assigned := session.SetupAssignments.Get(player)
  if !assigned || (assigned && step.IsDone()) {
    nextStep,error := session.findNextUndoneSetupStep(player)
    if ( error != nil ) {
      fmt.Println(error.Error())
      return step
    }
    session.SetupAssignments.Set(player, nextStep)
    return nextStep
  }
  return step
}
