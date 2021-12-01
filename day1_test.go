package aoc2021

import (
	"fmt"
	"testing"
)

func TestCountIncreases(t *testing.T) {
	fmt.Printf("# of increases: %d\n", CountIncreases("input/day1-1.txt"))
}

func TestCountWindowIncreases(t *testing.T) {
	fmt.Printf("# of window increases: %d\n", CountWindowIncreases("input/day1-1.txt"))
}
