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
  if nil != g {
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

func (session *Session) IsRuleDone(rule *game.SetupRule) bool {
  for _,step := range session.SetupSteps {
    if step.GetRule().Description == rule.Description && !step.IsDone() {
      return false
    }
  }
  return true
}

func (session *Session) AreStepDependenciesDone(step game.SetupStep) bool {
  for _,dep := range step.GetRule().Dependencies {
    if !session.IsRuleDone(dep) {
      return false
    }
  }
  return true
}

func (session *Session) findNextUndoneSetupStep(player *game.Player) (game.SetupStep, error) {
  for step,_ := range session.freeSetupSteps {
    if step.CanBeOwnedBy(player) && !step.IsDone() && session.AreStepDependenciesDone(step) {
      return step, nil
    }
  }
  return nil, fmt.Errorf("No undone steps available for %s", player.Name)
}

func (session *Session) Step(player *game.Player) game.SetupStep {
  step,hasAssignment := session.SetupAssignments.Get(player)
  if !hasAssignment || (hasAssignment && step.IsDone()) {
    nextStep,error := session.findNextUndoneSetupStep(player)
    if ( error != nil ) {
      fmt.Println(error.Error())
      return step
    }
    session.SetupAssignments.Set(player, nextStep)
    delete(session.freeSetupSteps, nextStep)
    return nextStep
  }
  return step
}

func (session *Session) Print() {
  for _,step := range session.SetupSteps {
    ownerName := ""
    if nil != step.GetOwner() {
      ownerName = step.GetOwner().Name
    }
    fmt.Printf("%s (%s) %t\n", step, ownerName, step.IsDone())
  }
}
