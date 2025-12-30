package day06

import (
	"fmt"
	"strings"
)

type Column struct {
	Data      []int64
	Operation rune // '+' for addition, '*' for multiplication
}

type Columns []Column

func ParseColumns(lines []string) (Columns, error) {
	if len(lines) == 0 {
		return nil, nil
	}

	operations := strings.Fields(lines[len(lines)-1])
	columns := make([]Column, len(operations))

	for i, operation := range operations {
		columns[i].Data = make([]int64, len(lines)-1)
		columns[i].Operation = rune(operation[0])
	}

	for row, line := range lines[:len(lines)-1] {
		// Parsing logic for each line to create Column instances
		data := strings.Fields(line)

		for col, str := range data {
			var value int64
			_, err := fmt.Sscanf(str, "%d", &value)
			if err != nil {
				return nil, fmt.Errorf("failed to parse integer '%s': %w", str, err)
			}
			columns[col].Data[row] = value

		}
	}

	fmt.Println("Parsed columns:", columns)

	return columns, nil
}

func ParseCephalopod(input []string) (Columns, error) {
	rows := len(input)
	if rows == 0 {
		return nil, fmt.Errorf("input is empty")
	}
	bits := len(input[0])

	data := []int64{}
	columns := []Column{}
	for bit := bits - 1; bit >= 0; bit-- {
		value := int64(0)
		operation := ' '
		for row := 0; row < rows; row++ {
			if input[row][bit] == ' ' {
				continue
			}
			switch input[row][bit] {
			case '+', '*':
				operation = rune(input[row][bit])
			default:
				digit := input[row][bit] - '0'
				value = int64(digit) + (value * 10)
			}
		}
		if value == 0 {
			continue
		}
		data = append(data, value)
		if operation != ' ' {
			columns = append(columns, Column{
				Data:      data,
				Operation: operation,
			})
			data = []int64{}
		}

	}

	return columns, nil
}

func (c Column) ApplyOperation() int64 {
	if len(c.Data) == 0 {
		return 0
	}

	result := c.Data[0]
	for _, value := range c.Data[1:] {
		switch c.Operation {
		case '+':
			result += value
		case '*':
			result *= value
		}
	}
	return result
}

func (c Columns) ComputeResults() []int64 {
	results := make([]int64, len(c))
	for i := range c {
		results[i] = c[i].ApplyOperation()
	}

	return results
}
