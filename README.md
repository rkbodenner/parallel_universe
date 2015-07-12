parallel_universe makes the setup of complex boardgames fast and easy.

It will hand out setup tasks to each player to minimize setup time, so that none of the players are sitting around wondering what to do. Eventually, players will be able to use a web interface on their mobile devices to work through their setup tasks.

Why?
----
Eclipse takes forever to set up. So does Mansions of Madness. If only all six players could start grabbing the cardboard bits and get them in the right places without waiting on the one person holding the rulebook, we could get down to the fun part sooner.

What
----
This library models tabletop game setup and includes the following:
- Types for games and their setup rules, game sessions, players, assignment of setup steps to players
- Dependencies between steps, e.g., you have to set up the board before you put the pieces on it
- Handing out steps to a variable number of players
- Visualizing the tree of steps using graphviz. See `graph/graph.go`.
- Rules for a few games. See `collection/collection.go`.

It doesn't yet include:
- Scenarios or variants of play for a single game. For example, if the setup varies a lot depending on the player count (e.g., the dummy player in Tokaido, or solo variants), you'll just have to write that out in the setup rule description, or create a separate Game. Similarly for the scenarios in Mansions of Madness or Descent.

It won't include:
- Storing stuff in a database or exposing it with a web API. That's what https://github.com/rkbodenner/meeple_mover does.

Using this library
----
* [Install Go](http://golang.org/doc/install)
* In your GOPATH, `go get github.com/rkbodenner/parallel_universe`

TODO
----
https://gist.github.com/rkbodenner/8a970d33a9c69b1471cd
