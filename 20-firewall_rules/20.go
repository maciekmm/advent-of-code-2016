package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	LowerBound, UpperBound int
}

type RangeEntries []Range

func (re RangeEntries) Len() int {
	return len(re)
}

func (re RangeEntries) Swap(i, j int) {
	re[i], re[j] = re[j], re[i]
}

func (re RangeEntries) Less(i, j int) bool {
	return re[i].LowerBound < re[j].LowerBound
}

const maxValue = 4294967295

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	re := RangeEntries{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		lowerBound, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		upperBound, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		re = append(re, Range{LowerBound: lowerBound, UpperBound: upperBound})
	}

	sort.Sort(re)
	prev := Range{
		LowerBound: 0,
		UpperBound: 0,
	}

	allowed := 0
	for _, r := range re {
		if r.LowerBound > prev.UpperBound+1 {
			if allowed == 0 {
				fmt.Println("Part 1:", prev.UpperBound+1)
			}
			allowed++
		}
		if r.UpperBound < prev.UpperBound {
			r.UpperBound = prev.UpperBound
		}
		prev = r
	}
	allowed += (maxValue - prev.UpperBound)
	fmt.Println("Part 2:", allowed)
}
