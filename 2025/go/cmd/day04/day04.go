package main

import (
	"fmt"

	"github.com/manning0218/adventOfCode/2025/go/day04"
	"github.com/manning0218/adventOfCode/2025/go/utils"
)

func main() {
	lines := utils.ReadInput(4)
	printDept := day04.NewGridPrintDept(lines)
	part1result := printDept.FindNumberPaperToMove(4)
	fmt.Println("Part 1:", part1result)

	part2Result := 0
	for pos := printDept.FindNumberPaperToMove(4); len(pos) > 0; pos = printDept.FindNumberPaperToMove(4) {
		part2Result += len(pos)
		printDept.RemoveRollLocations(pos)
	}

	fmt.Println("Part 2:", part2Result)
}
