/*

  Test program that sets up a 2-player game

*/
package main

import (
  "fmt"
  "os"
  "github.com/rkbodenner/parallel_universe/collection"
  "github.com/rkbodenner/parallel_universe/game"
  "github.com/rkbodenner/parallel_universe/session"
)

func step(session *session.Session, player *game.Player) {
  step := session.Step(player)
  step.Finish()
  fmt.Println(session.StepWithAssigneeString(step))
}

func printSetupRules(game *game.Game)  {
  for _,r := range game.SetupRules {
    fmt.Printf("%s\t%s\n", r.Description, r.Arity)
  }
}

func main() {
  var players = []*game.Player{
    &game.Player{1, "Player One"},
    &game.Player{2, "Player Two"},
  }
  game := collection.NewForbiddenIsland()
  session, err := session.NewSession(game, players)
  if nil != err {
    fmt.Printf("Error creating game session: %s", err)
    os.Exit(1)
  }

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

  fmt.Println("== Final session state ==========================================")
  session.Print()
  fmt.Println()

  fmt.Println("== Setup rules ==================================================")
  printSetupRules(game)
}
