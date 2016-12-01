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

		i, err := strconv.Atoi(instruction[2:])

		if err != nil {
			panic(err)
		}

		x = x + i*dir[f][0]
		y = y + i*dir[f][1]

		fmt.Printf("X: %d, Y: %d\n", x, y)
	}

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	fmt.Println(x + y)
}
