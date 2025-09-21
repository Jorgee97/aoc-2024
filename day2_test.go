package main

import "testing"

func TestDay2P1(t *testing.T) {
	// Arange
	data := [][]int{
		{1, 4, 5, 8, 11, 12, 9},
		{47, 46, 49, 52, 55},
		{84, 85, 86, 87, 90, 89, 86},
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	// Act
	got := Day2P1(data)
	want := 2

	if got != want {
		t.Errorf("Error got %d, wanted: %d\n", got, want)
	}
}

func TestDay2P2(t *testing.T) {
	// Arange
	data := [][]int{
		{54, 55, 58, 58, 59, 56},
		{19, 20, 23, 27, 28, 30, 33, 36},
		{65, 64, 63, 65, 66, 67, 70},
		{4, 3, 4, 5, 8, 15, 16, 14},
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	// Act
	got := Day2P2(data)
	want := 4

	if got != want {
		t.Errorf("Error got %d, wanted: %d\n", got, want)
	}
}
