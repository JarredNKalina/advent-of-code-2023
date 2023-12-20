package day_7

import (
	"advent_of_code_2023/utils"
	"strings"
)

func Part2(input string) {
	lines := utils.SplitLines(input)

	var camelHands []CamelHand

	for _, line := range lines {
		camelHands = append(camelHands, convertStringsToCamelHandP2(line))
	}

	sortedCards := sortCardsP2(camelHands)

	var result = 0
	for i, card := range sortedCards {

		result = result + card.bid*(i+1)
	}

	println("The result is: ", result)
}

func convertStringsToCamelHandP2(line string) CamelHand {
	line1, line2, _ := strings.Cut(line, " ")

	return CamelHand{
		hand: line1,
		bid:  utils.ParseInt(line2),
	}

}

func sortCardsP2(unsortedCards []CamelHand) []CamelHand {
	sortedCards := mergeSortP2(unsortedCards)
	return sortedCards
}

func mergeSortP2(items []CamelHand) []CamelHand {
	if len(items) < 2 {
		return items
	}
	first := mergeSortP2(items[:len(items)/2])
	second := mergeSortP2(items[len(items)/2:])
	return mergeP2(first, second)
}

func mergeP2(handA []CamelHand, handB []CamelHand) []CamelHand {
	var final []CamelHand
	i := 0
	j := 0
	for i < len(handA) && j < len(handB) {
		if isHandALessThanHandBP2(handA[i], handB[j]) {
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

func isHandALessThanHandBP2(handA CamelHand, handB CamelHand) bool {
	var cardRanking = map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

	handARank := getHandRankingP2(handA)
	handBRank := getHandRankingP2(handB)

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

func getHandRankingP2(hand CamelHand) typeRank {
	cardCounts := make(map[rune]int)
	for _, card := range hand.hand {
		cardCounts[card]++
	}
	jokerCount := cardCounts['J']

	switch {
	case hasCountP2(cardCounts, 5):
		return FiveOfAKind
	case hasCountP2(cardCounts, 4):
		// JKKKK
		if jokerCount == 1 {
			return FiveOfAKind
		}
		// JJJJK
		if jokerCount == 4 {
			return FiveOfAKind
		}
		return FourOfAKind
	case hasCountP2(cardCounts, 3) && hasCountP2(cardCounts, 2):
		// JJKKK
		if jokerCount == 2 {
			return FiveOfAKind
		}
		// JJJKK
		if jokerCount == 3 {
			return FiveOfAKind
		}
		return FullHouse
	case hasCountP2(cardCounts, 3):
		//JKKK1
		if jokerCount == 1 {
			return FourOfAKind
		}
		// JJJK1
		if jokerCount == 3 {
			return FourOfAKind
		}
		return ThreeOfAKind
	case hasCountP2(cardCounts, 2) && countPairsP2(cardCounts) == 2:
		// J22KK
		if jokerCount == 1 {
			return FullHouse
		}
		// JJ2KK
		if jokerCount == 2 {
			return FourOfAKind
		}
		return TwoPair
	case hasCountP2(cardCounts, 2):
		// KKJ12
		if jokerCount == 1 {
			return ThreeOfAKind
		}
		if jokerCount == 2 {
			return ThreeOfAKind
		}
		return OnePair
	default:
		if jokerCount == 1 {
			return OnePair
		}
		return HighCard
	}
}

func hasCountP2(counts map[rune]int, target int) bool {
	for _, count := range counts {
		if count == target {
			return true
		}
	}
	return false
}

func countPairsP2(counts map[rune]int) int {
	pairCount := 0
	for _, count := range counts {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount
}
