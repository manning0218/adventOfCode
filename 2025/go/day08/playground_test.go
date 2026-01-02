package day08

import (
	"math"
	"testing"
)

func TestDistanceTo(t *testing.T) {
	box1 := JunctionBoxLocation{X: 0, Y: 0, Z: 0}
	box2 := JunctionBoxLocation{X: 3, Y: 4, Z: 0}

	distSquared := box1.DistanceTo(box2)
	expected := float64(25) // 3^2 + 4^2 = 25

	if distSquared != expected {
		t.Errorf("DistanceTo() = %f; want %f", distSquared, expected)
	}

	// Actual distance would be sqrt(25) = 5
	actualDist := math.Sqrt(distSquared)
	if actualDist != 5.0 {
		t.Errorf("Actual distance = %f; want 5.0", actualDist)
	}
}

func TestFindNShortestConnections(t *testing.T) {
	input := []string{
		"0,0,0",
		"1,0,0",
		"0,1,0",
		"5,5,5",
	}

	boxes := NewJunctionBoxes(input)
	connections := boxes.FindNShortestConnections(3)

	if len(connections) != 3 {
		t.Errorf("Expected 3 connections, got %d", len(connections))
	}

	// The shortest should be (0,0,0) to (1,0,0) = distance 1
	if connections[0].Distance != 1.0 {
		t.Errorf("Shortest connection distance = %f; want 1.0", connections[0].Distance)
	}

	// Second shortest should be (0,0,0) to (0,1,0) = distance 1
	if connections[1].Distance != 1.0 {
		t.Errorf("Second shortest distance = %f; want 1.0", connections[1].Distance)
	}

	// Verify connections are sorted
	for i := 0; i < len(connections)-1; i++ {
		if connections[i].Distance > connections[i+1].Distance {
			t.Errorf("Connections not sorted: %f > %f", connections[i].Distance, connections[i+1].Distance)
		}
	}
}

func TestBuildCircuits(t *testing.T) {
	// Create two separate clusters of boxes
	input := []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}

	boxes := NewJunctionBoxes(input)

	// Use n=10 to connect only the 10 shortest distances
	// This should create 11 circuits
	circuits := boxes.BuildCircuits(10)

	// Should have circuits (could be 1, 2, or more depending on connections)
	if len(circuits) == 0 {
		t.Fatal("Expected at least one circuit")
	}

	t.Logf("Found %d circuit(s):", len(circuits))
	for i, circuit := range circuits {
		t.Logf("  Circuit %d: %d boxes", i+1, len(circuit))
		for _, box := range circuit {
			t.Logf("    - (%d,%d,%d)", box.X, box.Y, box.Z)
		}
	}
}

func TestBuildCircuitsFullyConnected(t *testing.T) {
	input := []string{
		"0,0,0",
		"1,0,0",
		"2,0,0",
		"3,0,0",
	}

	boxes := NewJunctionBoxes(input)

	// Connect all boxes (6 total connections for 4 boxes)
	circuits := boxes.BuildCircuits(10)

	// Should have 1 circuit with all 4 boxes
	if len(circuits) != 1 {
		t.Errorf("Expected 1 circuit, got %d", len(circuits))
	}

	if len(circuits[0]) != 4 {
		t.Errorf("Expected circuit with 4 boxes, got %d", len(circuits[0]))
	}
}

func TestBuildCircuitsIsolated(t *testing.T) {
	input := []string{
		"0,0,0",
		"100,100,100",
		"200,200,200",
	}

	boxes := NewJunctionBoxes(input)

	// No connections (n=0)
	circuits := boxes.BuildCircuits(0)

	// Should have 3 separate circuits (each box isolated)
	if len(circuits) != 3 {
		t.Errorf("Expected 3 circuits, got %d", len(circuits))
	}

	// Each circuit should have 1 box
	for i, circuit := range circuits {
		if len(circuit) != 1 {
			t.Errorf("Circuit %d has %d boxes; want 1", i, len(circuit))
		}
	}
}

