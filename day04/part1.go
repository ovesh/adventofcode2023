package day04

import (
	"strings"

	"ovesh.name/advent2023/utils"
)

type intset map[int]struct{}

func sliceAtoi(slice []string) intset {
	res := intset{}
	for _, s := range slice {
		res[utils.MustAtoi(s)] = struct{}{}
	}

	return res
}

func intersect(a, b intset) intset {
	res := intset{}
	for i, _ := range a {
		if _, ok := b[i]; ok {
			res[i] = struct{}{}
		}
	}
	return res
}

func pow2(exponent int) int {
	res := 1
	for i := 0; i < exponent; i++ {
		res *= 2
	}
	return res
}

func Part1() int {
	lines := strings.Split(sample, "\n")
	lines = lines[1:]
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, ": ")
		split = strings.Split(split[1], " | ")
		winners := sliceAtoi(strings.Fields(split[0]))
		existing := sliceAtoi(strings.Fields(split[1]))
		intersection := intersect(winners, existing)
		if len(intersection) > 0 {
			sum += pow2(len(intersection) - 1)
		}
	}

	return sum
}
