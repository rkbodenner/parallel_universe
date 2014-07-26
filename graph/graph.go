/*

Print graphviz dot files, for visualizing game setup steps and their dependencies

go run graph.go | dot -Tpng > game.png

*/

package main

import (
  "fmt"
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
  PrintDot(collection.NewForbiddenIsland())
}
