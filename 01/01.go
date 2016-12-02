package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var dir = map[int8][]int{
	//f - x,y
	0: []int{0, 1},  // 0deg
	1: []int{1, 0},  // 90deg
	2: []int{0, -1}, // 180deg
	3: []int{-1, 0}, // 270deg
}

func main() {
	//coords and facing direction
	x, y, f := 0, 0, int8(0)

	mx, my, met := 0, 0, false
	grid := map[string]bool{
		"0x0": true,
	}

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for _, instruction := range strings.Split(" "+string(content), ",") { //add space to the beginning to make instructions consistent
		switch instruction[1] { //take the space
		case 'R':
			f = (f + 1) & 3 // rotate 90 degrees right
		case 'L':
			f = (f - 1) & 3
		}

		steps, err := strconv.Atoi(instruction[2:])

		if err != nil {
			panic(err)
		}

		if !met {
			for i := x + dir[f][0]; i != x+(steps+1)*dir[f][0]; i = i + dir[f][0] {
				if grid[fmt.Sprintf("%dx%d", i, y)] {
					mx = i
					my = y
					met = true
					break
				}

				grid[fmt.Sprintf("%dx%d", i, y)] = true
			}

			for i := y + dir[f][1]; i != y+(steps+1)*dir[f][1]; i = i + dir[f][1] {
				if grid[fmt.Sprintf("%dx%d", x, i)] {
					my = i
					mx = x
					met = true
					break
				}

				grid[fmt.Sprintf("%dx%d", x, i)] = true
			}
		}

		x = x + steps*dir[f][0]
		y = y + steps*dir[f][1]
	}

	fmt.Printf("Distance from start: %d\n", calculateDistance(x, y))
	fmt.Printf("Distance from first meeting: %d\n", calculateDistance(mx, my))
}

func calculateDistance(x, y int) int {
	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}
	return x + y
}
