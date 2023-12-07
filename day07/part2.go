package day07

import (
	"fmt"
	"sort"
	"strings"

	"ovesh.name/advent2023/utils"
)

var cardOrder2 = map[uint8]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'T': 3,
	'9': 4,
	'8': 5,
	'7': 6,
	'6': 7,
	'5': 8,
	'4': 9,
	'3': 10,
	'2': 11,
	'J': 12,
}

func getHandType2(s string) handType {
	counts := map[rune]int{}
	jokerNum := 0
	for _, c := range s {
		if c == 'J' {
			jokerNum += 1
		}
	}
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
		if jokerNum == 1 {
			return onePair
		}
		return highCard
	}
	if len(counts) == 4 {
		if jokerNum == 1 || jokerNum == 2 {
			return threeOfAKind
		}
		return onePair
	}
	if len(counts) == 2 {
		if jokerNum > 0 {
			return fiveOfAKind
		}
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
			if jokerNum == 3 || jokerNum == 1 {
				return fourOfAKind
			}
			if jokerNum == 2 {
				return fiveOfAKind
			}
			return threeOfAKind
		}
		if num == 2 {
			if jokerNum == 1 {
				return fullHouse
			}
			if jokerNum == 2 {
				return fourOfAKind
			}
			return twoPairs
		}
	}
	panic(fmt.Sprintf("failed to determine for %s", s))
}

func Part2() int {
	hands := []hand{}
	lines := strings.Split(sample, "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")
		hands = append(hands, hand{handType: getHandType2(split[0]), s: split[0], bid: utils.MustAtoi(split[1])})
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType < hands[j].handType {
			return false
		}
		if hands[i].handType > hands[j].handType {
			return true
		}
		for k := 0; k < 5; k++ {
			if cardOrder2[hands[i].s[k]] < cardOrder2[hands[j].s[k]] {
				return false
			}
			if cardOrder2[hands[i].s[k]] > cardOrder2[hands[j].s[k]] {
				return true
			}
		}
		return true
	})

	sum := 0
	for i, h := range hands {
		fmt.Printf("%+v\n", h)
		sum += (h.bid * (i + 1))
	}
	return sum
}
