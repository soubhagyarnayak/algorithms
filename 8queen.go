package datastructure

import (
	"fmt"
)

func IsValidPlacement(placements [][]bool, row int, col int) bool {
	for i := 0; i < col; i++ {
		if placements[row][i] {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if placements[i][j] {
			return false
		}
	}
	for i, j := row, col; i < len(placements) && j >= 0; i, j = i+1, j-1 {
		if placements[i][j] {
			return false
		}
	}
	return true
}

func PrintPlacements(placements [][]bool) {
	for i := 0; i < len(placements); i++ {
		for j := 0; j < len(placements[i]); j++ {
			if placements[i][j] {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}
}

func SolveNQueenInternal(placements [][]bool, col int) bool {
	if col >= len(placements[0]) {
		return true
	}
	for i := 0; i < len(placements); i++ {
		if IsValidPlacement(placements, i, col) {
			placements[i][col] = true
			if SolveNQueenInternal(placements, col+1) {
				return true
			}
			placements[i][col] = false
		}
	}
	return false
}

func SolveNQueen(n int) {
	placements := make([][]bool, n)
	for i := 0; i < n; i++ {
		placements[i] = make([]bool, n)
	}
	if SolveNQueenInternal(placements, 0) {
		PrintPlacements(placements)
	} else {
		fmt.Printf("No solution exists!")
	}
}
