package utils

import "strconv"

func MustAtoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func SliceAtoi(slice []string) []int {
	res := make([]int, len(slice))
	for i, s := range slice {
		res[i] = MustAtoi(s)
	}

	return res
}
