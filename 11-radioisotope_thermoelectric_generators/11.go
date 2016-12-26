package main

import (
	"fmt"
	"sort"
)

const pairs = 7
const floors = 4

type Pairs [pairs][2]uint8

//elevator
//gen:mic floor
//gen1:mic2 floor...
type State struct {
	Elevator uint8
	Substate Pairs
}

// Len is part of sort.Interface.
func (s *State) Len() int {
	return len(s.Substate)
}

// Swap is part of sort.Interface.
func (s *State) Swap(i, j int) {
	s.Substate[i], s.Substate[j] = s.Substate[j], s.Substate[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *State) Less(i, j int) bool {
	if s.Substate[i][0] == s.Substate[j][0] {
		return s.Substate[i][1] > s.Substate[j][1]
	}
	return s.Substate[i][0] > s.Substate[j][0]
}

func (s *State) Simplify() {
	sort.Sort(s)
}

func (s *State) Clone() *State {
	new := &State{Elevator: s.Elevator}
	for i, pair := range s.Substate {
		for j, floor := range pair {
			new.Substate[i][j] = floor
		}
	}
	return new
}

func (s *State) Moves() []*State {
	moves := []*State{}
	currentFloor := s.Elevator
	//1 somewhere
	for i, pair := range s.Substate {
		for j := 0; j < 2; j++ {
			if pair[j] != currentFloor {
				continue
			}

			if currentFloor > 0 {
				//down
				cloned := s.Clone()
				cloned.Elevator--
				cloned.Substate[i][j]--
				if cloned.IsValid() {
					moves = append(moves, cloned)
				}
			}

			movedTwo := false
			//two somewhere
			l := j + 1
			for k := i; k < len(s.Substate); k++ {
				for ; l < 2; l++ {
					if s.Substate[k][l] != currentFloor {
						continue
					}

					//up
					if currentFloor < floors-1 {
						cloned := s.Clone()
						cloned.Elevator++
						cloned.Substate[i][j]++
						cloned.Substate[k][l]++
						if cloned.IsValid() {
							moves = append(moves, cloned)
							movedTwo = true
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
				l = 0
			}

			//move up
			//if two are being moved up, don't move one
			if !movedTwo && currentFloor < floors-1 {
				cloned := s.Clone()
				cloned.Elevator++       //elevator
				cloned.Substate[i][j]++ //item
				if cloned.IsValid() {
					moves = append(moves, cloned)
				}
			}
		}
		//item x

	}
	return moves
}

func (s *State) IsValid() bool {
	for i := uint8(0); i < floors; i++ {
		hasGen, hasUnpairedChip := false, false
		for _, pair := range s.Substate {
			if pair[0] != i && pair[1] != i {
				continue
			}

			if pair[0] == i {
				if hasUnpairedChip {
					return false
				}
				hasGen = true
			}
			if pair[0] != i && pair[1] == i {
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

func (s *State) IsFinished() bool {
	//return s == finished
	for _, sub := range s.Substate {
		for _, p := range sub {
			if p != floors-1 {
				return false
			}
		}
	}
	return true
}

func main() {
	initialState := &State{Elevator: 0, Substate: Pairs{[2]uint8{0, 0}, [2]uint8{0, 1}, [2]uint8{0, 1}, [2]uint8{2, 2}, [2]uint8{2, 2}, [2]uint8{0, 0}, [2]uint8{0, 0}}}
	visited := map[State]struct{}{*initialState: struct{}{}}
	toExplore := []*State{initialState}

	distance := 0

	for {
		distance++
		//fmt.Println(distance, len(toExplore))
		nextPoints := []*State{}
		for _, current := range toExplore {
			for _, adj := range current.Moves() {
				if adj.IsFinished() {
					fmt.Println(distance)
					return
				}
				if _, ok := visited[*adj]; !ok {
					adj.Simplify()
				} else {
					continue
				}
				if _, ok := visited[*adj]; !ok {
					visited[*adj] = struct{}{}
					nextPoints = append(nextPoints, adj)
				}
			}
		}
		toExplore = nextPoints
	}
}
