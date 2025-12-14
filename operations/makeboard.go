package operations

import "math"

func MakeBoard(tetrominoes [][]string) [][]rune {
	n := len(tetrominoes)
	size := int(math.Round(math.Sqrt(float64(n * 4))))
	size++
	board := make([][]rune, size)
	for i := range board {
		board[i] = make([]rune, size)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	return board
}
