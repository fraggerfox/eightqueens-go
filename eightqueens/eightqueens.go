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

// IsSafe checks if a row, col combinatation is safe to place a Queen.
func IsSafe(board Board, row, col int) bool {
	return checkRowColumn(board, row, col) && checkDiagonal(board, row, col)
}

// PlaceQueens recursively attempts to place Queens on a chessboard, backtracks in case of failures.
// Bruteforce the solution by exhaustive searching.
func PlaceQueens(board Board, row int, solutions *Boards) Boards {
	if row == BOARD_SIZE {
		*solutions = append(*solutions, board)
		return *solutions
	}

	for col := 0; col < BOARD_SIZE; col++ {
		if IsSafe(board, row, col) {
			board[row][col] = true
			PlaceQueens(board, row+1, solutions)
			board[row][col] = false
		}
	}

	return *solutions
}
