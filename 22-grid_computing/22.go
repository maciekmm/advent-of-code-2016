package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var spaceSplit = regexp.MustCompile(" +")

type Node struct {
	X, Y       int
	Size, Used int
}

func (n Node) Free() int {
	return n.Size - n.Used
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	nodes := []Node{}

	for scanner.Scan() {
		rawNode := spaceSplit.Split(scanner.Text(), -1)
		node := Node{}

		//coordinates
		coords := strings.Split(rawNode[0], "-")
		if x, err := strconv.Atoi(coords[len(coords)-2][1:]); err != nil {
			panic(err)
		} else {
			node.X = x
		}
		if y, err := strconv.Atoi(coords[len(coords)-1][1:]); err != nil {
			panic(err)
		} else {
			node.Y = y
		}

		//size
		if size, err := strconv.Atoi(rawNode[1][:len(rawNode[1])-1]); err != nil {
			panic(err)
		} else {
			node.Size = size
		}

		//Used
		if used, err := strconv.Atoi(rawNode[2][:len(rawNode[2])-1]); err != nil {
			panic(err)
		} else {
			node.Used = used
		}
		nodes = append(nodes, node)
	}

	possibilities := 0

	for _, source := range nodes {
		for _, destination := range nodes {
			if source.Used == 0 {
				continue
			}
			if source == destination {
				continue
			}
			if destination.Free() < source.Used {
				continue
			}
			possibilities++
		}
	}
	fmt.Println(possibilities)
}
