package day05

import (
	"sort"
	"strings"

	"ovesh.name/advent2023/utils"
)

type valWithRange struct {
	minKey, minVal, rangeLen int
}

type rangedMap struct {
	ranges []valWithRange
}

func (d *rangedMap) add(v valWithRange) {
	d.ranges = append(d.ranges, v)
	sort.Slice(d.ranges, func(i, j int) bool {
		return d.ranges[i].minKey < d.ranges[j].minKey
	})
}

func (d *rangedMap) get(k int) int {
	for _, r := range d.ranges {
		if k >= r.minKey && k < r.minKey+r.rangeLen {
			return r.minVal + (k - r.minKey)
		}
	}

	return k
}

func buildMap(s string) *rangedMap {
	lines := strings.Split(strings.TrimSpace(s), "\n")[1:]
	ranges := []valWithRange{}
	for _, line := range lines {
		vals := utils.SliceAtoi(strings.Split(line, " "))
		minVal := vals[0]
		minKey := vals[1]
		rangeLen := vals[2]
		ranges = append(ranges, valWithRange{minKey: minKey, minVal: minVal, rangeLen: rangeLen})
	}

	return &rangedMap{ranges: ranges}
}

func variadicMin(nums []int) int {
	curMin := nums[0]
	for _, cur := range nums {
		if cur < curMin {
			curMin = cur
		}
	}
	return curMin
}

func Part1() int {
	parts := strings.Split(sample, "\n\n")
	seedsPart := strings.Split(parts[0], "\n")[1:]
	seeds := utils.SliceAtoi(strings.Split(strings.Split(seedsPart[0], ": ")[1], " "))

	seed2Soil := buildMap(parts[1])
	soil2Fertilizer := buildMap(parts[2])
	fertilizer2Water := buildMap(parts[3])
	water2Light := buildMap(parts[4])
	light2Temp := buildMap(parts[5])
	temp2Humidity := buildMap(parts[6])
	humidity2Location := buildMap(parts[7])

	locations := []int{}
	for _, seed := range seeds {
		soil := seed2Soil.get(seed)
		fertilizer := soil2Fertilizer.get(soil)
		water := fertilizer2Water.get(fertilizer)
		light := water2Light.get(water)
		temp := light2Temp.get(light)
		humidity := temp2Humidity.get(temp)
		location := humidity2Location.get(humidity)
		locations = append(locations, location)
	}

	return variadicMin(locations)
}
