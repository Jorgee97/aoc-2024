package main

import (
	"testing"
)

func TestSolveDay1(t *testing.T) {
	// Arrange
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	// Act
	got := SolveDay1(left, right)

	// Assert
	want := 11
	if got != 11 {
		t.Errorf("got %d, want: %d", got, want)
	}
}

func TestSolveDay1Part2(t *testing.T) {
	// Arrange
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	// Act
	got := Day1PartTwo(left, right)

	// Assert
	want := 31
	if got != want {
		t.Errorf("got %d, want: %d", got, want)
	}
}
