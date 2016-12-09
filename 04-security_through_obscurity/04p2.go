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

	for _, line := range strings.Split(string(content), "\n") {
		parts := strings.Split(line, "-")
		last := parts[len(parts)-1]
		num, err := strconv.Atoi(last[:len(last)-7])
		if err != nil {
			panic(err)
		}
		if strings.Contains(strings.Map(func(run rune) rune {
			//ROT x
			return rune((int(run)-'a'+num)%26 + 'a')
		}, strings.Join(parts[:len(parts)-1], "")), "north") {
			fmt.Println(num)
			return
		}
	}
}
