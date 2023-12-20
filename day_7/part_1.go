package day_7

import (
	"advent_of_code_2023/utils"
	"strings"
)

type CamelHand struct {
	hand string
	bid  int
}

type typeRank string

// Define the possible typeRanks
const (
	FiveOfAKind  typeRank = "5OK"
	FourOfAKind  typeRank = "4OK"
	FullHouse    typeRank = "FH"
	ThreeOfAKind typeRank = "3OK"
	TwoPair      typeRank = "2P"
	OnePair      typeRank = "1P"
	HighCard     typeRank = "HC"
)

var typeRanking = map[typeRank]int{
	FiveOfAKind:  6,
	FourOfAKind:  5,
	FullHouse:    4,
	ThreeOfAKind: 3,
	TwoPair:      2,
	OnePair:      1,
	HighCard:     0,
}

func Part1(input string) {
	lines := utils.SplitLines(input)

	var camelHands []CamelHand

	for _, line := range lines {
		camelHands = append(camelHands, convertStringsToCamelHand(line))
	}

	sortedCards := sortCards(camelHands)

	var result = 0
	for i, card := range sortedCards {

		result = result + card.bid*(i+1)
	}

	println("The result is: ", result)
}

func convertStringsToCamelHand(line string) CamelHand {
	line1, line2, _ := strings.Cut(line, " ")

	return CamelHand{
		hand: line1,
		bid:  utils.ParseInt(line2),
	}

}

func sortCards(unsortedCards []CamelHand) []CamelHand {
	sortedCards := mergeSort(unsortedCards)
	return sortedCards
}

func mergeSort(items []CamelHand) []CamelHand {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(handA []CamelHand, handB []CamelHand) []CamelHand {
	var final []CamelHand
	i := 0
	j := 0
	for i < len(handA) && j < len(handB) {
		if isHandALessThanHandB(handA[i], handB[j]) {
			final = append(final, handA[i])
			i++
		} else {
			final = append(final, handB[j])
			j++
		}
	}
	for ; i < len(handA); i++ {
		final = append(final, handA[i])
	}
	for ; j < len(handB); j++ {
		final = append(final, handB[j])
	}
	return final
}

func isHandALessThanHandB(handA CamelHand, handB CamelHand) bool {
	var cardRanking = map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	handARank := getHandRanking(handA)
	handBRank := getHandRanking(handB)

	if handARank == handBRank {
		for i := 0; i < len(handA.hand); i++ {
			if cardRanking[string(handA.hand[i])] == cardRanking[string(handB.hand[i])] {
				continue
			}
			return cardRanking[string(handA.hand[i])] < cardRanking[string(handB.hand[i])]

		}
	}

	return typeRanking[handARank] < typeRanking[handBRank]
}

func getHandRanking(hand CamelHand) typeRank {
	cardCounts := make(map[rune]int)
	for _, card := range hand.hand {
		cardCounts[card]++
	}

	// Check for different hand types
	switch {
	case hasCount(cardCounts, 5):
		return FiveOfAKind
	case hasCount(cardCounts, 4):
		return FourOfAKind
	case hasCount(cardCounts, 3) && hasCount(cardCounts, 2):
		return FullHouse
	case hasCount(cardCounts, 3):
		return ThreeOfAKind
	case hasCount(cardCounts, 2) && countPairs(cardCounts) == 2:
		return TwoPair
	case hasCount(cardCounts, 2):
		return OnePair
	default:
		return HighCard
	}
}

func hasCount(counts map[rune]int, target int) bool {
	for _, count := range counts {
		if count == target {
			return true
		}
	}
	return false
}

func countPairs(counts map[rune]int) int {
	pairCount := 0
	for _, count := range counts {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount
}
