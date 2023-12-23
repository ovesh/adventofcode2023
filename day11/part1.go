package day11

import (
	"fmt"
	"strings"
)

type coord struct {
	y, x int
}

func Part1() int {
	lines := strings.Split(sample, "\n")[1:]
	width := len(lines)
	matrix := make([][]rune, width)
	for i, line := range lines {
		matrix[i] = make([]rune, width)
		for j, c := range line {
			matrix[i][j] = c
		}
	}

	emptyRows := []int{}
	for i := 0; i < width; i++ {
		empty := true
		for j := 0; j < width; j++ {
			if matrix[i][j] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, i)
		}
	}

	emptyCols := []int{}
	for i := 0; i < width; i++ {
		empty := true
		for j := 0; j < width; j++ {
			if matrix[j][i] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}
	}
	//fmt.Println(emptyRows)
	//fmt.Println(emptyCols)

	galaxyCoords := []coord{}
	expanded := make([][]rune, width+len(emptyRows))
	emptyRowsIdx := 0
	for i := 0; i < width; i++ {
		emptyColIdx := 0
		expanded[i+emptyRowsIdx] = make([]rune, width+len(emptyCols))
		for j := 0; j < width; j++ {
			expanded[i+emptyRowsIdx][j+emptyColIdx] = matrix[i][j]
			if matrix[i][j] == '#' {
				galaxyCoords = append(galaxyCoords, coord{i + emptyRowsIdx, j + emptyColIdx})
			}
			if emptyColIdx < len(emptyCols) && emptyCols[emptyColIdx] == j {
				emptyColIdx++
				expanded[i+emptyRowsIdx][j+emptyColIdx] = '.'
			}
		}
		if emptyRowsIdx < len(emptyRows) && emptyRows[emptyRowsIdx] == i {
			emptyRowsIdx++
			expanded[i+emptyRowsIdx] = make([]rune, width+len(emptyCols))
			for j := 0; j < width+len(emptyCols); j++ {
				expanded[i+emptyRowsIdx][j] = '.'
			}
		}
	}

	sum := 0
	for i, pointa := range galaxyCoords {
		for j := i + 1; j < len(galaxyCoords); j++ {
			pointb := galaxyCoords[j]
			curDist := distance(pointa, pointb)
			//fmt.Println("distance", pointa, pointb, curDist)
			sum += curDist
		}
	}

	return sum
}

func distance(a, b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func draw(matrix [][]rune) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%c", matrix[i][j])
		}
		fmt.Println()
	}
}
