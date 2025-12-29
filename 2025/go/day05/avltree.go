package day05

import (
	"fmt"
	"strconv"
	"strings"
)

// Interval represents a range [Start, End] inclusive
type Interval struct {
	Start, End int64
}

// IntervalNode represents a node in the AVL Interval Tree
type IntervalNode struct {
	Interval Interval
	Max      int64 // Maximum End value in this subtree
	Height   int   // Height of this node
	Left     *IntervalNode
	Right    *IntervalNode
}

// AVLIntervalTree is a self-balancing interval tree
type AVLIntervalTree struct {
	Root *IntervalNode
	Size int
}

// NewAVLIntervalTree creates a new empty AVL Interval Tree
func NewAVLIntervalTree() *AVLIntervalTree {
	return &AVLIntervalTree{}
}

// height returns the height of a node (0 for nil)
func (t *AVLIntervalTree) height(node *IntervalNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// maxValue returns the max value of a node (0 for nil)
func (t *AVLIntervalTree) maxValue(node *IntervalNode) int64 {
	if node == nil {
		return 0
	}
	return node.Max
}

// updateHeight updates the height of a node based on its children
func (t *AVLIntervalTree) updateHeight(node *IntervalNode) {
	if node == nil {
		return
	}
	leftHeight := t.height(node.Left)
	rightHeight := t.height(node.Right)
	node.Height = 1 + max(leftHeight, rightHeight)
}

// updateMax updates the max value of a node based on its interval and children
func (t *AVLIntervalTree) updateMax(node *IntervalNode) {
	if node == nil {
		return
	}
	maxVal := node.Interval.End
	if node.Left != nil && node.Left.Max > maxVal {
		maxVal = node.Left.Max
	}
	if node.Right != nil && node.Right.Max > maxVal {
		maxVal = node.Right.Max
	}
	node.Max = maxVal
}

// balanceFactor returns the balance factor (left height - right height)
func (t *AVLIntervalTree) balanceFactor(node *IntervalNode) int {
	if node == nil {
		return 0
	}
	return t.height(node.Left) - t.height(node.Right)
}

// rotateRight performs a right rotation
//
//	    y                    x
//	   / \                  / \
//	  x   C    ===>        A   y
//	 / \                      / \
//	A   B                    B   C
func (t *AVLIntervalTree) rotateRight(y *IntervalNode) *IntervalNode {
	x := y.Left
	B := x.Right

	// Perform rotation
	x.Right = y
	y.Left = B

	// Update heights
	t.updateHeight(y)
	t.updateHeight(x)

	// Update max values
	t.updateMax(y)
	t.updateMax(x)

	return x
}

// rotateLeft performs a left rotation
//
//	  x                      y
//	 / \                    / \
//	A   y      ===>        x   C
//	   / \                / \
//	  B   C              A   B
func (t *AVLIntervalTree) rotateLeft(x *IntervalNode) *IntervalNode {
	y := x.Right
	B := y.Left

	// Perform rotation
	y.Left = x
	x.Right = B

	// Update heights
	t.updateHeight(x)
	t.updateHeight(y)

	// Update max values
	t.updateMax(x)
	t.updateMax(y)

	return y
}

// Insert adds a new interval to the tree
func (t *AVLIntervalTree) Insert(interval Interval) {
	t.Root = t.insertNode(t.Root, interval)
	t.Size++
}

// insertNode recursively inserts an interval and rebalances
func (t *AVLIntervalTree) insertNode(node *IntervalNode, interval Interval) *IntervalNode {
	// Standard BST insertion
	if node == nil {
		return &IntervalNode{
			Interval: interval,
			Max:      interval.End,
			Height:   1,
		}
	}

	// Insert based on Start value
	if interval.Start < node.Interval.Start {
		node.Left = t.insertNode(node.Left, interval)
	} else {
		node.Right = t.insertNode(node.Right, interval)
	}

	// Update height and max
	t.updateHeight(node)
	t.updateMax(node)

	// Get balance factor
	balance := t.balanceFactor(node)

	// Left-Left case
	if balance > 1 && interval.Start < node.Left.Interval.Start {
		return t.rotateRight(node)
	}

	// Right-Right case
	if balance < -1 && interval.Start >= node.Right.Interval.Start {
		return t.rotateLeft(node)
	}

	// Left-Right case
	if balance > 1 && interval.Start >= node.Left.Interval.Start {
		node.Left = t.rotateLeft(node.Left)
		return t.rotateRight(node)
	}

	// Right-Left case
	if balance < -1 && interval.Start < node.Right.Interval.Start {
		node.Right = t.rotateRight(node.Right)
		return t.rotateLeft(node)
	}

	return node
}

// Contains checks if a value is within any interval in the tree
func (t *AVLIntervalTree) Contains(value int64) bool {
	return t.search(t.Root, value)
}

// search recursively searches for an interval containing the value
func (t *AVLIntervalTree) search(node *IntervalNode, value int64) bool {
	if node == nil {
		return false
	}

	// Check if current interval contains the value
	if value >= node.Interval.Start && value <= node.Interval.End {
		return true
	}

	// If left subtree exists and its max >= value, search left
	if node.Left != nil && node.Left.Max >= value {
		if t.search(node.Left, value) {
			return true
		}
	}

	// Search right subtree
	return t.search(node.Right, value)
}

// ParseInterval parses a string like "123-456" into an Interval
func ParseInterval(s string) (Interval, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return Interval{}, fmt.Errorf("invalid interval format: %s", s)
	}

	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return Interval{}, fmt.Errorf("invalid start value: %s", parts[0])
	}

	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return Interval{}, fmt.Errorf("invalid end value: %s", parts[1])
	}

	if start > end {
		return Interval{}, fmt.Errorf("start must be <= end: %d > %d", start, end)
	}

	return Interval{Start: start, End: end}, nil
}

