package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func Day3P1(corrupted string) int {
	r, _ := regexp.Compile(`mul[(]\d{1,3},\d{1,3}[)]`)

	idxs := r.FindAllStringIndex(corrupted, -1)

	result := 0
	for _, v := range idxs {
		mul := corrupted[v[0]:v[1]]
		parts := strings.Split(mul, ",")
		xs, ys := parts[0][4:], parts[1][:len(parts[1])-1]
		x, _ := strconv.Atoi(xs)
		y, _ := strconv.Atoi(ys)
		result += x * y
	}

	return result
}

func inRange(idx int, idxs [][]int) bool {
	for _, v := range idxs {
		from, to := v[0], v[1]
		if idx > from && idx < to {
			return true
		}
	}
	return false
}

func Day3P2(corrupted string) int {
	r, _ := regexp.Compile(`mul[(]\d{1,3},\d{1,3}[)]`)
	do_not, _ := regexp.Compile(`don't\(\)([\s\S]*?)(do\(\)|$)`)

	idxs := r.FindAllStringIndex(corrupted, -1)
	dontidx := do_not.FindAllStringIndex(corrupted, -1)
	fmt.Println(len(dontidx))

	result := 0
	for _, v := range idxs {
		lastIdx := v[1]
		if inRange(lastIdx, dontidx) {
			continue
		}
		mul := corrupted[v[0]:v[1]]
		parts := strings.Split(mul, ",")
		xs, ys := parts[0][4:], parts[1][:len(parts[1])-1]
		x, _ := strconv.Atoi(xs)
		y, _ := strconv.Atoi(ys)
		result += x * y
	}

	return result
}

func SolveD3() {
	b, err := os.ReadFile(path.Join("inputs", "3"))

	if err != nil {
		log.Fatal(err)
	}
	data := string(b)

	r := Day3P1(data)
	fmt.Printf("Result day 3 part 1: %d\n", r)
	r2 := Day3P2(data)
	fmt.Printf("Result day 3 part 2: %d\n", r2)
}
