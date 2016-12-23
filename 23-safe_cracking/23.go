package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cloneInstructions(instrs [][]string) [][]string {
	out := make([][]string, len(instrs))
	for i, instr := range instrs {
		out[i] = make([]string, len(instr))
		copy(out[i], instr)
	}
	return out
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var instructions [][]string
	for scanner.Scan() {
		str := scanner.Text()
		instructions = append(instructions, strings.Split(str, " "))
	}

	registerPart1 := map[byte]int{
		'a': 7,
		'b': 0,
		'c': 0,
		'd': 0,
	}
	run(registerPart1, cloneInstructions(instructions))
	fmt.Println("Part 1:", registerPart1['a'])

	registerPart2 := map[byte]int{
		'a': 12,
		'b': 0,
		'c': 0,
		'd': 0,
	}
	run(registerPart2, cloneInstructions(instructions))
	fmt.Println("Part 2:", registerPart2['a'])
}

func isRegister(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func run(register map[byte]int, instructions [][]string) {
	for i := 0; i < len(instructions); i++ {
		//fmt.Println(i, instructions[i:i+5])
		//HARDCORE optimizations :D
		fparam := instructions[i][1]

		switch instructions[i][0] {
		case "inc":
			if !isRegister(fparam[0]) {
				continue
			}
			register[fparam[0]]++
		case "dec":
			if !isRegister(fparam[0]) {
				continue
			}
			register[fparam[0]]--
		default:
			value, err := strconv.Atoi(fparam)
			if err != nil {
				value = register[fparam[0]]
			}

			switch instructions[i][0] {
			case "jnz":
				if value == 0 {
					continue
				}
				sec, err := strconv.Atoi(instructions[i][2])
				if err != nil {
					sec = register[instructions[i][2][0]]
				}
				i += sec - 1
			case "cpy":
				if !isRegister(instructions[i][2][0]) {
					continue
				}

				//NAIVE,shallow multiply optimization, don't even try to understand this
				//optimizes this: [[cpy 0 a] [cpy b c] [inc a] [dec c] [jnz c -2] [dec d] [jnz d -5]]
				//to a = b*d, c=0; d=0
				if i+6 < len(instructions) && value == 0 {
					if register[instructions[i+6][1][0]] != 0 && instructions[i+1][0] == "cpy" && instructions[i+6][0] == "jnz" && instructions[i+5][0] == "dec" {
						register[instructions[i][2][0]] = register[instructions[i+6][1][0]] * register[instructions[i+1][1][0]]
						register[instructions[i+6][1][0]] = 0
						register[instructions[i+1][2][0]] = 0
						i += 6
					}
					continue
				}

				register[instructions[i][2][0]] = value
			case "tgl":
				if i+value < 0 || i+value >= len(instructions) {
					continue
				}
				targetInstr := instructions[i+value]
				//if one argument (instr arg)
				if len(targetInstr) == 2 {
					if targetInstr[0] == "inc" {
						instructions[i+value][0] = "dec"
					} else {
						instructions[i+value][0] = "inc"
					}
				} else {
					if targetInstr[0] == "jnz" {
						instructions[i+value][0] = "cpy"
					} else {
						instructions[i+value][0] = "jnz"
					}
				}
			}
		}

	}
}
