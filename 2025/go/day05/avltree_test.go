package day05

import "testing"

func TestParseInterval(t *testing.T) {
	tests := []struct {
		input       string
		expected    Interval
		expectError bool
	}{
		{"10-20", Interval{10, 20}, false},
		{"291687894568177-292172488078380", Interval{291687894568177, 292172488078380}, false},
		{"0-0", Interval{0, 0}, false},
		{"100-100", Interval{100, 100}, false},
		{"invalid", Interval{}, true},
		{"10-20-30", Interval{}, true},
		{"abc-def", Interval{}, true},
		{"20-10", Interval{}, true}, // start > end
	}

	for _, test := range tests {
		result, err := ParseInterval(test.input)
		if test.expectError {
			if err == nil {
				t.Errorf("ParseInterval(%s) expected error but got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseInterval(%s) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ParseInterval(%s) = %+v; want %+v", test.input, result, test.expected)
			}
		}
	}
}

func TestAVLIntervalTreeInsertAndContains(t *testing.T) {
	tree := NewAVLIntervalTree()

	// Insert intervals
	intervals := []Interval{
		{10, 20},
		{50, 60},
		{5, 8},
		{15, 25},
		{100, 200},
	}

	for _, interval := range intervals {
		tree.Insert(interval)
	}

	// Test values that should be contained
	containedTests := []int64{
		10, 15, 20, // in [10, 20]
		5, 6, 7, 8, // in [5, 8]
		50, 55, 60, // in [50, 60]
		15, 20, 25, // in [15, 25] (overlaps with [10, 20])
		100, 150, 200, // in [100, 200]
	}

	for _, value := range containedTests {
		if !tree.Contains(value) {
			t.Errorf("Contains(%d) = false; want true", value)
		}
	}

	// Test values that should NOT be contained
	notContainedTests := []int64{
		0, 4, 9, 26, 49, 61, 99, 201, 1000,
	}

	for _, value := range notContainedTests {
		if tree.Contains(value) {
			t.Errorf("Contains(%d) = true; want false", value)
		}
	}
}

func TestAVLIntervalTreeLargeRanges(t *testing.T) {
	tree := NewAVLIntervalTree()

	// Insert large ranges like the fresh ingredients example
	tree.Insert(Interval{291687894568177, 292172488078380})
	tree.Insert(Interval{100000000000000, 200000000000000})
	tree.Insert(Interval{500000000000000, 600000000000000})

	// Test values within ranges
	if !tree.Contains(291687894568177) {
		t.Error("Contains(291687894568177) = false; want true")
	}
	if !tree.Contains(292172488078380) {
		t.Error("Contains(292172488078380) = false; want true")
	}
	if !tree.Contains(291900000000000) {
		t.Error("Contains(291900000000000) = false; want true")
	}

	// Test values outside ranges
	if tree.Contains(291687894568176) {
		t.Error("Contains(291687894568176) = true; want false")
	}
	if tree.Contains(292172488078381) {
		t.Error("Contains(292172488078381) = true; want false")
	}
	if tree.Contains(300000000000000) {
		t.Error("Contains(300000000000000) = true; want false")
	}
}

func TestAVLIntervalTreeOverlappingIntervals(t *testing.T) {
	tree := NewAVLIntervalTree()

	// Insert overlapping intervals
	tree.Insert(Interval{10, 30})
	tree.Insert(Interval{20, 40})
	tree.Insert(Interval{35, 50})

	// All values from 10 to 50 should be contained
	for i := int64(10); i <= 50; i++ {
		if !tree.Contains(i) {
			t.Errorf("Contains(%d) = false; want true (overlapping intervals)", i)
		}
	}

	// Values outside should not be contained
	if tree.Contains(9) || tree.Contains(51) {
		t.Error("Contains returned true for values outside all intervals")
	}
}

