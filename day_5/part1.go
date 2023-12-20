package day_5

import (
	"advent_of_code_2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Item struct {
	source      int
	destination int
	rangeLength int
}

func convertLineToItem(line string) Item {
	parts := strings.Fields(line)

	destination, _ := strconv.Atoi(parts[0])
	source, _ := strconv.Atoi(parts[1])
	length, _ := strconv.Atoi(parts[2])

	return Item{
		source:      source,
		destination: destination,
		rangeLength: length,
	}
}

func Part1(input string) {
	var conversionMaps [][]Item

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

	var locations []int

	_, seeds, _ := strings.Cut(mapContent[0], ": ")
	println(seeds)

	for _, seedStr := range strings.Fields(seeds) {
		seed := utils.ParseInt(seedStr)

		for _, conversionMap := range conversionMaps {
			seed = GetLocationFromMap(conversionMap, seed)
		}

		locations = append(locations, seed)
	}

	fmt.Printf("The minimum location is %d \n", utils.Min(locations...))
}

func GetLocationFromMap(conversionMap []Item, seed int) int {
	for _, item := range conversionMap {
		if seed >= item.source && seed <= item.source+item.rangeLength {
			return item.destination + (seed - item.source)
		}
	}
	return seed
}
