package tictactoe

import (
	"fmt"
	"strings"
)

func InputGridSize() (dim int) {
	fmt.Printf("What size grid would you like to play on?? Please enter only one dimension (ie 3, 4, or 5): ")
	_, err := fmt.Scanf("%d", &dim)
	if err != nil {
	}
	return
}

func InputMove(g *GameBoard) {
	for g.Win() == false {
		if g.Turn == "X" {
			var move int
			fmt.Println("Make your move: ")
			_, err := fmt.Scanf("%d", &move)
			if err != nil || ValidMove(g.RemainingIndices(), move-1) == false {
				fmt.Println("Please Enter A Valid Move")
			} else {
				g.Move(move-1, g.Turn)
			}
		} else {
			g.Move(g.AlphaBetaBestMove(), g.Turn)
		}
		DisplayBoard(g.Board)
		DisplayRemainingMoves(g)
		DisplayTurn(g)
		if len(g.RemainingIndices()) == 0 {
			break
		}
	}
}

func InputPlayAgain() {
	var yesOrNo string
	fmt.Printf("Would you like to play again?? ")
	_, err := fmt.Scanf("%s", &yesOrNo)
	if err != nil {
	} else {
		if strings.ToUpper(yesOrNo) == "YES" || strings.ToUpper(yesOrNo) == "Y" {
			Play()
		}
	}
}
