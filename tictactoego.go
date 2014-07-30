package main

import (
  . "./tictactoe/"
)

func main() {
  game := GameBoard{ CreateBoard(3), "X" }
  DisplayBoard(game.Board)
  game.Move(4, "X")
  DisplayBoard(game.Board)
}
