package day13

import (
	"strings"
)

func Part1() int {
	patterns := strings.Split(sample[1:], "\n\n")
	sum := 0
	for _, pattern := range patterns {
		horizontals := strings.Split(pattern, "\n")
		verticals := getVerticalLines(horizontals)
		for i := 0; i < len(verticals)-1; i++ {
			if verticals[i] == verticals[i+1] {
				symmetrical := checkSymmetry(verticals, i)
				if symmetrical {
					sum += i + 1
				}
				//fmt.Println("vertical", symmetrical, i)
			}
		}
		for i := 0; i < len(horizontals)-1; i++ {
			if horizontals[i] == horizontals[i+1] {
				symmetrical := checkSymmetry(horizontals, i)
				if symmetrical {
					sum += 100 * (i + 1)
				}
				//fmt.Println("horizontal", symmetrical, i)
			}
		}
	}
	return sum
}

func getVerticalLines(pattern []string) []string {
	length := len(pattern[0])
	res := make([]string, length)
	for j := 0; j < length; j++ {
		for _, line := range pattern {
			res[j] += string(line[j])
		}
	}
	return res
}

func checkSymmetry(pattern []string, center int) bool {
	for i := 0; center+i+1 < len(pattern) && center-i >= 0; i++ {
		if pattern[center+i+1] != pattern[center-i] {
			return false
		}
	}
	return true
}
