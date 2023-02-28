package eightqueens

const BOARD_SIZE = 8

type Board [BOARD_SIZE][BOARD_SIZE]bool
type Boards []Board

func checkRowColumn(board Board, row, col int) bool {
	panic("not implemented")
}

func checkDiagonal(board Board, row, col int) bool {
	panic("not implemented")
}

func IsSafe(board Board, row, col int) bool {
	panic("not implemented")
}

func placeQueen(board Board, row, col int) bool {
	panic("not implemented")
}

func PlaceQueens(board Board, row int) Boards {
	return nil
}
