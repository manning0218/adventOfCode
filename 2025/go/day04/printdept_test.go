package day04

import "testing"

func TestNumberOfNeighbors(t *testing.T) {
	lines := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	dept := NewGridPrintDept(lines)

	tests := []struct {
		row, col      int
		target        rune
		expectedCount int
	}{
		{0, 0, '@', 2},
		{1, 1, '@', 6},
		{4, 4, '@', 8},
		{9, 9, '@', 2},
		{5, 5, '.', 2},
	}

	for _, test := range tests {
		count := dept.NumberOfNeighbors(test.row, test.col, test.target)
		if count != test.expectedCount {
			t.Errorf("NumberOfNeighbors(%d, %d, '%c') = %d; want %d",
				test.row, test.col, test.target, count, test.expectedCount)
		}
	}
}

func TestFindNumberPaperToMove(t *testing.T) {
	lines := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	dept := NewGridPrintDept(lines)

	maxNeighbors := 4
	expectedCount := 13
	locations := dept.FindNumberPaperToMove(maxNeighbors)
	count := len(locations)
	if count != expectedCount {
		t.Errorf("FindNumberPaperToMove(%d) = %d; want %d",
			maxNeighbors, count, expectedCount)
	}
}

func TestFindNumberOfGridsThatCanBeRemoved(t *testing.T) {
	lines := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	dept := NewGridPrintDept(lines)

	maxNeighbors := 4
	expectedTotalRemoved := 43
	totalRemoved := 0

	for positions := dept.FindNumberPaperToMove(maxNeighbors); len(positions) > 0; positions = dept.FindNumberPaperToMove(maxNeighbors) {
		totalRemoved += len(positions)
		dept.RemoveRollLocations(positions)
	}

	if totalRemoved != expectedTotalRemoved {
		t.Errorf("Total grids removed = %d; want %d", totalRemoved, expectedTotalRemoved)
	}
}
