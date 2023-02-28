package eightqueens_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/fraggerfox/eightqueens-go/eightqueens"
)

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
	assert.PanicsWithValue(suite.T(), "not implemented", func() { eightqueens.IsSafe(suite.board, 0, 0) })
}

func (suite *EightQueensSuite) TestPlaceQueens() {
	assert.Nil(suite.T(), eightqueens.PlaceQueens(suite.board, 0))
}

func TestEightQueensSuite(t *testing.T) {
	suite.Run(t, new(EightQueensSuite))
}
