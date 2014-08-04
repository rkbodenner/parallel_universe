/*

Print graphviz dot files, for visualizing game setup steps and their dependencies

go run graph.go | dot -Tpng > game.png

*/

package main

import (
  "database/sql"
  "flag"
  "fmt"
  "os"
  _ "github.com/lib/pq"
  "github.com/rkbodenner/meeple_mover/record"
  "github.com/rkbodenner/parallel_universe/collection"
  "github.com/rkbodenner/parallel_universe/game"
)

func PrintDot(g *game.Game) {
  fmt.Println("digraph {")
  for _, rule := range g.SetupRules {
    fmt.Printf("  \"%s\";\n", rule.Description)
    for _, dep := range rule.Dependencies {
      fmt.Printf("  \"%s\" -> \"%s\";", dep.Description, rule.Description)
    }
  }
  fmt.Println("}")
}

func main() {
  var gameName string
  flag.StringVar(&gameName, "game", "", "Name of the game to print a dot file for")
  var useTestDB bool
  flag.BoolVar(&useTestDB, "test-db", false, "Load games from the test database")
  var verbose bool
  flag.BoolVar(&verbose, "verbose", false, "Print messages about progress to stdout")

  flag.Parse()
  if "" == gameName {
    flag.Usage()
    os.Exit(1)
  }

  if verbose {
    fmt.Printf("Searching for %s...\n", gameName)
  }
  var game *game.Game

  if useTestDB {
    if verbose {
      fmt.Println("Loading from the test database...")
    }

    connectString := fmt.Sprintf("user=ralph dbname=meeple_mover_test sslmode=disable")
    db, err := sql.Open("postgres", connectString)
    if nil != err {
      fmt.Print(err)
    }
    defer db.Close()

    rec := record.NewEmptyGameRecord()
    // FIXME: This prints messages to stdout, which breaks redirecting output straight to a dot file
    err = rec.FindByName(db, gameName)
    if nil != err {
      fmt.Println(err)
    } else {
      game = rec.Game
    }
  } else {
    if verbose {
      fmt.Println("Loading from the collection...")
    }

    shelf := collection.NewCollection().Games
    for _, g := range shelf {
      if gameName == g.Name {
        if verbose {
          fmt.Printf("Found %s in the collection\n", gameName)
        }
        game = g
        break
      }
    }
  }

  if nil != game {
    PrintDot(game)
  } else {
    fmt.Fprintf(os.Stderr, "Error: No such game")
  }
}
