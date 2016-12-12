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

	registerPart1 := map[byte]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
	}
	run(registerPart1, instructions)
	fmt.Println("Part 1:", registerPart1['a'])

	registerPart2 := map[byte]int{
		'a': 0,
		'b': 0,
		'c': 1,
		'd': 0,
	}
	run(registerPart2, instructions)
	fmt.Println("Part 2:", registerPart2['a'])
}

func run(register map[byte]int, instructions [][]string) {
	for i := 0; i < len(instructions); i++ {
		fparam := instructions[i][1]

		switch instructions[i][0] {
		case "inc":
			register[fparam[0]]++
		case "dec":
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
}
