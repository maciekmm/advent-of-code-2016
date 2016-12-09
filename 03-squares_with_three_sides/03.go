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
	lengths := []int{0, 0, 0}
	for _, line := range strings.Split(string(content), "\n") {
		line = strings.Trim(line, " ")
		slens := spaceSplit.Split(line, -1)

		for i, length := range slens {
			lengths[i], err = strconv.Atoi(length)
			if err != nil {
				panic(err)
			}
		}
		sort.Ints(lengths)
		if lengths[2] < (lengths[0] + lengths[1]) {
			possible++
		}
	}

	fmt.Println(possible)
}
