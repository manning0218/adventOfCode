package day04

type GridPrintDept [][]rune
type RollLocations [][2]int

var directions = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1}, // northwest, north, northeast
	{0, -1}, {0, 1}, // west,        , east
	{1, -1}, {1, 0}, {1, 1}, // southwest, south, southeast
}

func NewGridPrintDept(lines []string) GridPrintDept {
	height := len(lines)
	if height == 0 {
		return GridPrintDept{}
	}
	width := len(lines[0])
	dept := make(GridPrintDept, height)
	for i := 0; i < height; i++ {
		dept[i] = make([]rune, width)
		for j, ch := range lines[i] {
			dept[i][j] = ch
		}
	}
	return dept
}

func (g GridPrintDept) NumberOfNeighbors(row, col int, target rune) int {
	count := 0
	height := len(g)
	if height == 0 {
		return 0
	}
	width := len(g[0])

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newRow < height && newCol >= 0 && newCol < width {
			if g[newRow][newCol] == target {
				count++
			}
		}
	}
	return count
}

func (g GridPrintDept) FindNumberPaperToMove(maxNeighbors int) RollLocations {
	locations := RollLocations{}
	height := len(g)
	if height == 0 {
		return nil
	}
	width := len(g[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if g[i][j] != '@' {
				continue
			}
			neighborCount := g.NumberOfNeighbors(i, j, '@')
			if neighborCount < maxNeighbors {
				locations = append(locations, [2]int{i, j})
			}
		}
	}
	return locations
}

func (g GridPrintDept) RemoveRollLocations(locations RollLocations) {
	for _, loc := range locations {
		row, col := loc[0], loc[1]
		if g[row][col] == '@' {
			g[row][col] = '.'
		}
	}
}
