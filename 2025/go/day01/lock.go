package day01

import "fmt"

type Lock struct {
	Val  int
	Prev *Lock
	Next *Lock
}

var password int = 0

func ResetPassword() {
	password = 0
}

func NewLock(maxVal, startingPosition int) *Lock {
	head := &Lock{Val: 0}
	current := head
	for i := 1; i < maxVal; i++ {
		newLock := &Lock{Val: i, Prev: current}
		current.Next = newLock
		current = newLock
	}
	current.Next = head
	head.Prev = current

	for i := 0; i < startingPosition; i++ {
		head = head.Next
	}

	return head
}

func PrintLock(l *Lock, size int) {
	current := l
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
}

func (l *Lock) MoveLeft(steps int) *Lock {
	current := l
	for i := 0; i < steps; i++ {
		current = current.Prev
		if current.Val == 0 {
			password++
		}
	}
	return current
}

func (l *Lock) MoveRight(steps int) *Lock {
	current := l
	for i := 0; i < steps; i++ {
		current = current.Next
		if current.Val == 0 {
			password++
		}
	}
	return current
}

func GetPassword() int {
	return password
}

func (l *Lock) IsMagicNumber(magicNumbers ...int) bool {
	for _, num := range magicNumbers {
		if l.Val == num {
			return true
		}
	}
	return false
}

// IsInvalid2 checks if a number is made of a repeating sequence (e.g., 12121212, 123123123)
func IsInvalid2(id int) bool {
	s := fmt.Sprintf("%d", id)
	n := len(s)

	// Try all possible pattern lengths from 1 to n/2
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen != 0 {
			continue
		}

		pattern := s[:patternLen]
		repeats := n / patternLen
		repeated := ""
		for i := 0; i < repeats; i++ {
			repeated += pattern
		}

		if repeated == s {
			return true
		}
	}

	return false
}
