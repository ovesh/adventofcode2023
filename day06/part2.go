package day06

import (
	"math"
	"strings"

	"ovesh.name/advent2023/utils"
)

func Part2() int {
	lines := strings.Split(sample, "\n")
	totTime := float64(utils.MustAtoi(strings.Join(strings.Fields(lines[0])[1:], "")))
	minDist := float64(utils.MustAtoi(strings.Join(strings.Fields(lines[1])[1:], "")))
	// distance travelled: (totTime - pressTime) * pressTime
	// needs to be larger than minDist:
	// (totTime - pressTime) * presstime > minDist
	// =>
	// (-1) * (pressTime^2) + totTime * pressTime - minDist > 0
	// which is a quadratic equation
	// count all the int solutions
	sqrtDiscriminant := math.Sqrt(math.Pow(totTime, 2.0) - 4.0*minDist)
	sol1 := (-totTime + sqrtDiscriminant) / -2.0
	sol2 := (-totTime - sqrtDiscriminant) / -2.0
	if float64(int(sol2)) == sol2 {
		return int(sol2) - 1 - int(math.Ceil(sol1))
	}
	return int(math.Floor(sol2)) - int(math.Ceil(sol1)) + 1
}
