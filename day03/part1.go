package day03

import (
	"fmt"
	"strings"

	"ovesh.name/advent2023/utils"
)

func print(m [][]rune) {
	for _, row := range m {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

type number struct {
	y, xStart, xEnd, val int
	adjacent             bool
}

func (n *number) String() string {
	return fmt.Sprintf("%+v", *n)
}

var symbols = map[rune]struct{}{
	'*': {},
	'$': {},
	'#': {},
	'+': {},
	'&': {},
	'%': {},
	'-': {},
	'@': {},
	'=': {},
	'/': {},
}

func markAdjacent(symbolX int, numbers []*number) {
	for _, n := range numbers {
		if symbolX >= n.xStart-1 && symbolX <= n.xEnd+1 {
			n.adjacent = true
		}
	}
}

func Part1() int {
	lines := strings.Split(sample, "\n")
	lines = lines[1:]
	sum := 0
	matrix := make([][]rune, len(lines))
	numbersPerRow := make([][]*number, len(matrix))
	for i, line := range lines {
		rLine := []rune(line)
		matrix[i] = rLine
		numbersPerRow[i] = []*number{}
		curStart := -1
		curEnd := -1
		for j, c := range rLine {
			if c >= '0' && c <= '9' {
				if curStart == -1 {
					curStart = j
				}
				curEnd = j
				if j == len(rLine)-1 {
					val := utils.MustAtoi(line[curStart : curEnd+1])
					numbersPerRow[i] = append(numbersPerRow[i], &number{y: i, xStart: curStart, xEnd: curEnd, val: val})
				}
			} else {
				if curEnd != -1 {
					val := utils.MustAtoi(line[curStart : curEnd+1])
					numbersPerRow[i] = append(numbersPerRow[i], &number{y: i, xStart: curStart, xEnd: curEnd, val: val})
					curStart = -1
					curEnd = -1
				}
			}
		}
	}

	for i, rLine := range matrix {
		for j, c := range rLine {
			if c != '.' && !(c >= '0' && c <= '9') {
				if _, ok := symbols[c]; !ok {
					panic(fmt.Sprintf("unknown symbol %s", string(c)))
				}

				if i > 0 {
					markAdjacent(j, numbersPerRow[i-1])
				}
				markAdjacent(j, numbersPerRow[i])
				if i < len(numbersPerRow)-1 {
					markAdjacent(j, numbersPerRow[i+1])
				}
			}
		}
	}

	for _, perRow := range numbersPerRow {
		for _, n := range perRow {
			if n.adjacent {
				sum += n.val
			}
		}
	}

	return sum
}
