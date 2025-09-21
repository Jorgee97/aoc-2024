package main

import (
	"fmt"
	"math"
	"path"
	"slices"
	"strconv"
	"strings"
)

func DataExtract() [][]int {
	location := path.Join("inputs", "2")
	lines := ExtractFile(location)
	var items [][]int
	for _, v := range lines {
		parts := strings.Split(v, " ")
		var line []int
		for _, vi := range parts {
			l, _ := strconv.Atoi(vi)
			line = append(line, l)
		}
		items = append(items, line)
	}
	return items
}

func Day2P1(data [][]int) int {
	safe := 0
	for _, v := range data {
		decreasing := false
		increasing := false
		// valid := false
		for j := range len(v) - 1 {
			if v[j] > v[j+1] {
				decreasing = true
			}

			if v[j] < v[j+1] {
				increasing = true
			}

			rule := math.Abs(float64(v[j]) - float64(v[j+1]))
			if rule > float64(3) || rule == 0 {
				decreasing = true
				increasing = true
			}
		}
		if increasing != decreasing {
			safe += 1
		}
	}

	return safe
}

func invalidDifference(a, b int) (bool, int) {
	rule := math.Abs(float64(a) - float64(b))
	if rule > float64(3) || rule == 0 {
		return true, int(rule)
	}
	return false, int(rule)
}

func isSafe(v []int) bool {
	decreasing := false
	increasing := false

	for j := range len(v) - 1 {
		if v[j] > v[j+1] {
			decreasing = true
		}

		if v[j] < v[j+1] {
			increasing = true
		}

		rule, _ := invalidDifference(v[j], v[j+1])
		if rule {
			decreasing = true
			increasing = true
		}
	}

	return decreasing != increasing
}

func Day2P2(data [][]int) int {
	safe := 0
	for _, v := range data {
		if isSafe(v) {
			safe += 1
		} else {
			for i := range len(v) {
				a := make([]int, len(v))
				copy(a, v)
				a = slices.Delete(a, i, i+1)
				if isSafe(a) {
					safe += 1
					break
				}
			}
		}
	}

	return safe
}

func SolveD2() {
	data := DataExtract()
	d2r := Day2P1(data)
	fmt.Printf("Response Part 1 Day 2: %d\n", d2r)

	d2r = Day2P2(data)
	fmt.Printf("Response Part 2 Day 2: %d\n", d2r)
}
