package main

import (
	"fmt"

	"github.com/manning0218/adventOfCode/2025/go/day03"
	"github.com/manning0218/adventOfCode/2025/go/utils"
)

func main() {
	input := utils.ReadInput(3)
	part1Result := part1(input)
	fmt.Printf("Part 1 Result: %d\n", part1Result)
	part2Result := part2(input)
	fmt.Printf("Part 2 Result: %d\n", part2Result)
}

func part1(input []string) int {
	joltage := 0
	for _, bank := range input {
		b := day03.Bank(bank)
		largestJoltage := b.LargestJoltage(2)
		joltage += int(largestJoltage)
	}

	return joltage
}

func part2(input []string) int {
	joltage := 0
	for _, bank := range input {
		b := day03.Bank(bank)
		largestJoltage := b.LargestJoltage(12)
		joltage += int(largestJoltage)
	}

	return joltage
}
