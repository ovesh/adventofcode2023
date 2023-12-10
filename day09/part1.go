package day09

import (
	"strings"

	"ovesh.name/advent2023/utils"
)

func diffs(series []int) []int {
	res := make([]int, len(series)-1)
	for i := 1; i < len(series); i++ {
		res[i-1] = series[i] - series[i-1]
	}
	return res
}

func isAllZeros(series []int) bool {
	for _, item := range series {
		if item != 0 {
			return false
		}
	}
	return true
}

func copy(series []int) []int {
	res := make([]int, len(series))
	for i := range series {
		res[i] = series[i]
	}
	return res
}

func Part1() int {
	lines := strings.Split(sample, "\n")[1:]
	sum := 0
	for _, line := range lines {
		series := utils.SliceAtoi(strings.Split(line, " "))
		origs := [][]int{copy(series)}
		for !isAllZeros(series) {
			series = diffs(series)
			origs = append(origs, series)
		}
		for i := len(origs) - 1; i > 0; i-- {
			origs[i-1] = append(origs[i-1], origs[i-1][len(origs[i-1])-1]+origs[i][len(origs[i])-1])
		}
		sum += origs[0][len(origs[0])-1]
	}
	return sum
}
