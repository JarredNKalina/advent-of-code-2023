package day_10

import (
	"advent_of_code_2023/utils"
	"strings"
)

func Part1(input string) {
	lines := strings.Split(input, "\n")

	maze := make(pipeMaze)
	for j, ln := range lines {
		for i, r := range ln {
			p := utils.Point{X: i, Y: -j}
			maze[p] = r
		}
	}

	bfs := utils.NewBFS[utils.Point]()

	_, err := bfs.Run(maze)

	if err != utils.ErrBFSNotFound {
		println("Error running BFS:", err)
	}

	var max uint64
	for _, v := range bfs.Distance {
		if v > max {
			max = v
		}
	}

	println("Answer:", max)
}

type pipeMaze map[utils.Point]rune

func (m pipeMaze) GetInitial() utils.Point {
	for k, v := range m {
		if v == 'S' {
			return k
		}
	}

	return utils.Point{}
}

func (m pipeMaze) GetNeighbors(p utils.Point) []utils.Point {
	var directions []utils.Point

	switch m[p] {
	case '|':
		directions = []utils.Point{utils.North, utils.South}
	case '-':
		directions = []utils.Point{utils.East, utils.West}
	case 'L':
		directions = []utils.Point{utils.North, utils.East}
	case 'J':
		directions = []utils.Point{utils.North, utils.West}
	case '7':
		directions = []utils.Point{utils.South, utils.West}
	case 'F':
		directions = []utils.Point{utils.South, utils.East}
	case '.':
		directions = []utils.Point{}
	case 'S':
		directions = []utils.Point{utils.North, utils.South, utils.East, utils.West}
	}

	// Validate connector and add if there is a connecting pipe
	neighbors := []utils.Point{}

	for _, d := range directions {
		switch d {
		case utils.North:
			n := p.Add(d)
			if m[n] == '|' || m[n] == '7' || m[n] == 'F' {
				neighbors = append(neighbors, n)
			}
		case utils.South:
			n := p.Add(d)
			if m[n] == '|' || m[n] == 'L' || m[n] == 'J' {
				neighbors = append(neighbors, n)
			}
		case utils.East:
			n := p.Add(d)
			if m[n] == '-' || m[n] == 'J' || m[n] == '7' {
				neighbors = append(neighbors, n)
			}
		case utils.West:
			n := p.Add(d)
			if m[n] == '-' || m[n] == 'L' || m[n] == 'F' {
				neighbors = append(neighbors, n)
			}
		}
	}

	return neighbors
}

func (m pipeMaze) IsFinal(p utils.Point) bool {
	return false
}
