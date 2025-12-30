package main

import (
	"fmt"

	"github.com/manning0218/adventOfCode/2025/go/day07"
	"github.com/manning0218/adventOfCode/2025/go/utils"
)

func main() {
	input := utils.ReadInput(7)
	diagram := day07.NewDiagram(input)

	start, err := diagram.FindStart()
	if err != nil {
		panic(fmt.Sprintf("failed to find start point: %v", err))
	}

	splitCount := diagram.ShootBeam(start)
	fmt.Println("Part 1 number of splits:", splitCount)

	// Part 2
	pathFinder := day07.NewBeamPathFinder(diagram)
	numberOfTimelines := pathFinder.CountPaths(start)
	fmt.Println("Part 2 number of unique paths:", numberOfTimelines)
}
