package day11

import (
	"strings"
)

const (
	expansion = 100
)

func Part2() int {
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
	emptyRowsIdx := 0
	for i := 0; i < width; i++ {
		emptyColIdx := 0
		for j := 0; j < width; j++ {
			if matrix[i][j] == '#' {
				galaxyCoords = append(galaxyCoords, coord{i + emptyRowsIdx*(expansion-1), j + emptyColIdx*(expansion-1)})
			}
			if emptyColIdx < len(emptyCols) && emptyCols[emptyColIdx] == j {
				emptyColIdx++
			}
		}
		if emptyRowsIdx < len(emptyRows) && emptyRows[emptyRowsIdx] == i {
			emptyRowsIdx++
		}
	}
	//fmt.Println(galaxyCoords)

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

//type expandedMatrix struct {
//	orig                 [][]rune
//	emptyRows, emptyCols []int
//}
//
//func colsBetween(orig [][]rune, emptyRows, emptyCols []int, a, b coord) int {
//
//}
