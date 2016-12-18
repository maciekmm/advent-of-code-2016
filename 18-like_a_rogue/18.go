package main

import (
	"fmt"
	"io/ioutil"
)

var trapSettings = [4][3]bool{
	[3]bool{true, true, false},
	[3]bool{false, true, true},
	[3]bool{true, false, false},
	[3]bool{false, false, true},
}

func isNewTrapped(i int, row []bool) bool {
	prev := [3]bool{isTrapped(i-1, row), isTrapped(i, row), isTrapped(i+1, row)}
	for _, set := range trapSettings {
		if set == prev {
			return true
		}
	}
	return false
}

func isTrapped(i int, r []bool) bool {
	if i < 0 || i >= len(r) {
		return false
	}
	return r[i]
}

func countSafeTiles(input string, rows int) int {
	row := make([]bool, len(input))

	safeTiles := 0
	for i, b := range input {
		if b == '.' {
			row[i] = false
			safeTiles++
		} else {
			row[i] = true
		}
	}

	temporaryRow := make([]bool, len(input))
	for r := 0; r < rows-1; r++ {
		for i := range row {
			trapped := isNewTrapped(i, row)
			if !trapped {
				safeTiles++
			}
			temporaryRow[i] = trapped
		}
		copy(row, temporaryRow)
	}
	return safeTiles
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", countSafeTiles(string(content), 40))
	fmt.Println("Part 2:", countSafeTiles(string(content), 400000))
}
