package aoc2021

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput(path string) []int {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var out []int
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, depth)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return out
}

func CountIncreases(path string) int {
	depths := readInput(path)
	count := 0
	i := 1
	for i < len(depths) {
		if depths[i] > depths[i-1] {
			count++
		}
		i++
	}
	return count
}

func CountWindowIncreases(path string) int {
	depths := readInput(path)
	count := 0
	var sum_prev int
	sum_curr := depths[0] + depths[1] + depths[2]
	a, b, c := 1, 2, 3
	for c < len(depths) {
		sum_curr, sum_prev = (depths[a] + depths[b] + depths[c]), sum_curr
		if sum_curr > sum_prev {
			count++
		}
		a++
		b++
		c++
	}
	return count
}
