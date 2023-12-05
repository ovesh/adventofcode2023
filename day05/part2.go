package day05

import (
	"strings"

	"ovesh.name/advent2023/utils"
)

func Part2() int {
	parts := strings.Split(sample, "\n\n")

	seed2Soil := buildMap(parts[1])
	soil2Fertilizer := buildMap(parts[2])
	fertilizer2Water := buildMap(parts[3])
	water2Light := buildMap(parts[4])
	light2Temp := buildMap(parts[5])
	temp2Humidity := buildMap(parts[6])
	humidity2Location := buildMap(parts[7])

	res := -1
	seedsPart := strings.Split(parts[0], "\n")[1:]
	seedRanges := utils.SliceAtoi(strings.Split(strings.Split(seedsPart[0], ": ")[1], " "))
	for i := 0; i < len(seedRanges); i += 2 {
		start := seedRanges[i]
		rangeLen := seedRanges[i+1]
		for seed := start; seed < start+rangeLen; seed++ {
			soil := seed2Soil.get(seed)
			fertilizer := soil2Fertilizer.get(soil)
			water := fertilizer2Water.get(fertilizer)
			light := water2Light.get(water)
			temp := light2Temp.get(light)
			humidity := temp2Humidity.get(temp)
			location := humidity2Location.get(humidity)
			if res == -1 {
				res = location
			}
			res = min(res, location)
		}
	}

	return res
}
