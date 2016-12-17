package main

import "fmt"

var moves = []Point{Point{1, 0}, Point{0, 1}, Point{0, -1}, Point{-1, 0}}

type Point struct {
	x, y int
}

func (p Point) GetAdjacent() (adj []Point) {
	for _, move := range moves {
		pos := Point{p.x + move.x, p.y + move.y}
		if pos.IsOpen() {
			adj = append(adj, pos)
		}
	}
	return
}

func (p Point) IsOpen() bool {
	if p.x < 0 || p.y < 0 {
		return false
	}
	open := true
	sum := (p.x+p.y)*(p.x+p.y) + 3*p.x + p.y + input
	for sum > 0 {
		if (sum & 1) == 1 {
			open = !open
		}
		sum >>= 1
	}
	return open
}

func (p Point) ShortestPath(dest Point, maxDistance int) (int, int) {
	distance := 0
	points := []Point{p}
	visited := map[Point]struct{}{}

	for maxDistance == 0 || distance < maxDistance {
		distance++
		nextPoints := []Point{}
		for _, current := range points {
			for _, adj := range current.GetAdjacent() {
				if adj == dest {
					return distance, len(visited)
				}
				if _, ok := visited[adj]; !ok {
					visited[adj] = struct{}{}
					nextPoints = append(nextPoints, adj)
				}
			}
		}
		points = nextPoints
	}
	return distance, len(visited)
}

const input = 1352

func main() {
	d, _ := Point{1, 1}.ShortestPath(Point{31, 39}, 0)
	fmt.Println("Distance:", d)
	_, vis := Point{1, 1}.ShortestPath(Point{31, 39}, 50)
	fmt.Println("Points visited:", vis)
}
