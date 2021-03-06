package tictactoe

import (
	"fmt"
	"math"
	"strings"
)

func Play() {
	dim := InputGridSize()
	results := make(map[string]int)
	game := GameBoard{CreateBoard(dim), "X", 0, results}
	DisplayBoard(game.Board)
	DisplayRemainingMoves(&game)
	DisplayTurn(&game)
	InputMove(&game)
	DisplayWinner(game)
	InputPlayAgain()
}

func horizontalBars(dimension int) {
	fmt.Println("--" + strings.Repeat("-", (dimension*2)) + "-")
}

func DisplayBoard(board []string) {
	dimension := int(math.Sqrt(float64(len(board))))
	horizontalBars(dimension)
	for i := 0; i < len(board); {
		rowEnd := i + dimension
		fmt.Println("|", strings.Join(board[i:rowEnd], " "), "|")
		i += dimension
	}
	horizontalBars(dimension)
}

func DisplayRemainingMoves(g *GameBoard) {
	fmt.Println("Remaining Moves:", AddOneToSliceValues(g.RemainingIndices()))
}

func DisplayTurn(g *GameBoard) {
	fmt.Println("Current Turn:", g.Turn)
}

func DisplayWinner(g GameBoard) {
	if g.Win() == true {
		fmt.Printf("Congratulations %v!! You've won!!\n", g.PreviousTurn())
	} else {
		fmt.Println("It's a draw.")
	}
}
