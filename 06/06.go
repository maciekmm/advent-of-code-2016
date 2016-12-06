package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var letters [][]uint

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		for i, c := range str {
			if i > len(letters)-1 {
				letters = append(letters, make([]uint, 26))
			}
			letters[i][c-'a']++
		}
	}

	buf := new(bytes.Buffer)
	buf.WriteString("Part1: ")
	buf2 := new(bytes.Buffer)
	buf2.WriteString("Part2: ")
	for _, arr := range letters {
		maxIndex, minIndex := 0, 0
		for j, maxVal, minVal := 0, uint(0), ^uint(0); j < len(arr); j++ {
			if arr[j] > maxVal {
				maxVal = arr[j]
				maxIndex = j
			}
			if arr[j] != 0 && arr[j] < minVal {
				minVal = arr[j]
				minIndex = j
			}
		}
		buf.WriteString(string(maxIndex + int('a')))
		buf2.WriteString(string(minIndex + int('a')))
	}
	fmt.Println(buf.String())
	fmt.Println(buf2.String())
}
