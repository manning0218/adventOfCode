package day05

import "testing"

func TestCountFreshIngredientsAvailable(t *testing.T) {
	lines := []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	}
	freshIngredients, availableIngredients, _ := NewIngredients(lines)
	expectedCount := 3 // Ingredients 5, 11, and 17 are fresh

	count := availableIngredients.CountFreshIngredientsAvailable(freshIngredients)
	if count != expectedCount {
		t.Errorf("CountFreshIngredientsAvailable() = %d; want %d", count, expectedCount)
	}
}

func TestCountTotalFreshIngredients(t *testing.T) {
	lines := []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"", // Empty line separator required
	}
	freshIngredients, _, err := NewIngredients(lines)
	if err != nil {
		t.Fatalf("NewIngredients() unexpected error: %v", err)
	}
	expectedTotal := int64(14) // Unique fresh ingredients: 3,4,5,10,11,12,13,14,16,17,18,19,20

	total := freshIngredients.CountTotalFreshIngredients()
	if total != expectedTotal {
		t.Errorf("CountTotalFreshIngredients() = %d; want %d", total, expectedTotal)
	}
}
