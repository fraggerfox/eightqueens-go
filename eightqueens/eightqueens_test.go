package eightqueens_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/fraggerfox/eightqueens-go/eightqueens"
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
	board eightqueens.Board
}

func (suite *EightQueensSuite) SetupTest() {
	var board eightqueens.Board
	suite.board = board
}

func (suite *EightQueensSuite) TestEmptyBoard() {
	var emptyBoard eightqueens.Board
	assert.ElementsMatch(suite.T(), suite.board, emptyBoard)
}

func (suite *EightQueensSuite) TestIsSafe() {
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
			name:           "Existing queen in a8, c7 and e6, should be safe with b5",
			existingQueens: map[int]int{0: 0, 1: 2, 2: 4},
			row:            3,
			column:         1,
			want:           true,
		},
	}

	for _, tc := range tcs {
		for row, col := range tc.existingQueens {
			suite.board[row][col] = true
		}
		assert.Equalf(suite.T(), tc.want, eightqueens.IsSafe(suite.board, tc.row, tc.column), "%s: expected: %v, got: %v", tc.name, tc.want, !tc.want)
	}
}

func (suite *EightQueensSuite) TestPlaceQueens() {
	solutions := make(eightqueens.Boards, 0)
	solutions = eightqueens.PlaceQueens(suite.board, 0, &solutions)

	expectedSolutionCount := 92
	// This is the first solution discovered using this method.
	/*
	   Q - - - - - - -
	   - - - - Q - - -
	   - - - - - - - Q
	   - - - - - Q - -
	   - - Q - - - - -
	   - - - - - - Q -
	   - Q - - - - - -
	   - - - Q - - - -
	*/
	expectedFirstsolution := eightqueens.Board{
		{true, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, true},
		{false, false, false, false, false, true, false, false},
		{false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, true, false},
		{false, true, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false},
	}

	assert.Equalf(suite.T(), expectedSolutionCount, len(solutions), "PlaceQueens: expected %v, got: %v", expectedSolutionCount, len(solutions))
	assert.Equalf(suite.T(), expectedFirstsolution, solutions[0], "PlaceQueens: expected %v, got: %v", expectedFirstsolution, solutions[0])
}

func TestEightQueensSuite(t *testing.T) {
	suite.Run(t, new(EightQueensSuite))
}
