package main

import "fmt"

const pairs = 7
const floors = 4
const elements = pairs*2 + 1

//elevator
//gen:mic floor
//gen1:mic2 floor...
type State [elements]uint8

func (s State) Clone() (new State) {
	for i, e := range s {
		new[i] = e
	}
	return
}

func (s State) Moves() []State {
	moves := []State{}
	currentFloor := s[0]
	//1 somewhere
	for i := 1; i < elements; i++ {
		//item x
		if s[i] != currentFloor {
			continue
		}
		if currentFloor < floors-1 {
			cloned := s.Clone()
			cloned[0]++ //elevator
			cloned[i]++ //item
			if cloned.IsValid() {
				moves = append(moves, cloned)
			}
			moves = append(moves, cloned)
		}

		if currentFloor > 0 {
			//down
			cloned := s.Clone()
			cloned[0]--
			cloned[i]--
			if cloned.IsValid() {
				moves = append(moves, cloned)
			}
		}

		//two somewhere
		for j := i + 1; j < elements; j++ {
			if s[j] != currentFloor {
				continue
			}

			//up
			if currentFloor < floors-1 {
				cloned := s.Clone()
				cloned[0]++
				cloned[i]++
				cloned[j]++
				if cloned.IsValid() {
					moves = append(moves, cloned)
				}
			}

			//why would we move two items down, no idea
			// if currentFloor > 0 {
			// 	cloned := s.Clone()
			// 	cloned[0]--
			// 	cloned[i]--
			// 	cloned[j]--
			// 	if cloned.IsValid() {
			// 		moves = append(moves, cloned)
			// 	}
			// }
		}
	}
	return moves
}

func (s State) IsValid() bool {
	for i := uint8(0); i < floors; i++ {
		hasGen, hasUnpairedChip := false, false
		for j := 0; j < pairs; j++ {
			if s[1+j*2] != i && s[2+j*2] != i {
				continue
			}

			if s[1+j*2] == i {
				if hasUnpairedChip {
					return false
				}
				hasGen = true
			}
			if s[1+j*2] != i && s[2+j*2] == i {
				if hasGen {
					return false
				}
				hasUnpairedChip = true
			}
		}
		if hasGen && hasUnpairedChip {
			return false
		}
	}
	return true
}

func (s State) IsFinished() bool {
	//return s == finished
	for i := 0; i < elements; i++ {
		if s[i] != floors-1 {
			return false
		}
	}
	return true
}

func main() {
	initialState := State{0, 0, 0, 0, 1, 0, 1, 2, 2, 2, 2, 0, 0, 0, 0}
	visited := map[State]struct{}{initialState: struct{}{}}
	toExplore := []State{initialState}

	distance := 0

	for {
		distance++
		nextPoints := []State{}
		for _, current := range toExplore {

			for _, adj := range current.Moves() {
				if adj.IsFinished() {
					fmt.Println(distance)
					return
				}
				if _, ok := visited[adj]; !ok {
					visited[adj] = struct{}{}
					nextPoints = append(nextPoints, adj)
				}
			}
		}
		toExplore = nextPoints
	}
}
