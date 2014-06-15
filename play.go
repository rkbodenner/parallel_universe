package main

import (
  "fmt"
  "github.com/rkbodenner/parallel_universe/collection"
  "github.com/rkbodenner/parallel_universe/game"
  "github.com/rkbodenner/parallel_universe/session"
)

func step(session *session.Session, player *game.Player) {
  fmt.Printf("%s\t(%s)\n", session.Step(player), player.Name)
  session.Step(player).Finish()
}

func main() {
  var players = []*game.Player{
    &game.Player{1, "Player One"},
    &game.Player{2, "Player Two"},
  }
  game := collection.NewForbiddenIsland()
  session := session.NewSession(game, players)

  fmt.Println("game on")

  step(session, players[0])
  step(session, players[0])
  step(session, players[1])
  step(session, players[0])
  step(session, players[1])
  step(session, players[0])
  step(session, players[1])
  step(session, players[1])
  step(session, players[0])

  fmt.Println()
  game.PrintSetupRules()
}
