package main

import (
  . "./tictactoe/"
)

func main() {
  dim := GridSizeMessage()
  game := GameBoard{CreateBoard(dim), "X"}
  DisplayBoard(game.Board)
  DisplayTurn(game)
  LoopThroughMoves(game)
}
