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


type StepPlayerIdMap struct {
  stepMap map[int]*game.SetupStep
  playerMap map[int]*game.Player
}

func NewStepPlayerIdMap() *StepPlayerIdMap {
  return &StepPlayerIdMap{make(map[int]*game.SetupStep), make(map[int]*game.Player)}
}

func (m *StepPlayerIdMap) Get(player *game.Player) (*game.SetupStep, bool) {
  step, ok := m.stepMap[player.Id]
  return step, ok
}

func (m *StepPlayerIdMap) Set(player *game.Player, step *game.SetupStep) {
  m.stepMap[player.Id] = step
  m.playerMap[player.Id] = player
}

func (m *StepPlayerIdMap) IsAssigned(step *game.SetupStep) bool {
  return nil != m.GetAssignee(step)
}

func (m *StepPlayerIdMap) GetAssignee(s *game.SetupStep) *game.Player {
  for playerId, step := range m.stepMap {
    if nil != step && s.Equal(step) {
      return m.playerMap[playerId]
    }
  }
  return nil
}

func (m *StepPlayerIdMap) GetAssignees(rule *game.SetupRule) []*game.Player {
  assignees := make([]*game.Player, 0)
  for playerId, step := range m.stepMap {
    if nil != step && step.Rule.Equal(rule) {
      assignees = append(assignees, m.playerMap[playerId])
    }
  }
  return assignees
}

func (m *StepPlayerIdMap) MarshalJSON() ([]byte, error) {
  var mapWithStringIdKeys = make(map[string]*game.SetupStep)
  for key, value := range m.stepMap {
    mapWithStringIdKeys[strconv.Itoa(key)] = value
  }
  return json.Marshal(mapWithStringIdKeys)
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
    if step.Rule.Equal(rule) {
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
