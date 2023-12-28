package day_11

import (
	"advent_of_code_2023/utils"
	"strings"

	"github.com/devries/combs"
)

func Part2(input string) {
	lines := strings.Split(input, "\n")

	galaxyRows := make(map[int]bool)
	galaxyColumns := make(map[int]bool)

	galaxies := []utils.Point{}

	for j, ln := range lines {
		for i, r := range ln {
			if r == '#' {
				p := utils.Point{X: i, Y: j}
				galaxies = append(galaxies, p)
				galaxyRows[j] = true
				galaxyColumns[i] = true
			}
		}
	}

	sum := 0

	for combo := range combs.Combinations(2, galaxies) {
		d := 0
		for i := min(combo[0].X, combo[1].X); i < max(combo[0].X, combo[1].X); i++ {
			if galaxyColumns[i] {
				d++
			} else {
				d += 1000000
			}
		}

		for j := min(combo[0].Y, combo[1].Y); j < max(combo[0].Y, combo[1].Y); j++ {
			if galaxyRows[j] {
				d++
			} else {
				d += 1000000
			}
		}

		sum += d
	}
	println("Answer:", sum)
}
