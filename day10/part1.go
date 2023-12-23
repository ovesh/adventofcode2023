package day10

import (
	"fmt"
	"strings"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

const (
	// sample1
	//sY, sX = 1, 1
	//sPipe  = 'F'
	//firstDir = up

	// sample2
	//sY, sX   = 2, 0
	//sPipe    = 'F'
	//firstDir = up

	// sample3, sample 4
	//sY, sX   = 1, 1
	//sPipe    = 'F'
	//firstDir = up

	// sample5
	//sY, sX   = 4, 12
	//sPipe    = 'F'
	//firstDir = left

	// sample6
	//sY, sX   = 0, 4
	//sPipe    = '7'
	//firstDir = right

	// input
	sY, sX   = 86, 89
	sPipe    = '-'
	firstDir = left
)

func Part1() int {
	lines := strings.Split(sample6, "\n")[1:]
	width := len(lines)
	matrix := make([][]rune, width)
	for i := 0; i < width; i++ {
		matrix[i] = make([]rune, width)
		for j, c := range lines[i] {
			matrix[i][j] = c
		}
	}
	matrix[sY][sX] = sPipe

	y, x := sY, sX
	lastDir := firstDir
	stepCount := 0
	for {
		switch matrix[y][x] {
		case 'F':
			switch lastDir {
			case up:
				x++
				lastDir = right
			case left:
				y++
				lastDir = down
			}
		case '|':
			switch lastDir {
			case up:
				y--
				lastDir = up
			case down:
				y++
				lastDir = down
			}
		case '7':
			switch lastDir {
			case right:
				y++
				lastDir = down
			case up:
				x--
				lastDir = left
			}
		case 'J':
			switch lastDir {
			case right:
				y--
				lastDir = up
			case down:
				x--
				lastDir = left
			}
		case 'L':
			switch lastDir {
			case down:
				x++
				lastDir = right
			case left:
				y--
				lastDir = up
			}
		case '-':
			switch lastDir {
			case right:
				x++
				lastDir = right
			case left:
				x--
				lastDir = left
			}
		default:
			panic(fmt.Sprintf("unknown: %s", string(matrix[y][x])))
		}
		stepCount++
		if y == sY && x == sX {
			break
		}
	}

	if stepCount%2 == 0 {
		return stepCount / 2
	}
	return (stepCount / 2) + 1
}
