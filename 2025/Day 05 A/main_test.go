package main

import (
	"slices"
	"testing"
)

func TestHelloName(t *testing.T) {
	expected, _ := readLines("./out.txt")
	input, _ := readLines("./in.txt")

	actual := solve(input)
	if len(actual) == 0 {
		t.Fatal("No actual output from solve function")
	}
	if !slices.Equal(actual, expected) {
		t.Fatalf(`want %v, got %v`, expected, actual)
	}
}
