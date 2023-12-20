package day_8

import "advent_of_code_2023/utils"

func Part2(input string) {
	lines := utils.SplitLines(input)

	desertMap := getNodes(lines)

	startingNodes := make(map[string]Node)
	instructions := lines[0]
	for _, node := range desertMap {
		if node.name[2] == 'A' {
			startingNodes[node.name] = node
		}
	}
	var loops []int
	for _, node := range startingNodes {
		loops = append(loops, getLoops(node.name, instructions, desertMap))
	}

	for _, loops := range loops {
		println(loops)
	}

	println(utils.FindLCM(loops...))
}

func getLoops(startingNode string, instructions string, directionMap map[string]Node) int {
	var current string = startingNode

	var steps int = 0

	for current[2] != 'Z' {
		for _, instruction := range instructions {
			if instruction == 'R' {
				current = directionMap[current].right
			} else {
				current = directionMap[current].left
			}
			steps++
			if current[2] == 'Z' {
				break
			}
		}
	}
	return steps
}
