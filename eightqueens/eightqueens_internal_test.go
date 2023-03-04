package eightqueens

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Reference for tests using algebraic notation
// https://en.wikipedia.org/wiki/Algebraic_notation_(chess)
//
// + (0)(1)(2)(3)(4)(5)(6)(7)
// 8 |__|##|__|##|__|##|__|##| (0)
// 7 |##|__|##|__|##|__|##|__| (1)
// 6 |__|##|__|##|__|##|__|##| (2)
// 5 |##|__|##|__|##|__|##|__| (3)
// 4 |__|##|__|##|__|##|__|##| (4)
// 3 |##|__|##|__|##|__|##|__| (5)
// 2 |__|##|__|##|__|##|__|##| (6)
// 1 |##|__|##|__|##|__|##|__| (7)
// +  a  b  c  d  e  f  g  h
//
// For example board[0][0] is a8 and board[3][6] is g5 an so on.
type EightQueensSuite struct {
	suite.Suite
	board Board
}

func (suite *EightQueensSuite) SetupTest() {
	var board Board
	suite.board = board
}

func (suite *EightQueensSuite) TestEmptyBoard() {
	var emptyBoard Board
	assert.ElementsMatch(suite.T(), suite.board, emptyBoard)
}

func (suite *EightQueensSuite) TestCheckRowColumn() {
	tcs := []struct {
		name           string
		existingQueens map[int]int
		row            int
		column         int
		want           bool
	}{
		{
			name:           "Check for conflicts in an empty board, a2",
			existingQueens: map[int]int{},
			row:            6,
			column:         0,
			want:           true,
		},
		{
			name:           "Existing queen in a8, should conflict with b8",
			existingQueens: map[int]int{0: 0},
			row:            0,
			column:         1,
			want:           false,
		},
		{
			name:           "Existing queen in a8 and c7, should conflict with c1",
			existingQueens: map[int]int{0: 0, 1: 2},
			row:            7,
			column:         2,
			want:           false,
		},
		{
			name:           "Existing queen in a8, c7 and e6, should be safe with d5",
			existingQueens: map[int]int{0: 0, 1: 2, 2: 4},
			row:            3,
			column:         3,
			want:           true,
		},
	}

	for _, tc := range tcs {
		for row, col := range tc.existingQueens {
			suite.board[row][col] = true
		}
		assert.Equalf(suite.T(), tc.want, checkRowColumn(suite.board, tc.row, tc.column), "%s: expected: %v, got: %v", tc.name, tc.want, !tc.want)
	}
}

func (suite *EightQueensSuite) TestCheckDiagonal() {
	tcs := []struct {
		name           string
		existingQueens map[int]int
		row            int
		column         int
		want           bool
	}{
		{
			name:           "Check for conflicts in an empty board, a2",
			existingQueens: map[int]int{},
			row:            6,
			column:         0,
			want:           true,
		},
		{
			name:           "Existing queen in a8, should conflict with h1",
			existingQueens: map[int]int{0: 0},
			row:            7,
			column:         7,
			want:           false,
		},
		{
			name:           "Existing queen in a8 and c7, should conflict with f4",
			existingQueens: map[int]int{0: 0, 1: 2},
			row:            4,
			column:         5,
			want:           false,
		},
		{
			name:           "Existing queen in a8, c7 and b6, should be safe with g5",
			existingQueens: map[int]int{0: 0, 1: 2, 2: 1},
			row:            3,
			column:         6,
			want:           true,
		},
	}

	for _, tc := range tcs {
		for row, col := range tc.existingQueens {
			suite.board[row][col] = true
		}
		assert.Equalf(suite.T(), tc.want, checkDiagonal(suite.board, tc.row, tc.column), "%s: expected: %v, got: %v", tc.name, tc.want, !tc.want)
	}
}

func TestEightQueensSuite(t *testing.T) {
	suite.Run(t, new(EightQueensSuite))
}
