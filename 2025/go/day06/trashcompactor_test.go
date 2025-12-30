package day06

import (
	"testing"
)

func TestComputeResults(t *testing.T) {
	input := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}

	cols, err := ParseColumns(input)
	if err != nil {
		t.Fatalf("ParseColumns() unexpected error: %v", err)
	}

	expectedResults := []int64{33210, 490, 4243455, 401}
	results := cols.ComputeResults()

	for i, result := range results {
		if result != expectedResults[i] {
			t.Errorf("ComputeResults()[%d] = %d; want %d", i, result, expectedResults[i])
		}
	}
}

func TestParseCephalopod(t *testing.T) {
	input := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}

	cols, err := ParseCephalopod(input)
	if err != nil {
		t.Fatalf("ParseCephalopod() unexpected error: %v", err)
	}

	if len(cols) != 4 {
		t.Fatalf("ParseCephalopod() returned %d columns; want 4", len(cols))
	}

	expectedColumns := []Column{
		{Data: []int64{4, 431, 623}, Operation: '+'},
		{Data: []int64{175, 581, 32}, Operation: '*'},
		{Data: []int64{8, 248, 369}, Operation: '+'},
		{Data: []int64{356, 24, 1}, Operation: '*'},
	}

	for i, col := range cols {
		if col.Operation != expectedColumns[i].Operation {
			t.Errorf("Column %d operation = '%c'; want '%c'", i, col.Operation, expectedColumns[i].Operation)
		}
		if len(col.Data) != len(expectedColumns[i].Data) {
			t.Errorf("Column %d data length = %d; want %d", i, len(col.Data), len(expectedColumns[i].Data))
			continue
		}
		for j, val := range col.Data {
			if val != expectedColumns[i].Data[j] {
				t.Errorf("Column %d data[%d] = %d; want %d", i, j, val, expectedColumns[i].Data[j])
			}
		}
	}
}
