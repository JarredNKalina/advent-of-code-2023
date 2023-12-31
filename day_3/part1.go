package day_3

import (
	"advent_of_code_2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Part1(input string) {
	lines := utils.SplitLines(input)

	var directions = [][2]int{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0}, {1, 0},
		{-1, -1}, {0, -1}, {1, -1},
	}

	re := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,<>\/?]`)

	hasAdjacentSymbol := func(colStart int, colEnd int, row int) bool {
		for i := colStart; i < colEnd; i++ {
			for _, dir := range directions {
				newCol := i + dir[0]
				newRow := row + dir[1]

				if newCol >= 0 && newCol < len(lines[0]) && newRow >= 0 && newRow < len(lines) {
					if re.MatchString(string(lines[newRow][newCol])) {
						return true
					}
				}
			}
		}
		return false
	}

	totalSum := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if utils.IsDigit(lines[i][j]) {
				endX := j

				for endX < len(lines[0]) && lines[i][endX] != '.' && !re.MatchString(string(lines[i][endX])) {
					endX++
				}

				if hasAdjacentSymbol(j, endX, i) {
					numeral, _ := strconv.Atoi(lines[i][j:endX])
					totalSum += numeral
				}

				j = endX
			}
		}
	}

	fmt.Printf("The total sum of the good engine parts is %d\n", totalSum)
}
