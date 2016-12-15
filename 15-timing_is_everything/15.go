package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Disc struct {
	Positions int
	Start     int
}

func MustParseInt(str string) int {
	if i, err := strconv.Atoi(str); err != nil {
		panic(err)
	} else {
		return i
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	discs := []Disc{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		discs = append(discs, Disc{
			Positions: MustParseInt(parts[3]),
			Start:     MustParseInt(parts[11][:len(parts[11])-1]),
		})
	}

	discs = append(discs, Disc{
		Positions: 11,
		Start:     0,
	})

	hasPartOne := false

outer:
	for t := 0; ; t++ {
		for i, disc := range discs {
			if i == len(discs)-1 && !hasPartOne {
				fmt.Println("Part 1:", t)
				hasPartOne = true
			}

			if (disc.Start+(i+1)+t)%disc.Positions != 0 {
				continue outer
			}
		}
		fmt.Println("Part 2", t)
		break
	}
}
