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

	var letters [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		for i, c := range str {
			if i > len(letters)-1 {
				letters = append(letters, make([]int, 26))
			}
			letters[i][c-'a']++
		}
	}

	buf := new(bytes.Buffer)
	for _, arr := range letters {
		maxIndex := 0
		for j, maxVal := 0, 0; j < len(arr); j++ {
			if arr[j] > maxVal {
				maxVal = arr[j]
				maxIndex = j
			}
		}
		buf.WriteString(string(maxIndex + int('a')))
	}
	fmt.Println(buf.String())
}
