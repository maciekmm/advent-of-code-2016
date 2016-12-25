package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	register := map[byte]int{}

	for i := uint(0); i < ^uint(0); i++ {
		register['a'] = int(i)
		register['b'] = 0
		register['c'] = 0
		register['d'] = 0
		if run(register, instructions) {
			fmt.Println("Part 1: ", i)
			break
		}
	}
}

func run(register map[byte]int, instructions [][]string) bool {
	last, tries := 1, 64
	for i := 0; i < len(instructions); i++ {
		fparam := instructions[i][1]

		switch instructions[i][0] {
		case "inc":
			register[fparam[0]]++
		case "dec":
			register[fparam[0]]--
		case "out":
			num := register[fparam[0]]
			if (last == 1 && num == 0) || (last == 0 && num == 1) {
				tries--
				last = num
			} else {
				return false
			}
			if tries <= 0 {
				return true
			}
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
				if offset, err := strconv.Atoi(instructions[i][2]); err != nil {
					panic(err)
				} else {
					i += offset - 1
				}
			case "cpy":
				register[instructions[i][2][0]] = value
			}
		}
	}
	return false
}
