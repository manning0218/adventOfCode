package main

import (
	"fmt"

	"github.com/manning0218/adventOfCode/2025/go/day01"
	"github.com/manning0218/adventOfCode/2025/go/utils"
)

func main() {
	input := utils.ReadInput(1)
	part1Result := part1(input)
	fmt.Printf("Part 1 Result: %d\n", part1Result)
	part2Result := part2(input)
	fmt.Printf("Part 2 Result: %d\n", part2Result)
}

func part1(input []string) int {
	count := 0
	lock := day01.NewLock(100, 50)
	for _, line := range input {
		var direction rune
		var steps int
		fmt.Sscanf(line, "%c%d", &direction, &steps)
		if direction == 'L' {
			lock = lock.MoveLeft(steps)
		} else if direction == 'R' {
			lock = lock.MoveRight(steps)
		}

		if lock.IsMagicNumber(0) {
			count++
		}
	}

	return count
}

func part2(input []string) int {
	day01.ResetPassword()
	part1(input)
	return day01.GetPassword()
}
