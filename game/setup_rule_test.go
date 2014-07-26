package game

import (
  "testing"
)

func TestSetupStepEqual(t *testing.T) {
  r1 := SetupRule{}
  r1.Id = 42
  r2 := SetupRule{}
  r2.Id = r1.Id

  s1 := &SetupStep{&r1, nil, false}
  s2 := &SetupStep{&r2, nil, false}
  if ! s1.Equal(s2) {
    t.Fatal("Steps with the same rule by ID and no owner must be equal")
  }
  if s1.Equal(s2) != s2.Equal(s1) {
    t.Fatal("Equality must be commutative")
  }

  p1 := Player{Id: 37, Name: "Alice"}
  p2 := Player{Id: 37, Name: "Bob"}
  s1.Owner = &p1
  s2.Owner = &p2
  if ! s1.Equal(s2) {
    t.Fatal("Steps with the same rule by ID and same owner by ID must be equal")
  }
  if s1.Equal(s2) != s2.Equal(s1) {
    t.Fatal("Equality must be commutative")
  }
}