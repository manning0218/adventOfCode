package day07

import "testing"

func BenchmarkFindAllPaths(b *testing.B) {
	lines := []string{
		"..S..",
		".....",
		"..^..",
		".....",
		".^.^.",
		".....",
		".....",
		"..^..",
		".....",
	}

	diagram := NewDiagram(lines)
	finder := NewBeamPathFinder(diagram)
	start, _ := diagram.FindStart()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		finder.FindAllPaths(start)
	}
}

func BenchmarkFindAllPathsLarge(b *testing.B) {
	// Create a larger grid with more splitters
	lines := []string{
		"..S..",
		".....",
		"..^..",
		".....",
		".^.^.",
		".....",
		"..^..",
		".....",
		".^.^.",
		".....",
		"..^..",
		".....",
		".^.^.",
		".....",
		".....",
	}

	diagram := NewDiagram(lines)
	finder := NewBeamPathFinder(diagram)
	start, _ := diagram.FindStart()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		finder.FindAllPaths(start)
	}
}

func BenchmarkFindAllPathsNoSplits(b *testing.B) {
	lines := []string{
		"..S..",
		".....",
		".....",
		".....",
		".....",
	}

	diagram := NewDiagram(lines)
	finder := NewBeamPathFinder(diagram)
	start, _ := diagram.FindStart()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		finder.FindAllPaths(start)
	}
}

func BenchmarkCountPaths(b *testing.B) {
	lines := []string{
		"..S..",
		".....",
		"..^..",
		".....",
		".^.^.",
		".....",
		".....",
		"..^..",
		".....",
	}

	diagram := NewDiagram(lines)
	finder := NewBeamPathFinder(diagram)
	start, _ := diagram.FindStart()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		finder.CountPaths(start)
	}
}

func BenchmarkCountPathsLarge(b *testing.B) {
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
	start, _ := diagram.FindStart()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		finder.CountPaths(start)
	}
}
