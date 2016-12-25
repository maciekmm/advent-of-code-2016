package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var moves = []Point{Point{x: 1, y: 0}, Point{x: 0, y: 1}, Point{x: 0, y: -1}, Point{x: -1, y: 0}}

type Point struct {
	x, y int
	repr byte
}

func (p Point) GetAdjacent(maze [][]Point) (adj []Point) {
	for _, move := range moves {
		finY, finX := p.y+move.y, p.x+move.x
		if finY >= len(maze) || finY < 0 || finX < 0 || finX >= len(maze[finY]) {
			continue
		}
		pos := maze[finY][finX]
		if pos.IsOpen() {
			adj = append(adj, pos)
		}
	}
	return
}

func (p Point) IsOpen() bool {
	return p.repr != '#'
}

func (p Point) IsNumber() bool {
	return p.repr >= '0' && p.repr <= '9'
}

func (p Point) DistanceTo(maze [][]Point, dest Point) int {
	distance := 0
	points := []Point{p}
	visited := map[Point]struct{}{}

	for {
		distance++
		nextPoints := []Point{}
		for _, current := range points {
			for _, adj := range current.GetAdjacent(maze) {
				if adj == dest {
					return distance
				}
				if _, ok := visited[adj]; !ok {
					visited[adj] = struct{}{}
					nextPoints = append(nextPoints, adj)
				}
			}
		}
		points = nextPoints
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var maze [][]Point
	var numberPoints []Point
	var pointZero Point

	//build a maze
	for y := 0; scanner.Scan(); y++ {
		str := scanner.Text()

		var xes []Point
		for x, b := range str {
			p := Point{x: x, y: y, repr: byte(b)}
			xes = append(xes, p)
			if p.IsNumber() {
				if p.repr == '0' {
					pointZero = p
					continue
				}
				numberPoints = append(numberPoints, p)
			}
		}
		maze = append(maze, xes)
	}

	ptpDistances := map[Point]map[Point]int{}
	withZero := append(numberPoints, pointZero)
	//fmt.Println(withZero)
	//calculate distances between points
	for i, np := range withZero {
		//we could start at i+1 and only once calculate the distance or copy already computed one
		//but who cares
		for j := 0; j < len(withZero); j++ {
			if i == j {
				continue
			}
			if _, ok := ptpDistances[np]; !ok {
				ptpDistances[np] = map[Point]int{}
			}
			ptpDistances[np][withZero[j]] = np.DistanceTo(maze, withZero[j])
		}
	}

	shortestDistance := math.MaxInt64
	shortestDistancePart2 := math.MaxInt64
	currentLocation, currentDistance := pointZero, 0
	for _, perm := range permutations(numberPoints) {
		for _, loc := range perm {
			currentDistance += ptpDistances[currentLocation][loc]
			currentLocation = loc
		}
		if currentDistance < shortestDistance {
			shortestDistance = currentDistance
		}
		if currentDistance+ptpDistances[currentLocation][pointZero] < shortestDistancePart2 {
			shortestDistancePart2 = currentDistance + ptpDistances[currentLocation][pointZero]
		}
		currentDistance = 0
		currentLocation = pointZero
	}
	fmt.Println("Part 1:", shortestDistance)
	fmt.Println("Part 2:", shortestDistancePart2)
}

//http://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []Point) [][]Point {
	var helper func([]Point, int)
	res := [][]Point{}

	helper = func(arr []Point, n int) {
		if n == 1 {
			tmp := make([]Point, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
