package eightqueens

const BOARD_SIZE = 8

type Board [BOARD_SIZE][BOARD_SIZE]bool
type Boards []Board

// checkRowColumn returns true if it is safe to place the queen
// in given position, false if a conflict is found.
func checkRowColumn(board Board, row, col int) bool {
	for i := 0; i < BOARD_SIZE; i++ {
		if board[row][i] || board[i][col] {
			return false
		}
	}

	return true
}

// checkDiagonal returns true if it is safe to place the queen
// in given position, false if a conflict is found.
func checkDiagonal(board Board, row, col int) bool {
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < BOARD_SIZE; i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}

	return true
}

func IsSafe(board Board, row, col int) bool {
	return checkDiagonal(board, row, col) && checkRowColumn(board, row, col)
}

func placeQueen(board Board, row, col int) bool {
	panic("not implemented")
}

func PlaceQueens(board Board, row int) Boards {
	return nil
}
