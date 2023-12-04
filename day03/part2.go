package day03

import (
	"strings"

	"ovesh.name/advent2023/utils"
)

type number2 struct {
	y, xStart, xEnd, val int
}

type gear struct {
	y, x      int
	neighbors map[number2]struct{}
}

func findNeighbors(g *gear, numbers []*number2) {
	for _, n := range numbers {
		if g.x >= n.xStart-1 && g.x <= n.xEnd+1 {
			g.neighbors[*n] = struct{}{}
		}
	}
}

func Part2() int {
	lines := strings.Split(sample, "\n")
	lines = lines[1:]
	sum := 0
	matrix := make([][]rune, len(lines))
	numbersPerRow := make([][]*number2, len(matrix))
	gears := []*gear{}
	for i, line := range lines {
		rLine := []rune(line)
		matrix[i] = rLine
		numbersPerRow[i] = []*number2{}
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
					numbersPerRow[i] = append(numbersPerRow[i], &number2{y: i, xStart: curStart, xEnd: curEnd, val: val})
				}
			} else {
				if curEnd != -1 {
					val := utils.MustAtoi(line[curStart : curEnd+1])
					numbersPerRow[i] = append(numbersPerRow[i], &number2{y: i, xStart: curStart, xEnd: curEnd, val: val})
					curStart = -1
					curEnd = -1
				}
				if c == '*' {
					gears = append(gears, &gear{x: j, y: i, neighbors: map[number2]struct{}{}})
				}
			}
		}
	}

	for _, g := range gears {
		if g.y > 0 {
			findNeighbors(g, numbersPerRow[g.y-1])
		}
		findNeighbors(g, numbersPerRow[g.y])
		if g.y < len(numbersPerRow)-1 {
			findNeighbors(g, numbersPerRow[g.y+1])
		}

		if len(g.neighbors) == 2 {
			ns := [2]int{}
			i := 0
			for n := range g.neighbors {
				ns[i] = n.val
				i++
			}

			sum += (ns[0] * ns[1])
		}
	}

	return sum
}
