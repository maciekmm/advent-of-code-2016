package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func calculateLength(str string) int {
	totalLength := 0
	indice, afterSplit := false, false
	var charCountBuf, multiBuf bytes.Buffer

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '(':
			indice = true
		case ')':
			charCount, err := strconv.Atoi(charCountBuf.String())
			if err != nil {
				panic(err)
			}
			multiplier, err := strconv.Atoi(multiBuf.String())
			if err != nil {
				panic(err)
			}
			totalLength = totalLength + calculateLength(str[i+1:i+charCount+1])*multiplier

			i = i + charCount
			charCountBuf.Reset()
			multiBuf.Reset()
			indice = false
			afterSplit = false
		default:
			if indice {
				if str[i] == 'x' {
					afterSplit = true
				} else if !afterSplit {
					charCountBuf.WriteByte(str[i])
				} else {
					multiBuf.WriteByte(str[i])
				}
			} else {
				totalLength++
			}
		}
	}
	return totalLength
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	length := 0

	for scanner.Scan() {
		length = length + calculateLength(scanner.Text())
	}
	fmt.Println(length)
}
