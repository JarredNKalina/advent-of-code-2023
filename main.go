package main

import (
	"advent_of_code_2023/day_10"
	"advent_of_code_2023/day_11"
	"advent_of_code_2023/day_3"
	"advent_of_code_2023/day_4"
	"advent_of_code_2023/day_5"
	"advent_of_code_2023/day_6"
	"advent_of_code_2023/day_7"
	"advent_of_code_2023/day_8"
	"advent_of_code_2023/day_9"
	"advent_of_code_2023/utils"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("You need to provide both the day and the path to input!\n\t%s <day> <path/to/sample.txt>\n", os.Args[0])
	}

	day := os.Args[1]
	filePath := os.Args[2]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File not found: %s\n", filePath)
	}

	input, err := utils.ReadFileIntoString(filePath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	switch day {
	case "3":
		runDay(day, day_3.Part1, day_3.Part2, input)
	case "4":
		runDay(day, day_4.Part1, day_4.Part2, input)
	case "5":
		runDay(day, day_5.Part1, day_5.Part2, input)
	case "6":
		runDay(day, day_6.Part1, nil, input)
	case "7":
		runDay(day, day_7.Part1, day_7.Part2, input)
	case "8":
		runDay(day, nil, day_8.Part2, input)
	case "9":
		runDay(day, day_9.Part1, day_9.Part2, input)
	case "10":
		runDay(day, day_10.Part1, day_10.Part2, input)
	case "11":
		runDay(day, day_11.Part1, day_11.Part2, input)
	default:
		log.Fatalf("Unknown day: %s\n", day)
	}
}

func runDay(day string, part1, part2 func(string), input string) {
	fmt.Printf("===== %s =====\n", day)

	if part1 != nil {
		start := time.Now()
		fmt.Printf("Part 1: ")
		part1(input)
		elapsed := time.Since(start)
		fmt.Printf("\t(took %s)\n", elapsed)
	}

	if part2 != nil {
		start := time.Now()
		fmt.Printf("Part 2: ")
		part2(input)
		elapsed := time.Since(start)
		fmt.Printf("\t(took %s)\n", elapsed)
	}

	fmt.Println()
}
