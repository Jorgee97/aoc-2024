package main

import "testing"

func TestDay3P1(t *testing.T) {
	corruped := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	got := Day3P1(corruped)

	want := 161

	if got != want {
		t.Errorf("Got: %d, want: %d\n", got, want)
	}
}

// Part two enables or disables depending on the function do() or don't()
func TestDay3P2(t *testing.T) {
	corruped := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))?don't()!'$ what()where(353,94)select(){mul(997,311)@from()/mul(987,583)&select(207,730)mul(299,379)select()"

	got := Day3P2(corruped)

	want := 48

	if got != want {
		t.Errorf("Got: %d, want: %d\n", got, want)
	}
}
