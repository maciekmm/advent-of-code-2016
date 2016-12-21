package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scrambler []byte

func (s Scrambler) SwapPositions(x, y int) {
	s[x], s[y] = s[y], s[x]
}

func (s Scrambler) SwapLetters(x, y byte) {
	s.SwapPositions(bytes.IndexByte(s, x), bytes.IndexByte(s, y))
}

func (s Scrambler) Reverse(x, y int) {
	for i := 0; i <= (y-x)/2; i++ {
		s.SwapPositions(x+i, y-i)
	}
}

func (s Scrambler) Move(x, y int) {
	b := s[x]
	s = append(s[:x], s[x+1:]...)

	temp := []byte{}
	temp = append(temp, s[:y]...)
	temp = append(temp, b)
	temp = append(temp, s[y:]...)
	s = append(s, b)
	copy(s, temp)
}

func (s Scrambler) Rotate(right bool, steps int) {
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

func (s Scrambler) RotateByLetter(x byte) {
	rotations := bytes.IndexByte(s, x)
	if rotations >= 4 {
		rotations++
	}
	rotations++
	s.Rotate(true, rotations)
}

func Scramble(instructions [][]string, input string) string {
	g := Scrambler([]byte(input))
	for _, instr := range instructions {
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
	return string(g)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	instructions := [][]string{}
	for scanner.Scan() {
		instructions = append(instructions, strings.Split(scanner.Text(), " "))
	}
	fmt.Println("Part 1:", Scramble(instructions, "abcdefgh"))

	orig := []byte("fbgdceah")
	for p := make([]byte, len(orig)); int(p[0]) < len(p); nextPerm(p) {
		perm := string(getPerm(orig, p))
		if Scramble(instructions, perm) == "fbgdceah" {
			fmt.Println("Part 2:", perm)
		}
	}
}

//http://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func nextPerm(p []byte) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || int(p[i]) < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []byte) []byte {
	result := append([]byte{}, orig...)
	for i, v := range p {
		result[i], result[i+int(v)] = result[i+int(v)], result[i]
	}
	return result
}
