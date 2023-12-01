package day01

import (
	"strings"
)

func Part1() int {
	lines := strings.Split(sample, "\n")
	sum := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		chars := []rune(line)
		var firstDigit, lastDigit rune
		for _, char := range chars {
			if char >= '0' && char <= '9' {
				firstDigit = char
				break
			}
		}
		for i := len(chars) - 1; i >= 0; i-- {
			char := chars[i]
			if char >= '0' && char <= '9' {
				lastDigit = char
				break
			}
		}
		cur := (firstDigit-'0')*10 + (lastDigit - '0')
		println(cur)
		sum += int(cur)
	}
	return sum
}
