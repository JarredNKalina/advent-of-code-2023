package day_8

import (
	"advent_of_code_2023/utils"
	"regexp"
)

type Node struct {
	left  string
	right string
	name  string
}

func Part1(input string) {
	const Target = "ZZZ"
	lines := utils.SplitLines(input)

	instruction := lines[0]
	desertMap := getNodes(lines)

	var steps int = 0
	var current string = "AAA"

	for current != Target {
		for i := range instruction {
			direction := string(instruction[i])
			if direction == "R" {
				current = desertMap[current].right
			} else {
				current = desertMap[current].left
			}
			steps++
			if current == Target {
				break
			}
		}
	}
	println(steps)
}

func getNodes(lines []string) map[string]Node {
	desertMap := make(map[string]Node)
	re := regexp.MustCompile("[A-Z0-9]{3}")
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		matches := re.FindAllString(line, -1)
		desertMap[matches[0]] = Node{
			left:  matches[1],
			right: matches[2],
			name:  matches[0],
		}
	}
	return desertMap
}
