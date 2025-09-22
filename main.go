package main

import (
	"log"
	"os"
	"strings"
)

// Returns the lines of the file
func ExtractFile(location string) []string {
	b, err := os.ReadFile(location)

	if err != nil {
		log.Fatal(err)
	}
	file := string(b)
	return strings.Split(file, "\n")
}

func main() {
	// Solve()
	// SolveD2()
	SolveD3()
}
