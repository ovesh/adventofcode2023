package day04

import (
	"strings"

	"ovesh.name/advent2023/utils"
)

type card struct {
	id                int
	winners, existing intset
}

func Part2() int {
	lines := strings.Split(sample, "\n")
	lines = lines[1:]
	cards := []card{}
	for _, line := range lines {
		split := strings.Split(line, ": ")
		cardID := utils.MustAtoi(strings.Fields(split[0])[1])
		split = strings.Split(split[1], " | ")
		winners := sliceAtoi(strings.Fields(split[0]))
		existing := sliceAtoi(strings.Fields(split[1]))
		cards = append(cards, card{id: cardID, winners: winners, existing: existing})
	}
	for i := 0; i < len(cards); i++ {
		intersection := intersect(cards[i].winners, cards[i].existing)
		firstID := cards[i].id
		for id := firstID; id < firstID+len(intersection); id++ {
			cards = append(cards, cards[id])
		}
	}

	return len(cards)
}
