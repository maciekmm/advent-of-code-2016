package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, line := range strings.Split(string(content), "\n") {
		parts := strings.Split(line, "-")
		last := parts[len(parts)-1]
		if isChecksumValid(strings.Join(parts[:len(parts)-1], ""), last[len(last)-6:len(last)-1]) {
			num, err := strconv.Atoi(last[:len(last)-7])
			if err != nil {
				panic(err)
			}
			total = total + num
		}
	}
	fmt.Println(total)
}

func isChecksumValid(str, checksum string) bool {
	runes := make(map[rune]uint, 26)
	for _, char := range str {
		if _, ok := runes[char]; !ok {
			runes[char] = 0
		}
		runes[char]++
	}

	lowestValue := ^uint(0)
	var lowestRune rune

	for _, c := range checksum {
		if count, ok := runes[c]; ok {
			if count > lowestValue || (count == lowestValue && lowestRune > c) {
				return false
			}
			lowestValue = count
			lowestRune = c
			delete(runes, c)
			continue
		}
		return false
	}

	for c, count := range runes {
		if count > lowestValue || (count == lowestValue && lowestRune > c) {
			return false
		}
	}
	return true
}
