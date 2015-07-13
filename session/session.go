package session

import (
  "errors"
  "fmt"
  "github.com/rkbodenner/parallel_universe/game"
)

type Session struct {
  Id uint `json:"id"`
  Game *game.Game
  Players []*game.Player
  SetupAssignments StepAssignments
  SetupSteps []*game.SetupStep
}

func newStepAssignments() StepAssignments {
  return NewStepPlayerIdMap()
}

func NewEmptySession() (*Session) {
  return &Session{
    Game: nil,
    Players: make([]*game.Player, 0),
    SetupAssignments: newStepAssignments(),
    SetupSteps: make([]*game.SetupStep, 0),
  }
}

func NewSession(g *game.Game, players []*game.Player) (*Session, error) {
  if nil == g {
    return nil, errors.New("Session must have a non-nil Game")
  }
  if len(players) < g.MinPlayers {
    return nil, errors.New(fmt.Sprintf("Game requires at least %d players, but %d were given", g.MinPlayers, len(players)))
  }
  if len(players) > g.MaxPlayers {
    return nil, errors.New(fmt.Sprintf("Game requires at most %d players, but %d were given", g.MaxPlayers, len(players)))
  }

  setupSteps := make([]*game.SetupStep, 0)
  for _,rule := range g.SetupRules {
    if "Once" == rule.Arity {
      step, err := game.NewGlobalSetupStep(rule)
      if nil != err {
        return nil, err
      }
      setupSteps = append(setupSteps, step)
    } else if "Each player" == rule.Arity {
      for _,p := range players {
        step, err := game.NewSinglePlayerSetupStep(rule, p)
        if nil != err {
          return nil, err
        }
        setupSteps = append(setupSteps, step)
      }
    }
  }

  return &Session{
    Game: g,
    Players: players,
    SetupAssignments: newStepAssignments(),
    SetupSteps: setupSteps,
  }, nil
}

func (session *Session) IsRuleDone(rule *game.SetupRule) bool {
  for _,step := range session.SetupSteps {
    if step.Rule.Description == rule.Description && !step.Done {
      return false
    }
  }
  return true
}

func (session *Session) AreStepDependenciesDone(step *game.SetupStep) bool {
  for _,dep := range step.Rule.Dependencies {
    if !session.IsRuleDone(dep) {
      return false
    }
  }
  return true
}

func (session *Session) findNextUndoneSetupStep(player *game.Player) (*game.SetupStep, error) {
  var conditions [4]bool
  for _, step := range session.SetupSteps {
    conditions = [4]bool{
      !session.SetupAssignments.IsAssigned(step),
      step.CanBeOwnedBy(player),
      !step.Done,
      session.AreStepDependenciesDone(step),
    }
    fmt.Printf("Step %s available? [%t %t %t %t]\n", step.Rule.Description, conditions[0], conditions[1], conditions[2], conditions[3])
    if conditions[0] && conditions[1] && conditions[2] && conditions[3] {
      return step, nil
    }
  }
  return nil, fmt.Errorf("No undone steps available for %s", player.Name)
}

func (session *Session) Step(player *game.Player) *game.SetupStep {
  step,hasAssignment := session.SetupAssignments.Get(player)
  if !hasAssignment || (hasAssignment && step.Done) {
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

func (session *Session) StepAllPlayers() error {
  for _, player := range session.Players {
    session.Step(player)
  }
  return nil
}

func (session *Session) Print() {
  for _,step := range session.SetupSteps {
    ownerName := ""
    if nil != step.Owner {
      ownerName = step.Owner.Name
    }
    assigneeName := ""
    if assignee := session.SetupAssignments.GetAssignee(step); assignee != nil {
      assigneeName = assignee.Name
    }
    fmt.Printf("%s (%s)<-(%s) %t\n", step, ownerName, assigneeName, step.Done)
  }
}
