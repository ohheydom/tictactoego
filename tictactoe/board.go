package tictactoe

type GameBoard struct {
	Board  []string
	Turn   string
	Count  int
	Result map[string]int
}

func CreateBoard(size int) []string {
	dim := size * size
	board := make([]string, dim)
	for i := 0; i < dim; i++ {
		board[i] = "-"
	}
	return board
}

func (g *GameBoard) Move(ind int, turn string) {
	if g.Board[ind] == "-" {
		g.Board[ind] = turn
		g.Count++
		g.SwitchTurn()
	}
}

func (g GameBoard) PreviousTurn() string {
	if g.Turn == "X" {
		return "O"
	}
	return "X"
}

func (g GameBoard) RemainingIndices() (remainingIndices []int) {
	remainingIndices = make([]int, 0, len(g.Board))
	for i, value := range g.Board {
		if value == "-" {
			remainingIndices = append(remainingIndices, i)
		}
	}
	return
}

func (g *GameBoard) SwitchTurn() {
	if g.Turn == "X" {
		g.Turn = "O"
	} else {
		g.Turn = "X"
	}
}

func (g *GameBoard) UndoMove(ind int) {
	g.Board[ind] = "-"
	g.Count--
	g.SwitchTurn()
}

func ValidMove(remainingIndices []int, move int) bool {
	n := len(remainingIndices) / 2
	if remainingIndices[n] == move {
		return true
	}
	if remainingIndices[n] > move {
		for _, validMove := range remainingIndices[:n] {
			if validMove == move {
				return true
			}
		}
	} else {
		for _, validMove := range remainingIndices[n:] {
			if validMove == move {
				return true
			}
		}
	}
	return false
}
