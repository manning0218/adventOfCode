package main

import (
	"fmt"

	"github.com/manning0218/adventOfCode/2025/go/day05"
	"github.com/manning0218/adventOfCode/2025/go/utils"
)

func main() {
	lines := utils.ReadInput(5)
	freshIngredients, availableIngredients, err := day05.NewIngredients(lines)
	if err != nil {
		fmt.Println("Error creating ingredients:", err)
		return
	}

	part1Result := availableIngredients.CountFreshIngredientsAvailable(freshIngredients)
	fmt.Println("Part 1:", part1Result)

	part2Result := freshIngredients.CountTotalFreshIngredients()
	fmt.Println("Part 2:", part2Result)
}