func TestAVLIntervalTreeBalance(t *testing.T) {
	tree := NewAVLIntervalTree()

	// Insert intervals in sorted order (worst case for unbalanced tree)
	for i := int64(0); i < 100; i += 10 {
		tree.Insert(Interval{i, i + 5})
	}

	// Tree should remain balanced
	// Height should be O(log n), for 10 intervals: ~4
	if tree.Root.Height > 6 {
		t.Errorf("Tree is unbalanced, height = %d; expected <= 6", tree.Root.Height)
	}

	// All inserted values should still be found
	for i := int64(0); i < 100; i += 10 {
		if !tree.Contains(i) {
			t.Errorf("Contains(%d) = false; want true after balancing", i)
		}
		if !tree.Contains(i + 5) {
			t.Errorf("Contains(%d) = false; want true after balancing", i+5)
		}
	}
}

func TestNewAVLIntervalTreeFromStrings(t *testing.T) {
	intervals := []string{
		"10-20",
		"50-60",
		"291687894568177-292172488078380",
	}

	tree, err := NewAVLIntervalTreeFromStrings(intervals)
	if err != nil {
		t.Fatalf("NewAVLIntervalTreeFromStrings() unexpected error: %v", err)
	}

	if tree.Size != 3 {
		t.Errorf("tree.Size = %d; want 3", tree.Size)
	}

	// Test contains
	if !tree.Contains(15) {
		t.Error("Contains(15) = false; want true")
	}
	if !tree.Contains(55) {
		t.Error("Contains(55) = false; want true")
	}
	if !tree.Contains(291900000000000) {
		t.Error("Contains(291900000000000) = false; want true")
	}
}

func TestNewAVLIntervalTreeFromStringsInvalid(t *testing.T) {
	intervals := []string{
		"10-20",
		"invalid",
		"50-60",
	}

	_, err := NewAVLIntervalTreeFromStrings(intervals)
	if err == nil {
		t.Error("NewAVLIntervalTreeFromStrings() expected error for invalid input")
	}
}

func TestAVLIntervalTreeEmpty(t *testing.T) {
	tree := NewAVLIntervalTree()

	if tree.Contains(0) {
		t.Error("Empty tree Contains(0) = true; want false")
	}
	if tree.Size != 0 {
		t.Errorf("Empty tree Size = %d; want 0", tree.Size)
	}
}

func TestAVLIntervalTreeSinglePoint(t *testing.T) {
	tree := NewAVLIntervalTree()
	tree.Insert(Interval{42, 42})

	if !tree.Contains(42) {
		t.Error("Contains(42) = false; want true for single-point interval")
	}
	if tree.Contains(41) || tree.Contains(43) {
		t.Error("Contains returned true for values outside single-point interval")
	}
}

// Benchmark for large tree
func BenchmarkAVLIntervalTreeInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := NewAVLIntervalTree()
		for j := int64(0); j < 1000; j++ {
			tree.Insert(Interval{j * 100, j*100 + 50})
		}
	}
}

func BenchmarkAVLIntervalTreeContains(b *testing.B) {
	tree := NewAVLIntervalTree()
	for j := int64(0); j < 1000; j++ {
		tree.Insert(Interval{j * 100, j*100 + 50})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Contains(50025)
	}
}

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name     string
		input    []Interval
		expected []Interval
	}{
		{
			name:     "No intervals",
			input:    []Interval{},
			expected: []Interval{},
		},
		{
			name:     "Single interval",
			input:    []Interval{{10, 20}},
			expected: []Interval{{10, 20}},
		},
		{
			name:     "Non-overlapping intervals",
			input:    []Interval{{10, 20}, {30, 40}, {50, 60}},
			expected: []Interval{{10, 20}, {30, 40}, {50, 60}},
		},
		{
			name:     "Overlapping intervals",
			input:    []Interval{{10, 20}, {15, 25}, {22, 30}},
			expected: []Interval{{10, 30}},
		},
		{
			name:     "Adjacent intervals",
			input:    []Interval{{10, 20}, {21, 30}},
			expected: []Interval{{10, 30}},
		},
		{
			name:     "Mix of overlapping and non-overlapping",
			input:    []Interval{{3, 5}, {10, 14}, {12, 18}, {16, 20}, {30, 40}},
			expected: []Interval{{3, 5}, {10, 20}, {30, 40}},
		},
		{
			name:     "Fully contained interval",
			input:    []Interval{{10, 30}, {15, 20}},
			expected: []Interval{{10, 30}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MergeIntervals(test.input)
			if len(result) != len(test.expected) {
				t.Errorf("MergeIntervals() returned %d intervals; want %d", len(result), len(test.expected))
				return
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("MergeIntervals()[%d] = %+v; want %+v", i, result[i], test.expected[i])
				}
			}
		})
	}
}

func TestCountTotalElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []Interval
		expected int64
	}{
		{
			name:     "Empty",
			input:    []Interval{},
			expected: 0,
		},
		{
			name:     "Single interval",
			input:    []Interval{{10, 20}},
			expected: 11, // 10,11,12,...,20 = 11 elements
		},
		{
			name:     "Multiple intervals",
			input:    []Interval{{10, 20}, {30, 40}},
			expected: 22, // 11 + 11
		},
		{
			name:     "Single point interval",
			input:    []Interval{{42, 42}},
			expected: 1,
		},
		{
			name:     "Large range",
			input:    []Interval{{1, 1000000}},
			expected: 1000000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountTotalElements(test.input)
			if result != test.expected {
				t.Errorf("CountTotalElements() = %d; want %d", result, test.expected)
			}
		})
	}
}

func TestCountUniqueElements(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		expected  int64
	}{
		{
			name:      "Empty tree",
			intervals: []Interval{},
			expected:  0,
		},
		{
			name:      "Non-overlapping intervals",
			intervals: []Interval{{10, 20}, {30, 40}, {50, 60}},
			expected:  33, // 11 + 11 + 11
		},
		{
			name:      "Overlapping intervals",
			intervals: []Interval{{10, 20}, {15, 25}},
			expected:  16, // Merged to [10, 25] = 16 elements
		},
		{
			name:      "Adjacent intervals",
			intervals: []Interval{{10, 20}, {21, 30}},
			expected:  21, // Merged to [10, 30] = 21 elements
		},
		{
			name:      "Test case intervals",
			intervals: []Interval{{3, 5}, {10, 14}, {16, 20}, {12, 18}},
			expected:  14, // [3-5]=3, [10-20]=11 (after merging overlapping intervals), total=14
		},
		{
			name:      "Fully contained",
			intervals: []Interval{{10, 30}, {15, 20}},
			expected:  21, // Merged to [10, 30]
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := NewAVLIntervalTree()
			for _, interval := range test.intervals {
				tree.Insert(interval)
			}
			result := tree.CountUniqueElements()
			if result != test.expected {
				t.Errorf("CountUniqueElements() = %d; want %d", result, test.expected)
			}
		})
	}
}

func TestGetAllIntervals(t *testing.T) {
	tree := NewAVLIntervalTree()

	// Insert in random order
	tree.Insert(Interval{50, 60})
	tree.Insert(Interval{10, 20})
	tree.Insert(Interval{30, 40})
	tree.Insert(Interval{5, 8})

	intervals := tree.GetAllIntervals()

	// Should be sorted by Start value
	expected := []Interval{{5, 8}, {10, 20}, {30, 40}, {50, 60}}

	if len(intervals) != len(expected) {
		t.Errorf("GetAllIntervals() returned %d intervals; want %d", len(intervals), len(expected))
		return
	}

	for i := range intervals {
		if intervals[i] != expected[i] {
			t.Errorf("GetAllIntervals()[%d] = %+v; want %+v", i, intervals[i], expected[i])
		}
	}
}
