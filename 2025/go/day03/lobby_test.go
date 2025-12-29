package day03

import "testing"

func TestLargestJoltage(t *testing.T) {
	tests := []struct {
		name        string
		bank        Bank
		cellsToKeep int
		expected    Joltage
	}{
		{
			name:        "Example 1",
			bank:        Bank("987654321111111"),
			cellsToKeep: 2,
			expected:    98,
		},
		{
			name:        "Example 2",
			bank:        Bank("811111111111119"),
			cellsToKeep: 2,
			expected:    89,
		},
		{
			name:        "Example 3",
			bank:        Bank("234234234234278"),
			cellsToKeep: 2,
			expected:    78,
		},
		{
			name:        "Example 4",
			bank:        Bank("818181911112111"),
			cellsToKeep: 2,
			expected:    92,
		},
		{
			name:        "Example 5",
			bank:        Bank("987654321111111"),
			cellsToKeep: 12,
			expected:    987654321111,
		},
		{
			name:        "Example 6",
			bank:        Bank("811111111111119"),
			cellsToKeep: 12,
			expected:    811111111119,
		},
		{
			name:        "Example 7",
			bank:        Bank("234234234234278"),
			cellsToKeep: 12,
			expected:    434234234278,
		},
		{
			name:        "Example 8",
			bank:        Bank("818181911112111"),
			cellsToKeep: 12,
			expected:    888911112111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.bank.LargestJoltage(tt.cellsToKeep)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
