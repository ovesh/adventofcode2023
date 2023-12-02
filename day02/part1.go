package day02

import (
	"strings"

	"ovesh.name/advent2023/utils"
)

var maxPerColor = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Part1() int {
	lines := strings.Split(sample, "\n")
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, ":")

		sID := split[0][5:]
		id := utils.MustAtoi(sID)

		reveals := strings.Split(split[1], "; ")
		valid := true
		for _, reveal := range reveals {
			sets := strings.Split(strings.TrimSpace(reveal), ", ")
			for _, set := range sets {
				perColor := strings.Split(set, " ")
				count := utils.MustAtoi(perColor[0])
				color := perColor[1]
				if maxPerColor[color] < count {
					valid = false
					break
				}
			}
		}
		if valid {
			sum += id
		}
	}
	return sum
}
