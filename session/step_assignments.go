package session

import (
  "encoding/json"
  "strconv"
  "github.com/rkbodenner/parallel_universe/game"
)

type StepAssignments interface {
  Get(*game.Player) (game.SetupStep,bool)
  Set(*game.Player, game.SetupStep)
}

type StepMap struct {
  stepMap map[*game.Player]game.SetupStep
}

func NewStepMap() *StepMap {
  return &StepMap{make(map[*game.Player]game.SetupStep)}
}

func (m *StepMap) Get(player *game.Player) (game.SetupStep, bool) {
  step,ok := m.stepMap[player]
  return step,ok
}

func (m *StepMap) Set(player *game.Player, step game.SetupStep) {
  m.stepMap[player] = step
}

func (m *StepMap) MarshalJSON() ([]byte, error) {
  var mapWithStringIdKeys = make(map[string]game.SetupStep)
  for key,value := range m.stepMap {
    mapWithStringIdKeys[strconv.Itoa(key.Id)] = value
  }
  return json.Marshal(mapWithStringIdKeys)
}