// NewAVLIntervalTreeFromStrings creates a tree from a slice of interval strings
func NewAVLIntervalTreeFromStrings(intervals []string) (*AVLIntervalTree, error) {
	tree := NewAVLIntervalTree()
	for _, s := range intervals {
		interval, err := ParseInterval(s)
		if err != nil {
			return nil, err
		}
		tree.Insert(interval)
	}
	return tree, nil
}

// GetAllIntervals returns all intervals in sorted order (by Start value)
func (t *AVLIntervalTree) GetAllIntervals() []Interval {
	intervals := make([]Interval, 0, t.Size)
	t.inOrderTraversal(t.Root, &intervals)
	return intervals
}

// inOrderTraversal performs an in-order traversal to collect intervals
func (t *AVLIntervalTree) inOrderTraversal(node *IntervalNode, intervals *[]Interval) {
	if node == nil {
		return
	}
	t.inOrderTraversal(node.Left, intervals)
	*intervals = append(*intervals, node.Interval)
	t.inOrderTraversal(node.Right, intervals)
}

// MergeIntervals merges overlapping or adjacent intervals
// Input must be sorted by Start value (which GetAllIntervals provides)
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}

	merged := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1]
		curr := intervals[i]

		// Check if intervals overlap or are adjacent
		// Adjacent means curr.Start == last.End + 1
		if curr.Start <= last.End+1 {
			// Merge by extending the end if necessary
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			// No overlap, add as new interval
			merged = append(merged, curr)
		}
	}

	return merged
}

// CountTotalElements counts total unique elements across all merged intervals
func CountTotalElements(intervals []Interval) int64 {
	var total int64
	for _, interval := range intervals {
		// Interval [start, end] inclusive has (end - start + 1) elements
		total += interval.End - interval.Start + 1
	}
	return total
}

// CountUniqueElements returns the count of unique elements across all intervals in the tree
// This handles overlapping intervals by merging them first
func (t *AVLIntervalTree) CountUniqueElements() int64 {
	if t.Size == 0 {
		return 0
	}

	// Get all intervals in sorted order
	intervals := t.GetAllIntervals()

	// Merge overlapping intervals
	merged := MergeIntervals(intervals)

	// Count total elements
	return CountTotalElements(merged)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
