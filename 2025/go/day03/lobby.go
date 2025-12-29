package day03

type Bank string

type Joltage int

func (b Bank) LargestJoltage(cellsToKeep int) Joltage {
	if cellsToKeep >= len(b) {
		// Convert entire bank to Joltage
		joltage := Joltage(0)
		for _, d := range b {
			joltage = joltage*10 + Joltage(d-'0')
		}
		return joltage
	}
	if cellsToKeep <= 0 {
		return 0
	}

	cellsToRemove := len(b) - cellsToKeep
	stack := make([]byte, 0, len(b))
	removed := 0

	for i := 0; i < len(b); i++ {
		// While we can still remove digits and current digit is larger than top of stack
		for len(stack) > 0 && removed < cellsToRemove && b[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			removed++
		}
		stack = append(stack, b[i])
	}

	// Trim to the number we need to keep
	stack = stack[:cellsToKeep]

	// Convert result to Joltage
	joltage := Joltage(0)
	for _, d := range stack {
		joltage = joltage*10 + Joltage(d-'0')
	}

	return joltage
}
