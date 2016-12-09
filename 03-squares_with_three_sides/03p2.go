package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var spaceSplit = regexp.MustCompile(" +")

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	possible := 0
	triangles := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for i, line := range strings.Split(string(content), "\n") {
		line = strings.Trim(line, " ")
		slens := spaceSplit.Split(line, -1)

		for j, length := range slens {
			triangles[j][i%3], err = strconv.Atoi(length)
			if err != nil {
				panic(err)
			}
		}

		if (i-2)%3 == 0 {
			for k := 0; k < 3; k++ {
				sort.Ints(triangles[k])
				if triangles[k][2] < (triangles[k][0] + triangles[k][1]) {
					possible++
				}
			}
		}
	}

	fmt.Println(possible)
}
