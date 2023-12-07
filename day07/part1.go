package day07

import (
	"fmt"
	"sort"
	"strings"

	"ovesh.name/advent2023/utils"
)

type handType int

const (
	fiveOfAKind handType = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPairs
	onePair
	highCard
)

var cardOrder = map[uint8]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'J': 3,
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
}

func getHandType(s string) handType {
	counts := map[rune]int{}
	for _, c := range s {
		if _, ok := counts[c]; !ok {
			counts[c] = 1
		} else {
			counts[c] += 1
		}
	}
	if len(counts) == 1 {
		return fiveOfAKind
	}
	if len(counts) == 5 {
		return highCard
	}
	if len(counts) == 4 {
		return onePair
	}
	if len(counts) == 2 {
		for _, num := range counts { // only need one iteration
			if num == 3 || num == 2 {
				return fullHouse
			}
			return fourOfAKind
		}
	}
	// len(counts) == 3
	for _, num := range counts {
		if num == 3 {
			return threeOfAKind
		}
		if num == 2 {
			return twoPairs
		}
	}
	panic(fmt.Sprintf("failed to determine for %s", s))
}

type hand struct {
	handType handType
	s        string
	bid      int
}

func Part1() int {
	hands := []hand{}
	lines := strings.Split(sample, "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")
		hands = append(hands, hand{handType: getHandType(split[0]), s: split[0], bid: utils.MustAtoi(split[1])})
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType < hands[j].handType {
			return false
		}
		if hands[i].handType > hands[j].handType {
			return true
		}
		for k := 0; k < 5; k++ {
			if cardOrder[hands[i].s[k]] < cardOrder[hands[j].s[k]] {
				return false
			}
			if cardOrder[hands[i].s[k]] > cardOrder[hands[j].s[k]] {
				return true
			}
		}
		return true
	})

	sum := 0
	for i, h := range hands {
		sum += (h.bid * (i + 1))
	}
	return sum
}
