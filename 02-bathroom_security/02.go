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

	digit := 5
	code := ""
	for _, line := range strings.Split(string(content), "\n") {
		for _, instr := range line {
			switch instr {
			case 'L':
				//1,4,7
				if (digit-1)%3 == 0 {
					continue
				}
				digit--
			case 'R':
				//3,6,9
				if digit%3 == 0 {
					continue
				}
				digit++
			case 'U':
				if digit < 4 {
					continue
				}
				digit = digit - 3
			case 'D':
				if digit > 6 {
					continue
				}
				digit = digit + 3
			}
		}
		code += strconv.Itoa(digit)
	}
	fmt.Println(code)
}
