package day02

import (
	"strings"

	"ovesh.name/advent2023/utils"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Part2() int {
	lines := strings.Split(sample, "\n")
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, ":")

		reveals := strings.Split(split[1], "; ")
		maxPerColor := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, reveal := range reveals {
			sets := strings.Split(strings.TrimSpace(reveal), ", ")
			for _, set := range sets {
				perColor := strings.Split(set, " ")
				count := utils.MustAtoi(perColor[0])
				color := perColor[1]
				maxPerColor[color] = max(maxPerColor[color], count)
			}
		}
		product := 1
		for _, count := range maxPerColor {
			product *= count
		}
		sum += product
	}
	return sum
}
