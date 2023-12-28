package day_9

import (
	"advent_of_code_2023/utils"
	"fmt"
	"strings"
)

func Part1(input string) {
	lines := strings.Split(input, "\n")
	totalSum := 0

	for _, line := range lines {
		elems := strings.Fields(line)
		var elemsInt []int

		for i := range elems {
			elemsInt = append(elemsInt, utils.ParseInt(elems[i]))
		}

		totalSum += RecursivelyFindNextProgression(elemsInt)
	}

	fmt.Printf("The sum is %d\n", totalSum)
}

func RecursivelyFindNextProgression(elems []int) int {
	var diffs []int

	for i := 0; i < len(elems)-1; i++ {
		diffs = append(diffs, elems[i+1]-elems[i])
	}

	if isLineEmpty(diffs) {
		return elems[len(elems)-1] + diffs[0]
	}

	lastDiff := RecursivelyFindNextProgression(diffs)

	return elems[len(elems)-1] + lastDiff
}

func isLineEmpty(line []int) bool {
	for _, num := range line {
		if num != 0 {
			return false
		}
	}
	return true
}
