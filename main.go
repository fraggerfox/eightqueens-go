package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fraggerfox/eightqueens-go/eightqueens"
)

func main() {
	os.Exit(realMain(os.Stdout))
}

// https://mj-go.in/golang/test-the-main-function-in-go
// https://github.com/hashicorp/terraform/blob/main/main.go#L57
func realMain(out io.Writer) int {
	var board eightqueens.Board
	solutions := eightqueens.PlaceQueens(board, 0)
	fmt.Fprintf(out, "Found %d solutions.\n", len(solutions))

	return 0
}
