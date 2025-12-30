package day07

import (
	"errors"
	"fmt"
)

var ErrStartNotFound = errors.New("start point 'S' not found in diagram")

type Diagram [][]rune

func NewDiagram(lines []string) Diagram {
	diagram := make(Diagram, len(lines))
	for i, line := range lines {
		diagram[i] = []rune(line)
	}
	return diagram
}

type Point struct{ X, Y int }
type Path []Point

type Direction int

const (
	Down Direction = iota
	Left
	Right
)

func (d Diagram) FindStart() (Point, error) {
	for x, row := range d {
		for y, data := range row {
			if data == 'S' {
				return Point{X: x, Y: y}, nil
			}
		}
	}
	return Point{}, ErrStartNotFound
}

func (d Diagram) ShootBeam(start Point) int {
	splitCount := 0
	visited := map[Point]bool{}
	visit := []Point{start}

	// Use depth-first search to explore the diagram
	for len(visit) > 0 {
		current := visit[len(visit)-1]
		visit = visit[:len(visit)-1]

		if visited[current] {
			continue
		}
		visited[current] = true

		// Check if the current position is within bounds
		if current.X < 0 || current.X >= len(d) || current.Y < 0 || current.Y >= len(d[0]) {
			continue
		}

		fmt.Println("Visiting:", current)

		cell := d[current.X][current.Y]
		if cell == '^' {
			fmt.Println("Split at:", current)
			splitCount++
			visit = append(visit, Point{X: current.X, Y: current.Y - 1})
			visit = append(visit, Point{X: current.X, Y: current.Y + 1})
			continue
		}

		visit = append(visit, Point{X: current.X + 1, Y: current.Y})
	}

	return splitCount
}

type BeamWork struct {
	Pos         Point
	Direction   Direction
	CurrentPath []Point
}

type BeamPathFinder struct {
	Diagram Diagram
	Dir     Direction
	Paths   []Path
}

func NewBeamPathFinder(diagram Diagram) *BeamPathFinder {
	return &BeamPathFinder{
		Diagram: diagram,
		Dir:     Down,
		Paths:   []Path{},
	}
}

func (d Diagram) IsInBounds(p Point) bool {
	return p.X >= 0 && p.X < len(d) && p.Y >= 0 && p.Y < len(d[0])
}

func (bpf *BeamPathFinder) FindAllPaths(start Point) []Path {
	// Pre-allocate based on grid size estimate
	estimatedPaths := (len(bpf.Diagram) * len(bpf.Diagram[0])) / 10
	if estimatedPaths < 32 {
		estimatedPaths = 32
	}
	bpf.Paths = make([]Path, 0, estimatedPaths)

	// Pre-allocate stack with reasonable capacity
	estimatedStackSize := len(bpf.Diagram) * 2
	if estimatedStackSize < 64 {
		estimatedStackSize = 64
	}
	stack := make([]BeamWork, 0, estimatedStackSize)

	// Initial beam with pre-allocated path
	initialPath := make([]Point, 1, len(bpf.Diagram)*2)
	initialPath[0] = start
	stack = append(stack, BeamWork{
		Pos:         start,
		Direction:   Down,
		CurrentPath: initialPath,
	})

	for len(stack) > 0 {
		// Pop from BACK of stack (DFS)
		beam := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nextPos := bpf.move(beam.Pos, beam.Direction)

		// Check if exited grid
		if !bpf.Diagram.IsInBounds(nextPos) {
			// Path complete - only allocate for final storage
			completePath := make(Path, len(beam.CurrentPath))
			copy(completePath, beam.CurrentPath)
			bpf.Paths = append(bpf.Paths, completePath)
			continue
		}

		// Reuse capacity from beam.CurrentPath when extending
		// This avoids a full copy + allocation
		newPath := append(beam.CurrentPath[:len(beam.CurrentPath):len(beam.CurrentPath)], nextPos)

		cell := bpf.Diagram[nextPos.X][nextPos.Y]

		switch cell {
		case '.', 'S':
			// Push to stack
			stack = append(stack, BeamWork{
				Pos:         nextPos,
				Direction:   beam.Direction,
				CurrentPath: newPath,
			})

		// Handle splits
		case '^':
			leftPos := bpf.move(nextPos, Left)
			rightPos := bpf.move(nextPos, Right)

			// Process both split directions
			bpf.handleSplit(leftPos, rightPos, newPath, &stack)
		}
	}

	return bpf.Paths
}

