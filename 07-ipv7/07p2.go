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

	var parts map[string]interface{}

outer:
	for scanner.Scan() {
		str := scanner.Text()
		fields := strings.FieldsFunc(str, f)
		parts = make(map[string]interface{})

		//never starts with [] part, so we assume positive indexes are not within the brackets
		//push all values within brackets to a map
		//this is also not very efficient, but works
		for i := 1; i < len(fields); i = i + 2 {
			for j := 0; j < len(fields[i])-2; j++ {
				part := fields[i][j : j+3]
				if part[0] != part[2] {
					continue
				}
				//transform and put the lookup-able form into map
				parts[string(part[1])+string(part[0])+string(part[1])] = struct{}{}
			}
		}

		//check if any sequence has its sister in map
		for i := 0; i < len(fields); i = i + 2 {
			for j := 0; j < len(fields[i])-2; j++ {
				if _, ok := parts[fields[i][j:j+3]]; ok {
					supports++
					continue outer
				}
			}
		}
	}
	fmt.Println(supports)
}