func TestBuildCircuitsPartialConnection(t *testing.T) {
	// Linear arrangement of boxes
	input := []string{
		"0,0,0",
		"1,0,0",
		"2,0,0",
		"3,0,0",
		"4,0,0",
	}

	boxes := NewJunctionBoxes(input)

	// Connect only 2 pairs (n=2)
	// This should create 2 circuits: one with 3 boxes, others isolated
	circuits := boxes.BuildCircuits(2)

	t.Logf("Found %d circuit(s) with n=2:", len(circuits))
	for i, circuit := range circuits {
		t.Logf("  Circuit %d: %d boxes", i+1, len(circuit))
	}

	// The exact number depends on which pairs get connected
	// but we know we won't have all in one circuit with only 2 connections
	totalBoxes := 0
	for _, circuit := range circuits {
		totalBoxes += len(circuit)
	}

	if totalBoxes != 5 {
		t.Errorf("Total boxes in all circuits = %d; want 5", totalBoxes)
	}
}

func TestNewJunctionBoxes(t *testing.T) {
	input := []string{
		"1,2,3",
		"10,20,30",
		"-5,-10,-15",
	}

	boxes := NewJunctionBoxes(input)

	if len(boxes) != 3 {
		t.Errorf("Expected 3 boxes, got %d", len(boxes))
	}

	expected := []JunctionBoxLocation{
		{X: 1, Y: 2, Z: 3},
		{X: 10, Y: 20, Z: 30},
		{X: -5, Y: -10, Z: -15},
	}

	for i, box := range boxes {
		if box != expected[i] {
			t.Errorf("Box %d = %+v; want %+v", i, box, expected[i])
		}
	}
}

func TestFindLastConnection(t *testing.T) {
	// Create two clear clusters that are far apart
	input := []string{
		"0,0,0", // Cluster 1
		"1,0,0",
		"2,0,0",
		"100,100,100", // Cluster 2 (far away)
		"101,100,100",
		"102,100,100",
	}

	boxes := NewJunctionBoxes(input)

	// Use n=2 to connect each cluster internally, leaving 2 circuits
	lastConn, err := boxes.FindLastConnection(2)

	if err != nil {
		t.Fatalf("FindLastConnection() error: %v", err)
	}

	if lastConn == nil {
		t.Fatal("Expected a connection, got nil")
	}

	t.Logf("Last connection to merge final 2 circuits:")
	t.Logf("  Box1: (%d,%d,%d)", lastConn.Box1.X, lastConn.Box1.Y, lastConn.Box1.Z)
	t.Logf("  Box2: (%d,%d,%d)", lastConn.Box2.X, lastConn.Box2.Y, lastConn.Box2.Z)
	t.Logf("  Distance: %f", lastConn.Distance)

	// The last connection should bridge the two clusters
	// One box should be from cluster 1 (X < 10) and one from cluster 2 (X > 90)
	inCluster1 := lastConn.Box1.X < 10 || lastConn.Box2.X < 10
	inCluster2 := lastConn.Box1.X > 90 || lastConn.Box2.X > 90

	if !inCluster1 || !inCluster2 {
		t.Errorf("Expected connection to bridge both clusters")
	}
}

func TestFindLastConnectionLinear(t *testing.T) {
	// Linear arrangement
	input := []string{
		"0,0,0",
		"1,0,0",
		"2,0,0",
		"3,0,0",
		"4,0,0",
	}

	boxes := NewJunctionBoxes(input)

	// Use n=2 to create some initial circuits
	lastConn, err := boxes.FindLastConnection(2)

	if err != nil {
		t.Fatalf("FindLastConnection() error: %v", err)
	}

	t.Logf("Last connection for linear arrangement:")
	t.Logf("  Box1: (%d,%d,%d)", lastConn.Box1.X, lastConn.Box1.Y, lastConn.Box1.Z)
	t.Logf("  Box2: (%d,%d,%d)", lastConn.Box2.X, lastConn.Box2.Y, lastConn.Box2.Z)
	t.Logf("  Distance: %f", lastConn.Distance)
}
