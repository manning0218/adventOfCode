package main

import (
	"fmt"
	"sort"

	"github.com/manning0218/adventOfCode/2025/go/day08"
	"github.com/manning0218/adventOfCode/2025/go/utils"
)

func main() {
	input := utils.ReadInput(8)

	junctionBoxes := day08.NewJunctionBoxes(input)

	// Part 1: Find the 1000 shortest connections
	circuits := junctionBoxes.BuildCircuits(1000)
	fmt.Println("Part 1: 1000 Shortest Connections:")
	circuitSizes := make([]int, len(circuits))
	for i, circuit := range circuits {
		fmt.Printf("Circuit %d size: %d\n", i+1, len(circuit))
		circuitSizes[i] = len(circuit)
	}
	sort.IntSlice(circuitSizes).Sort()
	part1Result := 1
	for _, size := range circuitSizes[len(circuitSizes)-3:] {
		fmt.Println("Multiplying by size:", size)
		part1Result *= size
	}
	fmt.Println("Part 1 Result (product of 4 smallest circuits):", part1Result)

	// Part 2: Find the last connection
	lastConnection, err := junctionBoxes.FindLastConnection(1000)
	if err != nil {
		panic(fmt.Sprintf("failed to find last connection: %v", err))
	}
	fmt.Printf("Part 2: Last Connection between boxes at (%d,%d,%d) and (%d,%d,%d) with distance %.2f\n",
		lastConnection.Box1.X, lastConnection.Box1.Y, lastConnection.Box1.Z,
		lastConnection.Box2.X, lastConnection.Box2.Y, lastConnection.Box2.Z,
		lastConnection.Distance)

	fmt.Println("Part 2 Result: ", lastConnection.Box1.X*lastConnection.Box2.X)
}
