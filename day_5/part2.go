package day_5

import (
	"advent_of_code_2023/utils"
	"regexp"
	"strings"
)

func Part2(input string) {
	var conversionMaps [][]Item
	const groupSize = 100000
	re := regexp.MustCompile(`\r?\n\r?\n`)
	mapContent := re.Split(input, -1)

	for _, item := range mapContent[1:] {
		mapStrs := strings.FieldsFunc(item, func(r rune) bool { return r == '\n' })

		var conversionMap []Item
		for _, mapStr := range mapStrs[1:] {
			if len(mapStr) == 0 {
				continue
			}
			conversionMap = append(conversionMap, convertLineToItem(mapStr))
		}
		conversionMaps = append(conversionMaps, conversionMap)
	}

	_, initialSeeds, _ := strings.Cut(mapContent[0], ": ")
	initialSeedArr := strings.Fields(initialSeeds)
	var seedPairs [][2]int

	for i, seedStr := range initialSeedArr {
		if i%2 == 0 {
			firstNum := utils.ParseInt(seedStr)
			secondNum := utils.ParseInt(initialSeedArr[i+1])
			seedPairs = append(seedPairs, [2]int{firstNum, secondNum})
		}
	}

	var seeds []int

	for _, seedPair := range seedPairs {
		start := seedPair[0]
		rangeNum := seedPair[1]

		for i := 0; i < rangeNum; i++ {
			seeds = append(seeds, start+i)
		}
	}

	println(len(seeds))
	numGroups := (len(seeds) + groupSize - 1) / groupSize

	var locations []int

	for i := 0; i < numGroups; i++ {
		startIndex := i * groupSize
		endIndex := (i + 1) * groupSize

		if endIndex > len(seeds) {
			endIndex = len(seeds)
		}

		group := seeds[startIndex:endIndex]

		locations = append(locations, processGroup(group, conversionMaps))
	}
	println("final")
	println(utils.Min(locations...))
}

func processGroup(group []int, conversionMaps [][]Item) int {
	var locations []int
	for _, seed := range group {
		for _, conversionMap := range conversionMaps {
			seed = GetLocationFromMap(conversionMap, seed)
		}

		locations = append(locations, seed)
	}

	return utils.Min(locations...)
}
