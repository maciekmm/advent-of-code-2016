package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	keyboard := []rune{
		' ', ' ', '1', ' ', ' ',
		' ', '2', '3', '4', ' ',
		'5', '6', '7', '8', '9',
		' ', 'A', 'B', 'C', ' ',
		' ', ' ', 'D', ' ', ' ',
	}
	i := 10
	code := ""

	for _, line := range strings.Split(string(content), "\n") {
		for _, instr := range line {
			switch instr {
			case 'L':
				//0,5,10
				if i%5 == 0 || keyboard[i-1] == ' ' {
					continue
				}
				i--
			case 'R':
				//4,9,14
				if (i+1)%5 == 0 || keyboard[i+1] == ' ' {
					continue
				}
				i++
			case 'U':
				if i < 5 || keyboard[i-5] == ' ' {
					continue
				}
				i = i - 5
			case 'D':
				if i > 19 || keyboard[i+5] == ' ' {
					continue
				}
				i = i + 5
			}
		}
		code += string(keyboard[i])
	}
	fmt.Println(code)
}
