package main

import (
	"fmt"
	"math"
	"path"
	"slices"
	"strconv"
	"strings"
)

func ProcessDay1File() ([]int, []int) {
	location := path.Join("inputs", "1")
	lines := ExtractFile(location)
	var left []int
	var right []int
	for _, v := range lines {
		parts := strings.Split(v, "   ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func SolveDay1(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	differences := make([]int, len(left))
	for i := range len(left) {
		vi := left[i]
		vj := right[i]
		v := math.Abs(float64(vi - vj))
		differences = append(differences, int(v))
	}

	var response int
	for _, v := range differences {
		response += v
	}
	return response
}

func Day1PartTwo(left, right []int) int {
	o := make(map[int]int)

	for _, v := range right {
		o[v] += 1
	}

	var result int
	for _, v := range left {
		result += v * o[v]
	}

	return result
}

func Solve() {
	left, right := ProcessDay1File()
	d1r := SolveDay1(left, right)
	fmt.Printf("Response Part 1: %d\n", d1r)

	d1r = Day1PartTwo(left, right)
	fmt.Printf("Response Part 2: %d\n", d1r)
}
