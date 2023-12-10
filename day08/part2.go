package day08

import (
	"fmt"
	"strings"
)

func getFirstZ(n node, instructions string, nodesByID map[string]node) int64 {
	curNode := n
	for i := int64(0); true; i++ {
		instructionIdx := i % int64(len(instructions))
		instruction := instructions[instructionIdx]
		if instruction == 'L' {
			curNode = nodesByID[curNode.left]
		} else {
			curNode = nodesByID[curNode.right]
		}
		if curNode.id[2] == 'Z' {
			return i
		}
	}
	panic("no")
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Part2() int {
	split := strings.Split(sample3, "\n\n")
	instructions := split[0]
	sNodes := strings.Split(split[1], "\n")
	nodesByID := map[string]node{}
	curNodes := []node{}
	for _, sNode := range sNodes {
		n := node{id: sNode[:3], left: sNode[7:10], right: sNode[12:15]}
		nodesByID[sNode[:3]] = n
		if n.id[2] == 'A' {
			curNodes = append(curNodes, n)
		}
	}

	loopSizes := []int64{}
	// avoid a brute force solution. Each of the ghost paths settles on a loop. Every loop has a different length
	// so need to find the lowest common multiple of the cycle lengths.
	// However, the below code is not a generalized solution, but rather relies 2 idiosyncracies of the input data:
	// (1) the distance from the first node to the beginning of a loop is the same as from the first Z to the
	// beginning of the loop.
	// (2) each cycle contains only one Z
	for _, curNode := range curNodes {
		z := getFirstZ(curNode, instructions, nodesByID)
		loopSizes = append(loopSizes, z+1)
		//fmt.Printf("%+v %d %+v\n", curNode, z+1)
	}
	fmt.Println(loopSizes)
	fmt.Println(LCM(loopSizes[0], loopSizes[1], loopSizes[2:]...))

	return 0
}
