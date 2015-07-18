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
        step, err := game.NewPerPlayerSetupStep(rule, p)
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
    if step.Rule.Equal(rule) && !step.Done {
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

func (session *Session) findNextUndoneSetupStep(player *game.Player) (*game.SetupStep) {
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
      return step
    }
  }
  return nil
}

func (session *Session) Step(player *game.Player) *game.SetupStep {
  step,hasAssignment := session.SetupAssignments.Get(player)
  if !hasAssignment || (hasAssignment && step.Done) {
    nextStep := session.findNextUndoneSetupStep(player)
    if nil != nextStep {
      session.SetupAssignments.Set(player, nextStep)
      step = nextStep
    }
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
    fmt.Println(session.StepWithAssigneeString(step))
  }
}

func (session *Session) StepWithAssigneeString(step *game.SetupStep) string {
  assigneeName := ""
  if assignee := session.SetupAssignments.GetAssignee(step); assignee != nil {
    assigneeName = assignee.Name
  }
  return fmt.Sprintf("%s->(%s) %t", step, assigneeName, step.Done)
}
