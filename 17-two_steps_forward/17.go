package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var moves = []Point{Point{0, -1, "U"}, Point{0, 1, "D"}, Point{-1, 0, "L"}, Point{1, 0, "R"}}

type Point struct {
	x, y int
	path string
}

func (p Point) GetAdjacent(path string) (adj []Point) {
	md := md5.New()
	md.Write(passcode)
	md.Write([]byte(path))
	hash := fmt.Sprintf("%x", md.Sum(nil))[:4]

	for i, move := range moves {
		pos := Point{p.x + move.x, p.y + move.y, path + move.path}
		if pos.x < 0 || pos.x > 3 || pos.y < 0 || pos.y > 3 {
			continue
		}
		if strings.ContainsAny(string(hash[i]), "bcdef") {
			adj = append(adj, pos)
		}
	}
	return
}

func (p Point) Path(dest Point) (shortest string, longestLength int) {
	points := []Point{p}
	for len(points) > 0 {
		nextPoints := []Point{}
		for _, current := range points {
			for _, adj := range current.GetAdjacent(current.path) {
				if adj.x == dest.x && adj.y == dest.y {
					if len(shortest) == 0 {
						shortest = adj.path
					}
					longestLength = len(adj.path)
					continue
				}
				nextPoints = append(nextPoints, adj)
			}
		}
		points = nextPoints
	}
	return
}

var destination = Point{x: 3, y: 3}
var passcode = []byte("pgflpeqp")

func main() {
	fmt.Println(Point{x: 0, y: 0}.Path(destination))
}
