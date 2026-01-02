package day08

import (
	"fmt"
)

type JunctionBoxLocation struct {
	X int
	Y int
	Z int
}

type JunctionBoxes []JunctionBoxLocation

func NewJunctionBoxes(input []string) JunctionBoxes {
	boxes := make(JunctionBoxes, 0, len(input))
	for _, line := range input {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic(fmt.Sprintf("failed to parse junction box location from line %q: %v", line, err))
		}
		boxes = append(boxes, JunctionBoxLocation{X: x, Y: y, Z: z})
	}
	return boxes
}

type Circuit []JunctionBoxLocation

// Connection represents a connection between two junction boxes
type Connection struct {
	Box1     JunctionBoxLocation
	Box2     JunctionBoxLocation
	Distance float64
}

// DistanceTo calculates 3D Euclidean distance squared between two junction boxes
func (j JunctionBoxLocation) DistanceTo(other JunctionBoxLocation) float64 {
	dx := float64(j.X - other.X)
	dy := float64(j.Y - other.Y)
	dz := float64(j.Z - other.Z)
	return dx*dx + dy*dy + dz*dz // Return squared distance to avoid sqrt (faster)
}

// FindNShortestConnections finds the n shortest distances between junction boxes
func (jb JunctionBoxes) FindNShortestConnections(n int) []Connection {
	connections := []Connection{}

	// Calculate all pairwise distances
	for i := 0; i < len(jb); i++ {
		for j := i + 1; j < len(jb); j++ {
			dist := jb[i].DistanceTo(jb[j])
			connections = append(connections, Connection{
				Box1:     jb[i],
				Box2:     jb[j],
				Distance: dist,
			})
		}
	}

	// Sort by distance (selection sort for simplicity, can use sort.Slice)
	for i := 0; i < len(connections)-1 && i < n; i++ {
		minIdx := i
		for j := i + 1; j < len(connections); j++ {
			if connections[j].Distance < connections[minIdx].Distance {
				minIdx = j
			}
		}
		connections[i], connections[minIdx] = connections[minIdx], connections[i]
	}

	// Return the n shortest
	if n > len(connections) {
		n = len(connections)
	}
	return connections[:n]
}

// BuildCircuits builds circuits from the n shortest connections using Union-Find
func (jb JunctionBoxes) BuildCircuits(n int) []Circuit {
	connections := jb.FindNShortestConnections(n)

	// Union-Find data structure
	parent := make(map[JunctionBoxLocation]JunctionBoxLocation)

	// Initialize each box as its own parent
	for _, box := range jb {
		parent[box] = box
	}

	// Find with path compression
	var find func(JunctionBoxLocation) JunctionBoxLocation
	find = func(box JunctionBoxLocation) JunctionBoxLocation {
		if parent[box] != box {
			parent[box] = find(parent[box]) // Path compression
		}
		return parent[box]
	}

	// Union two sets
	union := func(box1, box2 JunctionBoxLocation) {
		root1 := find(box1)
		root2 := find(box2)
		if root1 != root2 {
			parent[root2] = root1
		}
	}

	// Connect boxes based on shortest connections
	for _, conn := range connections {
		union(conn.Box1, conn.Box2)
	}

	// Group boxes by their root (circuit)
	circuits := make(map[JunctionBoxLocation][]JunctionBoxLocation)
	for _, box := range jb {
		root := find(box)
		circuits[root] = append(circuits[root], box)
	}

	// Convert to slice of circuits
	result := []Circuit{}
	for _, circuit := range circuits {
		result = append(result, circuit)
	}

	return result
}

// FindLastConnection finds the connection that would merge the last two separate circuits
// Optimized: uses existing circuits from BuildCircuits(n) to avoid recalculating
func (jb JunctionBoxes) FindLastConnection(initialN int) (*Connection, error) {
	// Build initial circuits using the first n shortest connections
	circuits := jb.BuildCircuits(initialN)

	// If already down to 1 circuit, we're done
	if len(circuits) <= 1 {
		return nil, fmt.Errorf("already fully connected with %d circuits", len(circuits))
	}

	// Create a map from box to circuit index
	boxToCircuit := make(map[JunctionBoxLocation]int)
	for circuitIdx, circuit := range circuits {
		for _, box := range circuit {
			boxToCircuit[box] = circuitIdx
		}
	}

	// Find shortest connections BETWEEN circuits (not within)
	type CircuitConnection struct {
		Connection
		Circuit1 int
		Circuit2 int
	}

	interCircuitConnections := []CircuitConnection{}

	// Only check connections between different circuits
	for i := 0; i < len(circuits); i++ {
		for j := i + 1; j < len(circuits); j++ {
			// Find shortest connection between circuit i and circuit j
			var shortestConn *Connection
			shortestDist := float64(1e18)

			for _, box1 := range circuits[i] {
				for _, box2 := range circuits[j] {
					dist := box1.DistanceTo(box2)
					if dist < shortestDist {
						shortestDist = dist
						shortestConn = &Connection{
							Box1:     box1,
							Box2:     box2,
							Distance: dist,
						}
					}
				}
			}

			if shortestConn != nil {
				interCircuitConnections = append(interCircuitConnections, CircuitConnection{
					Connection: *shortestConn,
					Circuit1:   i,
					Circuit2:   j,
				})
			}
		}
	}

	// Sort inter-circuit connections by distance
	for i := 0; i < len(interCircuitConnections)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(interCircuitConnections); j++ {
			if interCircuitConnections[j].Distance < interCircuitConnections[minIdx].Distance {
				minIdx = j
			}
		}
		interCircuitConnections[i], interCircuitConnections[minIdx] = interCircuitConnections[minIdx], interCircuitConnections[i]
	}

	// Union-Find on circuits (not boxes)
	circuitParent := make([]int, len(circuits))
	for i := range circuitParent {
		circuitParent[i] = i
	}

	// Declare before defining to allow recursion
	var findCircuit func(int) int
	findCircuit = func(c int) int {
		if circuitParent[c] != c {
			circuitParent[c] = findCircuit(circuitParent[c])
		}
		return circuitParent[c]
	}

	unionCircuits := func(c1, c2 int) {
		root1 := findCircuit(c1)
		root2 := findCircuit(c2)
		if root1 != root2 {
			circuitParent[root2] = root1
		}
	}

	countCircuitComponents := func() int {
		roots := make(map[int]bool)
		for i := range circuits {
			roots[findCircuit(i)] = true
		}
		return len(roots)
	}

	// Merge circuits until we have exactly 2 left
	for _, conn := range interCircuitConnections {
		componentsCount := countCircuitComponents()

		// If down to 2 components, next connection is the answer
		if componentsCount == 2 {
			root1 := findCircuit(conn.Circuit1)
			root2 := findCircuit(conn.Circuit2)

			if root1 != root2 {
				return &conn.Connection, nil
			}
		}

		// Merge these circuits
		unionCircuits(conn.Circuit1, conn.Circuit2)
	}

	return nil, fmt.Errorf("could not find last connection")
}
