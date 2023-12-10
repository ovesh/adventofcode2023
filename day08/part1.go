package day08

import (
	"strings"
)

type node struct {
	id, left, right string
}

func Part1() int {
	split := strings.Split(sample1, "\n\n")
	instructions := split[0]
	sNodes := strings.Split(split[1], "\n")
	nodesByID := map[string]node{}
	for _, sNode := range sNodes {
		nodesByID[sNode[:3]] = node{id: sNode[:3], left: sNode[7:10], right: sNode[12:15]}
	}

	curNode := nodesByID["AAA"]
	for i := 0; true; i++ {
		instruction := instructions[i%len(instructions)]
		if instruction == 'L' {
			curNode = nodesByID[curNode.left]
		} else {
			curNode = nodesByID[curNode.right]
		}
		if curNode.id == "ZZZ" {
			return i + 1
		}
	}
	return 0
}
