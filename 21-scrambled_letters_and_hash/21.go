package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scramble []byte

func (s Scramble) SwapPositions(x, y int) {
	s[x], s[y] = s[y], s[x]
}

func (s Scramble) SwapLetters(x, y byte) {
	s.SwapPositions(bytes.IndexByte(s, x), bytes.IndexByte(s, y))
}

func (s Scramble) Reverse(x, y int) {
	for i := 0; i <= (y-x)/2; i++ {
		s.SwapPositions(x+i, y-i)
	}
}

func (s Scramble) Move(x, y int) {
	b := s[x]
	s = append(s[:x], s[x+1:]...)

	temp := []byte{}
	temp = append(temp, s[:y]...)
	temp = append(temp, b)
	temp = append(temp, s[y:]...)
	s = append(s, b)
	copy(s, temp)
}

func (s Scramble) Rotate(right bool, steps int) {
	//overflow precautions
	steps = steps % len(s)

	if !right {
		steps = len(s) - steps
	}

	temp := []byte{}
	temp = append(temp, s[len(s)-steps:]...)
	copy(s[steps:], s[:len(s)-steps])
	copy(s, temp)
}

func (s Scramble) RotateByLetter(x byte) {
	rotations := bytes.IndexByte(s, x)
	if rotations >= 4 {
		rotations++
	}
	rotations++
	s.Rotate(true, rotations)
}

func main() {
	g := Scramble([]byte("abcdefgh"))
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instr := strings.Split(scanner.Text(), " ")
		switch instr[0] {
		case "swap":
			if instr[1] == "position" {
				x, _ := strconv.Atoi(instr[2])
				y, _ := strconv.Atoi(instr[5])
				g.SwapPositions(x, y)
			} else {
				g.SwapLetters(instr[2][0], instr[5][0])
			}
		case "reverse":
			x, _ := strconv.Atoi(instr[2])
			y, _ := strconv.Atoi(instr[4])
			g.Reverse(x, y)
		case "move":
			x, _ := strconv.Atoi(instr[2])
			y, _ := strconv.Atoi(instr[5])
			g.Move(x, y)
		case "rotate":
			if instr[1] == "based" {
				g.RotateByLetter(instr[6][0])
			} else {
				right := true
				if instr[1] == "left" {
					right = false
				}
				x, _ := strconv.Atoi(instr[2])
				g.Rotate(right, x)
			}
		}
	}

	fmt.Println(string(g))
}
