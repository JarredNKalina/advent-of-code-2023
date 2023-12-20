package day_4

import (
	"fmt"
	"regexp"
	"strings"
)

type Card struct {
	leftNumbers  map[string]int
	rightNumbers map[string]int
	totalCount   int
}

func Part2(input string) {

	var cards []Card

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		card := processLineToCard(line)
		cards = append(cards, card)
	}

	for i, card := range cards {
		points := getPointsForCard(card)

		for j := 1; j <= points; j++ {
			cards[i+j].totalCount += 1 * cards[i].totalCount
		}
	}

	totalCards := 0
	for _, card := range cards {
		totalCards += card.totalCount
	}

	fmt.Printf("There are %d total cards in the stack!\n", totalCards)
}

func getPointsForCard(card Card) int {
	points := 0
	for copy, count := range card.rightNumbers {
		if winningCount, ok := card.leftNumbers[copy]; ok {
			points += count * winningCount
		}
	}
	return points
}

func processLineToCard(line string) Card {
	_, cardDataStr, _ := strings.Cut(line, ": ")
	cardData := strings.Split(cardDataStr, " | ")

	re := regexp.MustCompile("[0-9]{1,2}")

	leftNumbers := make(map[string]int)
	for _, point := range re.FindAllString(cardData[0], -1) {
		leftNumbers[point]++
	}

	rightNumbers := make(map[string]int)
	for _, point := range re.FindAllString(cardData[1], -1) {
		rightNumbers[point]++
	}

	return Card{
		leftNumbers:  leftNumbers,
		rightNumbers: rightNumbers,
		totalCount:   1,
	}
}
