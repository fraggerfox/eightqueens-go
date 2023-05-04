# Eight Queens solution in Golang

## Introduction

This was an attempt to learn some of the constructs of writing and testing a Golang program when I first got introduced to the language in my job in 2019. Due to the authoritative and abusive nature of my back then Tech lead, who had no clue on how to teach people on programming, I came up with the idea of a programming club where we (couple of my colleages and I) tried to do a Test Driven approach to learning Golang away from this abusive and authoritative person.

The Eight Queens problem is one of the first problems I got introduced to during my early days as a programmer and it was implemented by me by reading the original paper "Problem solving by stepwise refinement" by Nikalus Wirth. Since it is a familiar problem to me, it helps me to focus on the nuances of the language I am learning and the various libraries I use to interact with it.

Due to some unforseen circumstances, I was not able to fulfill the promise of setting up a programming club and never completed the task. But with the advent of ChatGPT was trying trying out some simple programming exercises and thought of why not ressurect this original problem

## ChatGPT's solution

ChatGPT gave a very compact solution which I shall post here for reference.

```go
package main

import "fmt"

const BOARD_SIZE = 8

type Board [BOARD_SIZE][BOARD_SIZE]bool

var Solutions int

func PrintSolution(board Board) {
    fmt.Printf("Solution %d:\n", Solutions)
    for i := 0; i < BOARD_SIZE; i++ {
        for j := 0; j < BOARD_SIZE; j++ {
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

func IsSafe(board Board, row, col int) bool {
    // Check row and column
    for i := 0; i < BOARD_SIZE; i++ {
        if board[row][i] || board[i][col] {
            return false
        }
    }

    // Check diagonals
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

func PlaceQueens(board Board, row int) {
    if row == BOARD_SIZE {
        Solutions++
        PrintSolution(board)
        return
    }

    for col := 0; col < BOARD_SIZE; col++ {
        if IsSafe(board, row, col) {
            board[row][col] = true
            PlaceQueens(board, row+1)
            board[row][col] = false
        }
    }
}

func main() {
    var board Board
    Solutions = 0
    PlaceQueens(board, 0)
    fmt.Printf("Found %d solutions.\n", Solutions)
}
```

### Explanation by ChatGPT

The program uses a 2D array called `Board` to represent the chessboard, where each element is a boolean value indicating whether a queen is placed on that square. The `PlaceQueens` function uses a recursive backtracking algorithm to place queens on the board, checking at each step whether it is safe to place a queen in a particular row and column.

The `IsSafe` function checks whether it is safe to place a queen at a particular row and column by checking whether there is already a queen in the same row, column, or diagonal.

The `PrintSolution` function prints out a solution in a human-readable format.

Finally, the main program calls `PlaceQueens` to find all solutions and prints out the number of solutions found.

## What is done by me?

Once the boilerplate solution was given, it was a process of breaking it down into sizeable chunks which, started off as the base point for which I wrote the tests using `testify` library.

The tests for the respective broken down blocks have been written in the respective `eightqueens_test.go` and `eightqueens_internal_test.go`. The internal test is specifically to test the unexposed parts of the Golang code in the `eightqueens` package.

In addition to this `main.go` has a simple test written to show how to improve the testing of logic that would reside in `func main()` in Golang.