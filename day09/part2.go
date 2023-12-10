package day09

import (
	"strings"

	"ovesh.name/advent2023/utils"
)

func Part2() int {
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
			origs[i-1] = append([]int{origs[i-1][0] - origs[i][0]}, origs[i-1]...)
		}
		sum += origs[0][0]
	}
	return sum
}
