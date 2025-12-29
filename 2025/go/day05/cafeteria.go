package day05

import (
	"fmt"
	"strconv"
)

type FreshIngredients struct {
	tree *AVLIntervalTree
}

type AvailableIngredients []int64

func NewIngredients(lines []string) (*FreshIngredients, AvailableIngredients, error) {
	tree := NewAVLIntervalTree()
	availableStart := 0

	// Parse fresh ingredient ranges and build AVL tree
	for i, line := range lines {
		if line == "" {
			availableStart = i + 1
			break
		}

		fmt.Println("Processing line for fresh ingredients:", line)

		// Parse the interval using the AVL tree's ParseInterval function
		interval, err := ParseInterval(line)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse interval '%s': %w", line, err)
		}
		tree.Insert(interval)
	}

	freshIngredients := &FreshIngredients{tree: tree}

	// Parse available ingredients
	availableIngredients := make(AvailableIngredients, 0)
	for _, line := range lines[availableStart:] {
		if line == "" {
			continue
		}
		fmt.Println("Processing line for available ingredients:", line)
		ingredient, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse ingredient '%s': %w", line, err)
		}
		availableIngredients = append(availableIngredients, ingredient)
	}

	return freshIngredients, availableIngredients, nil
}

func (ai AvailableIngredients) CountFreshIngredientsAvailable(fi *FreshIngredients) int {
	count := 0
	for _, available := range ai {
		if fi.tree.Contains(available) {
			count++
		}
	}
	return count
}

func (fi *FreshIngredients) Contains(ingredient int64) bool {
	return fi.tree.Contains(ingredient)
}

func (fi *FreshIngredients) CountTotalFreshIngredients() int64 {
	return fi.tree.CountUniqueElements()
}
