package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	supports := 0

	scanner := bufio.NewScanner(file)
	f := func(c rune) bool {
		return c == '[' || c == ']'
	}

outer:
	for scanner.Scan() {
		str := scanner.Text()
		fields := strings.FieldsFunc(str, f)

		//never starts with [] part, so we assume positive indexes are not within the brackets
		found := false
		for i, field := range fields {
			if i%2 == 1 && isValidABBA(field) {
				continue outer
			}
			if i%2 == 0 && isValidABBA(field) {
				found = true
			}
		}
		if found {
			supports++
		}
	}
	fmt.Println(supports)
}

func isValidABBA(sequence string) bool {
	for i := 0; i < len(sequence)-3; i++ {
		part := sequence[i : i+4]
		//not different characters
		if part[0] == part[1] || part[0] != part[3] || part[1] != part[2] {
			continue
		}
		return true
	}
	return false
}
