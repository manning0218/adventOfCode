package main

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/manning0218/adventOfCode/2025/go/day02"
	"github.com/manning0218/adventOfCode/2025/go/utils"
)

func main() {
	input := utils.ReadInput(2)
	records := parseRanges(input[0])
	testInput := [][2]day02.ProductID{
		{
			day02.ProductID("11"),
			day02.ProductID("22"),
		},
		{
			day02.ProductID("95"),
			day02.ProductID("115"),
		},
		{
			day02.ProductID("998"),
			day02.ProductID("1012"),
		},
		{
			day02.ProductID("1188511880"),
			day02.ProductID("1188511890"),
		},
		{
			day02.ProductID("222220"),
			day02.ProductID("222224"),
		},
		{
			day02.ProductID("1698522"),
			day02.ProductID("1698528"),
		},
		{
			day02.ProductID("446443"),
			day02.ProductID("446449"),
		},
		{
			day02.ProductID("38593856"),
			day02.ProductID("38593862"),
		},
		{
			day02.ProductID("565653"),
			day02.ProductID("565659"),
		},
		{
			day02.ProductID("824824821"),
			day02.ProductID("824824827"),
		},
		{
			day02.ProductID("2121212118"),
			day02.ProductID("2121212124"),
		},
	}
	//part1TestResult := part1(testInput)
	//fmt.Printf("Part 1 Test Result: %d\n", part1TestResult)

	//part1Result := part1(records)
	//fmt.Printf("Part 1 Result: %d\n", part1Result)

	part2TestResult := part2(testInput)
	fmt.Printf("Part 2 Test Result: %d\n", part2TestResult)

	part2Result := part2(records)
	fmt.Printf("Part 2 Result: %d\n", part2Result)
}

func part1(ranges [][2]day02.ProductID) int {
	invalidIDSum := 0
	for _, r := range ranges {
		//fmt.Printf("Checking range: %s - %s\n", r[0], r[1])
		for id := r[0].Value(); id <= r[1].Value(); id++ {
			pid := day02.ProductID(fmt.Sprintf("%d", id))
			if pid.IsInvalid() {
				fmt.Println("Invalid ID:", pid)
				invalidIDSum += pid.Value()
			}
		}
	}
	return invalidIDSum
}

func part2(ranges [][2]day02.ProductID) int {
	invalidIDSum := 0
	for _, r := range ranges {
		fmt.Printf("Checking range: %s - %s\n", r[0], r[1])
		for id := r[0].Value(); id <= r[1].Value(); id++ {
			pid := day02.ProductID(fmt.Sprintf("%d", id))
			if pid.IsInvalid2(id) {
				fmt.Println("Invalid ID (part 2):", pid)
				invalidIDSum += pid.Value()
			}
		}
	}
	return invalidIDSum
}

func parseRanges(line string) [][2]day02.ProductID {
	r := csv.NewReader(strings.NewReader(line))
	fields, err := r.Read()
	if err != nil {
		panic(fmt.Sprintf("failed to parse line: %v", err))
	}

	var ranges [][2]day02.ProductID
	for _, field := range fields {
		ids := strings.Split(field, "-")
		if len(ids) != 2 {
			panic(fmt.Sprintf("invalid range: %s", field))
		}
		ranges = append(ranges, [2]day02.ProductID{day02.ProductID(ids[0]), day02.ProductID(ids[1])})
	}

	return ranges
}
