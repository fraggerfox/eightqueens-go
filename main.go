package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fraggerfox/eightqueens-go/eightqueens"
)

func PrintSolution(boards eightqueens.Boards) {
	for _, board := range boards {
		for i := 0; i < eightqueens.BOARD_SIZE; i++ {
			for j := 0; j < eightqueens.BOARD_SIZE; j++ {
				if board[i][j] {
					fmt.Print("Q ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {
	os.Exit(realMain(os.Stdout))
}

// There is no proper way to test main() in golang, so the workaround is to
// keep main() as minimal as possible and write the logic into a proxy method
// like shown below.
//
// References:
// https://mj-go.in/golang/test-the-main-function-in-go
// https://github.com/hashicorp/terraform/blob/main/main.go#L57
func realMain(out io.Writer) int {
	var board eightqueens.Board
	solutions := make(eightqueens.Boards, 0)
	solutions = eightqueens.PlaceQueens(board, 0, &solutions)
	PrintSolution(solutions)
	fmt.Fprintf(out, "Found %d solutions.\n", len(solutions))

	return 0
}
