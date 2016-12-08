package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const width, height = 50, 6

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	display := make([][]bool, width)
	for i := 0; i < width; i++ {
		display[i] = make([]bool, height)
	}

	for scanner.Scan() {
		str := scanner.Text()
		if str[0:4] == "rect" { // light rectangle
			dimensions := strings.SplitN(str[5:], "x", 2)
			pw, err := strconv.Atoi(dimensions[0])
			if err != nil {
				panic(err)
			}
			ph, err := strconv.Atoi(dimensions[1])
			if err != nil {
				panic(err)
			}
			for i := 0; i < pw; i++ {
				for j := 0; j < ph; j++ {
					display[i][j] = true
				}
			}
		} else { //rotate
			opts := strings.SplitN(str[strings.LastIndexByte(str, byte('='))+1:], " by ", 2)
			id, err := strconv.Atoi(opts[0])
			if err != nil {
				panic(err)
			}
			pixels, err := strconv.Atoi(opts[1])
			if err != nil {
				panic(err)
			}

			if str[7] == 'c' { // rotate column
				temp := display[id][height-pixels:]
				display[id] = append(make([]bool, pixels), display[id][:height-pixels]...)
				copy(display[id], temp)
			} else { //rotate row
				//this can be improved vastly memory wise
				row := make([]bool, width)
				for i := 0; i < width; i++ {
					row[i] = display[i][id]
				}
				temp := row[width-pixels:]
				row = append(make([]bool, pixels), row[:width-pixels]...)
				copy(row, temp)

				for i := 0; i < width; i++ {
					display[i][id] = row[i]
				}
			}
		}
	}
	draw(display)
	lit := 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if display[i][j] {
				lit++
			}
		}
	}
	fmt.Println(lit)
}

func draw(disp [][]bool) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if disp[j][i] == true {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
