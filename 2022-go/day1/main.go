package main

import (
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	max := 0
	running := 0

	for _, line := range lines {
		if line == "" {
			if running > max {
				max = running
			}
			running = 0
			continue
		}

		// string to int
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		running += num
	}
	if running > max {
		max = running
	}
	return max
}

func minIndex(arr []int) int {
	min := 0
	for i, v := range arr {
		if v < arr[min] {
			min = i
		}
	}
	return min
}

func part2(lines []string) int {
	tops := make([]int, 3)
	running := 0

	for _, line := range lines {
		if line == "" {
			min := minIndex(tops)
			if running > tops[min] {
				tops[min] = running
			}
			running = 0
			continue
		}

		// string to int
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		running += num
	}
	sum := 0
	for _, v := range tops {
		sum += v
	}
	return sum
}

func main() {
	dat, err := os.ReadFile("./day1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")

	println("Part 1:", part1(lines))
	println("Part 2:", part2(lines))
}
