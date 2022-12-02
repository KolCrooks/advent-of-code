package main

import (
	"os"
	"strings"
)

func part1(lines []string) int {
	beats := map[byte]byte{
		'X': 'C',
		'Y': 'A',
		'Z': 'B',
	}

	score := 0
	for _, line := range lines {
		elf := line[0]
		me := line[2]
		score += int(me-'X') + 1
		if elf-'A' == me-'X' { // Draw
			score += 3
		} else if elf == beats[me] { // Win
			score += 6
		}

	}
	return score
}

func part2(lines []string) int {
	score_map := [][]int{
		/*	 ELF
			  A B C
		OUT	X 3 1 2
			Y 4 5 6
			Z 8 9 7
		*/
		{3, 1, 2},
		{4, 5, 6},
		{8, 9, 7},
	}

	score := 0
	for _, line := range lines {
		elf := line[0]
		outcome := line[2]
		score += score_map[outcome-'X'][elf-'A']
	}
	return score

}

func main() {
	dat, err := os.ReadFile("./day2.txt")
	if err != nil {
		panic(err)
	}
	// 	dat := []byte(`A Y
	// B X
	// C Z`)
	lines := strings.Split(string(dat), "\n")

	println("Part 1:", part1(lines))
	println("Part 2:", part2(lines))
}
