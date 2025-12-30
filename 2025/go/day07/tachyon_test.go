package day07

import (
	"testing"
)

func TestFindStart(t *testing.T) {
	lines := []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^.....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}

	diagram := NewDiagram(lines)

	start, err := diagram.FindStart()
	if err != nil {
		t.Fatalf("expected to find start point, got error: %v", err)
	}

	expected := Point{X: 0, Y: 7}
	if start != expected {
		t.Errorf("expected start point %v, got %v", expected, start)
	}
}

func TestShootBeam(t *testing.T) {
	lines := []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^.....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}

	diagram := NewDiagram(lines)

	start, err := diagram.FindStart()
	if err != nil {
		t.Fatalf("expected to find start point, got error: %v", err)
	}

	splitCount := diagram.ShootBeam(start)
	expectedSplits := 21
	if splitCount != expectedSplits {
		t.Errorf("expected %d splits, got %d", expectedSplits, splitCount)
	}
}

func TestFindPaths(t *testing.T) {
	lines := []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^.....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}

	diagram := NewDiagram(lines)
	finder := NewBeamPathFinder(diagram)

	start, err := diagram.FindStart()
	if err != nil {
		t.Fatalf("expected to find start point, got error: %v", err)
	}

	paths := finder.FindAllPaths(start)
	expectedPathCount := 40

	if len(paths) != expectedPathCount {
		t.Errorf("expected %d paths, got %d", expectedPathCount, len(paths))
	}
}

func TestCountPaths(t *testing.T) {
	lines := []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^.....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}

	diagram := NewDiagram(lines)
	finder := NewBeamPathFinder(diagram)

	start, err := diagram.FindStart()
	if err != nil {
		t.Fatalf("expected to find start point, got error: %v", err)
	}

	pathCount := finder.CountPaths(start)
	expectedCount := 40

	if pathCount != expectedCount {
		t.Errorf("CountPaths() = %d; want %d", pathCount, expectedCount)
	}

	t.Logf("Found %d unique paths (21 splits)", pathCount)
}

func TestCountPathsSimple(t *testing.T) {
	lines := []string{
		"..S..",
		".....",
		"..^..",
		".....",
	}

	diagram := NewDiagram(lines)
	finder := NewBeamPathFinder(diagram)
	start, _ := diagram.FindStart()

	pathCount := finder.CountPaths(start)
	expectedCount := 2 // Splits into 2 paths

	if pathCount != expectedCount {
		t.Errorf("CountPaths() = %d; want %d", pathCount, expectedCount)
	}
}
