package day_6

import (
	"advent_of_code_2023/utils"
	"regexp"
)

func extractNumbers(line string) []int {
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(line, -1)
	var numbers []int

	for _, match := range matches {
		numbers = append(numbers, utils.ParseInt(match))
	}
	return numbers
}

func Part1(input string) {
	lines := utils.SplitLines(input)

	times := extractNumbers(lines[0])
	distances := extractNumbers(lines[1])

	var result = 1
	for i, distance := range distances {

		var validDistances int
		time := times[i]

		// println(time, distance)
		for j := 0; j < time; j++ {
			if willFinish(j, time, distance) {
				validDistances++
			}
		}
		// println(validDistances)
		result = result * validDistances
	}

	println("The result is", result)

}

func willFinish(pressTime int, timeToBeat int, distanceToTravel int) bool {
	// println((timeToBeat - pressTime) * pressTime)
	return (timeToBeat-pressTime)*pressTime > distanceToTravel
}
