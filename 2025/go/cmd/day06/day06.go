package main

import (
	"fmt"

	"github.com/manning0218/adventOfCode/2025/go/day06"
	"github.com/manning0218/adventOfCode/2025/go/utils"
)

func main() {
	lines := utils.ReadInput(6)
	columns, err := day06.ParseColumns(lines)
	if err != nil {
		fmt.Println("Error parsing columns:", err)
		return
	}

	part1Results := columns.ComputeResults()
	fmt.Println("Results: ", grandTotal(part1Results))

	cephalopodColumns, err := day06.ParseCephalopod(lines)
	if err != nil {
		fmt.Println("Error parsing cephalopod columns:", err)
		return
	}

	part2Results := cephalopodColumns.ComputeResults()
	fmt.Println("Cephalopod Results: ", grandTotal(part2Results))
}

func grandTotal(results []int64) int64 {
	var total int64
	for _, result := range results {
		total += result
	}
	return total
}
