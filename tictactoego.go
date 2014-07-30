package main

import (
  . "./tictactoe/"
)

func main() {
  game := GameBoard{ CreateBoard(3) }
  turn := "X"
  DisplayBoard(game.Board)
  game.Move(4, turn)
  DisplayBoard(game.Board)
  SwitchTurn(&turn)
  game.Move(5, turn)
  DisplayBoard(game.Board)
}
