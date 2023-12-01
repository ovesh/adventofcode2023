package day01

import (
	"strings"
)

var digits = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func isDigit(s string, index int) rune {
	char := s[index]
	if char >= '0' && char <= '9' {
		return rune(char)
	}
	for k, v := range digits {
		if strings.HasPrefix(s[index:], k) {
			return v
		}
	}
	return rune(-1)
}

func Part2() int {
	lines := strings.Split(sample, "\n")
	sum := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		chars := []rune(line)
		var firstDigit, lastDigit rune
		for i := range chars {
			c := isDigit(line, i)
			if c != -1 {
				firstDigit = c
				break
			}
		}
		for i := len(chars) - 1; i >= 0; i-- {
			c := isDigit(line, i)
			if c != -1 {
				lastDigit = c
				break
			}
		}
		cur := (firstDigit-'0')*10 + (lastDigit - '0')
		println(cur)
		sum += int(cur)
	}
	return sum
}
