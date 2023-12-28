package day_10

import (
	"advent_of_code_2023/utils"
	"strings"
)

func Part2(input string) {
	println("Part 2 not implemented yet!")

	lines := strings.Split(input, "\n")

	maze := make(PipeMaze)
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

	// Remove all non-loop pipe pieces
	for p := range maze {
		if !bfs.Visited[p] {
			maze[p] = '.'
		}
	}

	// Replace initial starting point with actual piece
	start := maze.GetInitial()
	neighbors := maze.GetNeighbors(start)

	switch {
	case neighbors[0] == start.Add(utils.North) && neighbors[1] == start.Add(utils.South):
		maze[start] = '|'
	case neighbors[0] == start.Add(utils.North) && neighbors[1] == start.Add(utils.East):
		maze[start] = 'L'
	case neighbors[0] == start.Add(utils.North) && neighbors[1] == start.Add(utils.West):
		maze[start] = 'J'
	case neighbors[0] == start.Add(utils.South) && neighbors[1] == start.Add(utils.East):
		maze[start] = 'F'
	case neighbors[0] == start.Add(utils.South) && neighbors[1] == start.Add(utils.West):
		maze[start] = '7'
	case neighbors[0] == start.Add(utils.East) && neighbors[1] == start.Add(utils.West):
		maze[start] = '-'
	default:
		println("Unable to find pipe that fits")
	}

	sum := 0
	for p, v := range maze {
		if v == '.' {
			// Empty point
			count := 0
			for j := p.Y + 1; j <= 0; j++ {
				nv := maze[utils.Point{X: p.X, Y: j}]
				if nv == '-' || nv == 'F' || nv == 'L' {
					count++
				}
			}

			if count%2 == 0 {
				maze[p] = 'O'
			} else {
				maze[p] = 'I'
				sum++
			}
		}
	}

	println("Answer:", sum)
}

type PipeMaze map[utils.Point]rune


func (m PipeMaze) GetInitial() utils.Point {
	for k, v := range m {
		if v == 'S' {
			return k
		}
	}

	return utils.Point{}
}

func (m PipeMaze) GetNeighbors(p utils.Point) []utils.Point {
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

func (m PipeMaze) IsFinal(p utils.Point) bool {
	// Going to run until there are no more points
	return false
}
