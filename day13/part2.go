package day13

import (
	"strings"
)

func Part2() int {
	patterns := strings.Split(sample[1:], "\n\n")
	sum := 0
	for _, pattern := range patterns {
		//fmt.Println(pattern)
		horizontals := strings.Split(pattern, "\n")
		verticals := getVerticalLines(horizontals)
		for i := 0; i < len(verticals)-1; i++ {
			diffs := strcmpWithErr(verticals[i], verticals[i+1])
			if diffs == 0 || diffs == 1 {
				symmetrical := checkSymmetryWithErr(verticals, i)
				if symmetrical {
					sum += i + 1
				}
				//fmt.Println("vertical", symmetrical, i)
			}
		}
		for i := 0; i < len(horizontals)-1; i++ {
			diffs := strcmpWithErr(horizontals[i], horizontals[i+1])
			if diffs == 0 || diffs == 1 {
				symmetrical := checkSymmetryWithErr(horizontals, i)
				if symmetrical {
					sum += 100 * (i + 1)
				}
				//fmt.Println("horizontal", symmetrical, i)
			}
		}
	}
	return sum
}

func checkSymmetryWithErr(pattern []string, center int) bool {
	diffs := 0
	for i := 0; center+i+1 < len(pattern) && center-i >= 0; i++ {
		diffs += strcmpWithErr(pattern[center+i+1], pattern[center-i])
		if diffs > 1 {
			return false
		}
	}
	return diffs == 1
}

func strcmpWithErr(a, b string) int {
	diffs := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++
		}
		if diffs > 1 {
			return diffs
		}
	}
	return diffs
}
