package session

import (
  "encoding/json"
  "strconv"
  "github.com/rkbodenner/parallel_universe/game"
)

type StepAssignments interface {
  Get(*game.Player) (*game.SetupStep,bool)
  Set(*game.Player, *game.SetupStep)

  IsAssigned(*game.SetupStep) bool
  GetAssignee(*game.SetupStep) *game.Player
  GetAssignees(*game.SetupRule) []*game.Player
}

type StepMap struct {
  stepMap map[*game.Player]*game.SetupStep
}

func NewStepMap() *StepMap {
  return &StepMap{make(map[*game.Player]*game.SetupStep)}
}

func (m *StepMap) Get(player *game.Player) (*game.SetupStep, bool) {
  step,ok := m.stepMap[player]
  return step,ok
}

func (m *StepMap) Set(player *game.Player, step *game.SetupStep) {
  m.stepMap[player] = step
}

func (m *StepMap) IsAssigned(step *game.SetupStep) bool {
  assignees := m.GetAssignees(step.Rule)
  for _, assignee := range assignees {
    if step.CanBeOwnedBy(assignee) {
      return true
    }
  }
  return false
}

func (m *StepMap) GetAssignee(s *game.SetupStep) *game.Player {
  for player,step := range m.stepMap {
    if s.Equal(step) {
      return player
    }
  }
  return nil
}

func (m *StepMap) GetAssignees(rule *game.SetupRule) []*game.Player {
  playersAssigned := make([]*game.Player, 0)
  for player,step := range m.stepMap {
    if step.Rule.Description == rule.Description {
      playersAssigned = append(playersAssigned, player)
    }
  }
  return playersAssigned
}

func (m *StepMap) MarshalJSON() ([]byte, error) {
  var mapWithStringIdKeys = make(map[string]*game.SetupStep)
  for key,value := range m.stepMap {
    mapWithStringIdKeys[strconv.Itoa(key.Id)] = value
  }
  return json.Marshal(mapWithStringIdKeys)
}
