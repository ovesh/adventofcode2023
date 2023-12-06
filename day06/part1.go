package day06

import (
	"math"
	"strings"

	"ovesh.name/advent2023/utils"
)

func Part1() int {
	lines := strings.Split(sample, "\n")
	totTimes := utils.SliceAtoi(strings.Fields(lines[0])[1:])
	minDists := utils.SliceAtoi(strings.Fields(lines[1])[1:])
	product := 1
	for i := 0; i < len(minDists); i++ {
		// distance travelled: (totTime - pressTime) * pressTime
		// needs to be larger than minDist:
		// (totTime - pressTime) * presstime > minDist
		// =>
		// (-1) * (pressTime^2) + totTime * pressTime - minDist > 0
		// which is a quadratic equation
		// count all the int solutions
		minDist := float64(minDists[i])
		totTime := float64(totTimes[i])
		sqrtDiscriminant := math.Sqrt(math.Pow(totTime, 2.0) - 4.0*minDist)
		sol1 := (-totTime + sqrtDiscriminant) / -2.0
		sol2 := (-totTime - sqrtDiscriminant) / -2.0
		var numSolutions int
		if float64(int(sol2)) == sol2 {
			numSolutions = int(sol2) - 1 - int(math.Ceil(sol1))
		} else {
			numSolutions = int(math.Floor(sol2)) - int(math.Ceil(sol1)) + 1
		}
		product *= numSolutions
	}
	return product
}
