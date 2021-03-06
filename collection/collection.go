package collection

import (
  "github.com/rkbodenner/parallel_universe/game"
)

func NewTicTacToe() *game.Game {
  var setup = []*game.SetupRule{
    game.NewSetupRule("Draw 3x3 grid", "Once"),
    game.NewSetupRule("Choose X or O", "Each player"),
  }

  return game.NewGame("Tic-Tac-Toe", setup, 2, 2)
}

func NewForbiddenIsland() *game.Game {
  var setup = []*game.SetupRule{
    game.NewSetupRule("Create Forbidden Island", "Once"),     //0
    game.NewSetupRule("Place the treasures", "Once"),
    game.NewSetupRule("Divide the cards", "Once"),            //2
    game.NewSetupRule("The island starts to sink", "Once"),   //3
    game.NewSetupRule("The Adventurers appear", "Once"),      //4
    game.NewSetupRule("Place Adventurer pawn", "Each player"),//5
    game.NewSetupRule("Hand out Treasure deck cards", "Once"),//6
    game.NewSetupRule("Set the water level", "Once"),
  }
  setup[0].Details = "Shuffle the 24 Island tiles and randomly place them colorful-side-up into a 4x4 grid, then place 2 tiles next to each of the two middle tiles on every side of the square. Leave a small gap between the tiles."
  setup[1].Details = "Place the 4 treasure figurines--The Earth Stone, The Statue of the Wind, The Crystal of Fire, and The Ocean's Chalice--anywhere off to the side of the island"
  setup[2].Details = "Separate the cards into three decks according to the card backs: Flood deck (blue back), Treasure deck (red), and Adventurer cards (6 cards)"
  setup[3].Details = "Shuffle the Flood deck and place it face down on one side of the island, forming the Flood draw pile. Draw the top 6 cards (1 at a time) and place them face up next to the draw pile, forming the Flood discard pile. For each card drawn, flip the corresponding Island tile over to its flooded (blue & white) side."
  setup[4].Details = "Shuffle the Adventurer cards and randomly deal 1 to each player. Put undealt cards and their matching pawns back in the box."
  setup[5].Details = "Take the pawn matching the color of your Adventurer card and place it on the corresponding Island tile. Look for the matching pawn icon in the lower right corner of the Gates and Fools' Landing tiles. It's OK to start on a flooded tile."
  setup[6].Details = "Shuffle the Treasure deck and deal 2 cards to each player. Place your cards face up in front of you. If anyone gets a Waters Rise! card, give them a replacement and shuffle Waters Rise! back into the deck. Place the Treasure deck face down by one side of the island."
  setup[7].Details = "Place the Water Level slider on the left side of the Water Meter board and set it to the difficulty level of your choice"

  setup[3].Dependencies = []*game.SetupRule{setup[0], setup[2]}
  setup[4].Dependencies = []*game.SetupRule{setup[2]}
  setup[5].Dependencies = []*game.SetupRule{setup[4]}
  setup[6].Dependencies = []*game.SetupRule{setup[2]}

  return game.NewGame("Forbidden Island", setup, 2, 4)
}

func NewOraEtLaboraShortMultiplayer() *game.Game {
  var setup = []*game.SetupRule{
    game.NewSetupRule("Choose game board for short 3-4 player game", "Once"),

    game.NewSetupRule("Attach production wheel to game board", "Once"),
    game.NewSetupRule("Place 7 wooden goods indicators on game board", "Once"),
    game.NewSetupRule("Sort the building cards", "Once"),
    game.NewSetupRule("Place the start buildings", "Once"),
    game.NewSetupRule("Place the A, B, C, D buildings", "Once"),

    game.NewSetupRule("Place the black stone goods indicator", "Once"),
    game.NewSetupRule("Place the purple grapes goods indicator", "Once"),
    game.NewSetupRule("Take a heartland landscape board", "Each player"),
    game.NewSetupRule("Place moor and forest cards on landscape board", "Each player"),
    game.NewSetupRule("Choose a color", "Each player"),

    game.NewSetupRule("Take 1 prior and 1 lay brother of your color", "Each player"),
    game.NewSetupRule("Take 8 settlement cards of your color", "Each player"),
    game.NewSetupRule("Take 1 of each of the 6 starting goods", "Each player"),
    game.NewSetupRule("Remove unused tiles", "Once"),
    game.NewSetupRule("Sort districts and plots by cost", "Once"),
  }

  setup[0].Details = "The correct board will have an icon with two players, in the center on the reverse side. Place the board in the middle of the table."

  setup[1].Details = "Side showing 0/2/3/4/... should face up. Orient the wheel so that the beam points to the bible symbol. You can unscrew the wheel from the board with a fingernail."
  setup[1].Dependencies = []*game.SetupRule{setup[0]}
  setup[2].Details = "Place onto the board where the production wheel indicates 0 (clay, coins, grain, livestock, wood, peat, joker)"
  setup[2].Dependencies = []*game.SetupRule{setup[1]}
  // TODO: Player number variation
  setup[3].Details = "3-player game: Remove the cards with a 4 or a 3+ in the lower right corner. 4-player game: Remove the cards with a 4 in the lower right corner. Turn each card so that the chosen country variant (France or Ireland) faces up. Sort the buildings into stacks by the letter or bible symbol in the middle left of the card."
  setup[4].Details = "Start buildings have a bible symbol in the middle left of the card. Place the stack anywhere all players can see them."
  setup[4].Dependencies = []*game.SetupRule{setup[3]}
  setup[5].Details = "Place each stack next to the matching blue A, B, C, D symbol on the edge of the game board."
  setup[5].Dependencies = []*game.SetupRule{setup[1], setup[3]}

  setup[6].Details = "Place it at the position indicated by the matching symbol on the edge of the game board."
  setup[6].Dependencies = []*game.SetupRule{setup[1]}
  // TODO: Variant
  setup[7].Details = "Only if playing the France variant. Place it at the position indicated by the matching symbol on the edge of the game board."
  setup[7].Dependencies = []*game.SetupRule{setup[1]}
  setup[9].Details = "Place 1 moor and 2 forest. Leave the left-most two spaces empty on the upper row of the landscape board."
  setup[9].Dependencies = []*game.SetupRule{setup[8]}

  setup[11].Dependencies = []*game.SetupRule{setup[10]}
  setup[12].Details = "Stack buildings marked A, B, C, D under the respective piles of building cards next to the board."
  setup[12].Dependencies = []*game.SetupRule{setup[10]}
  setup[13].Details = "Clay, coin, grain, livestock, wood, peat. Place them right-side up."
  // TODO: Variant
  setup[14].Details = "France variant: Remove malt/beer. Ireland variant: Remove flour/bread and grapes/wine."
  setup[15].Details = "Lowest cost on top."

  return game.NewGame("Ora et Labora", setup, 3, 4)
}

type Collection struct {
  Games []*game.Game
}

func NewCollection() *Collection {
  return &Collection{
    []*game.Game{
      NewTicTacToe(),
      NewForbiddenIsland(),
      NewOraEtLaboraShortMultiplayer(),
    },
  }
}
