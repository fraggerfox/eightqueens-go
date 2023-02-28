package eightqueens

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

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
	assert.PanicsWithValue(suite.T(), "not implemented", func() { checkRowColumn(suite.board, 0, 0) })
}

func (suite *EightQueensSuite) TestCheckDiagonal() {
	assert.PanicsWithValue(suite.T(), "not implemented", func() { checkDiagonal(suite.board, 0, 0) })
}

func (suite *EightQueensSuite) TestPlaceQueen() {
	assert.PanicsWithValue(suite.T(), "not implemented", func() { placeQueen(suite.board, 0, 0) })
}

func TestEightQueensSuite(t *testing.T) {
	suite.Run(t, new(EightQueensSuite))
}