// handleSplit processes a beam split, reusing path capacity
func (bpf *BeamPathFinder) handleSplit(leftPos, rightPos Point, basePath []Point, stack *[]BeamWork) {
	// Left beam
	if bpf.Diagram.IsInBounds(leftPos) {
		// Reuse basePath capacity for left
		leftPath := append(basePath[:len(basePath):len(basePath)], leftPos)
		*stack = append(*stack, BeamWork{
			Pos:         leftPos,
			Direction:   Down,
			CurrentPath: leftPath,
		})
	} else {
		// Exits immediately
		exitPath := append(basePath[:len(basePath):len(basePath)], leftPos)
		bpf.Paths = append(bpf.Paths, exitPath)
	}

	// Right beam - need fresh copy since left may have reused capacity
	rightPath := make([]Point, len(basePath), len(basePath)+16)
	copy(rightPath, basePath)

	if bpf.Diagram.IsInBounds(rightPos) {
		rightPath = append(rightPath, rightPos)
		*stack = append(*stack, BeamWork{
			Pos:         rightPos,
			Direction:   Down,
			CurrentPath: rightPath,
		})
	} else {
		// Exits immediately
		rightPath = append(rightPath, rightPos)
		bpf.Paths = append(bpf.Paths, rightPath)
	}
}

func (bpf *BeamPathFinder) move(p Point, dir Direction) Point {
	switch dir {
	case Down:
		return Point{X: p.X + 1, Y: p.Y}
	case Left:
		return Point{X: p.X, Y: p.Y - 1}
	case Right:
		return Point{X: p.X, Y: p.Y + 1}
	}

	return p
}

// CountPaths counts total unique paths without storing them (much faster!)
func (bpf *BeamPathFinder) CountPaths(start Point) int {
	// Track active beams at each column (column -> beam count)
	currentBeams := make(map[int]int)
	currentBeams[start.Y] = 1

	totalPaths := 0

	// Process row by row from start
	for row := start.X; row < len(bpf.Diagram); row++ {
		nextBeams := make(map[int]int)

		for col, beamCount := range currentBeams {
			// Move down one row
			nextRow := row + 1

			// Check if beams exit the grid
			if nextRow >= len(bpf.Diagram) {
				totalPaths += beamCount
				continue
			}

			// Check what's in the next position
			nextCol := col

			// First check bounds
			if nextCol < 0 || nextCol >= len(bpf.Diagram[0]) {
				totalPaths += beamCount
				continue
			}

			cell := bpf.Diagram[nextRow][nextCol]

			switch cell {
			case '.', 'S':
				// Beams continue straight down
				nextBeams[nextCol] += beamCount

			case '^':
				// Split into left and right beams
				leftCol := nextCol - 1
				rightCol := nextCol + 1

				// Left beam
				if leftCol >= 0 && leftCol < len(bpf.Diagram[0]) {
					// Continue down from left position
					nextBeams[leftCol] += beamCount
				} else {
					// Exits immediately
					totalPaths += beamCount
				}

				// Right beam
				if rightCol >= 0 && rightCol < len(bpf.Diagram[0]) {
					// Continue down from right position
					nextBeams[rightCol] += beamCount
				} else {
					// Exits immediately
					totalPaths += beamCount
				}
			}
		}

		currentBeams = nextBeams

		// If no more beams, we're done
		if len(currentBeams) == 0 {
			break
		}
	}

	// Any remaining beams have exited
	for _, beamCount := range currentBeams {
		totalPaths += beamCount
	}

	return totalPaths
}
