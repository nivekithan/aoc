package main

import (
	"fmt"
	"log"
)

type Pipe = int

const (
	NorthSouth Pipe = iota
	EastWest
	NorthEast
	NorthWest
	SouthWest
	SouthEast
	Ground
	S
)

var RuneToPipe = map[rune]Pipe{
	'|': NorthSouth,
	'-': EastWest,
	'L': NorthEast,
	'J': NorthWest,
	'7': SouthWest,
	'F': SouthEast,
	'.': Ground,
	'S': S,
}

type Grid struct {
	grid            []Pipe
	colLength       int
	startingPostion int
}

type Cell struct {
	row int
	col int
}

func NewGrid(colLength int) *Grid {

	return &Grid{grid: []Pipe{}, colLength: colLength}
}

func (g *Grid) Append(p Pipe) {
	g.grid = append(g.grid, p)

	if p == S {
		pos := len(g.grid) - 1
		g.startingPostion = pos
	}
}

func (g *Grid) Get(cell Cell) Pipe {
	pos := g.toPos(cell)
	return g.grid[pos]
}

func (g *Grid) Walk() {
	distanceFromStartingPostion := make([]int, len(g.grid))

	startingCell := g.toCell(g.startingPostion)

	fromNorth, isValid := g.south(startingCell)

	if isValid {
		canWalkFromNorth := g.checkCanWalk(fromNorth, North)

		if canWalkFromNorth {
			g.walkImpl(fromNorth, North, 1, &distanceFromStartingPostion)
		}

		log.Printf("CanWalkFromNorth %v", canWalkFromNorth)
	}

	fromSouth, isValid := g.north(startingCell)

	if isValid {
		canWalkFromSouth := g.checkCanWalk(fromSouth, South)

		if canWalkFromSouth {
			g.walkImpl(fromSouth, South, 1, &distanceFromStartingPostion)
		}

		log.Printf("CanWalkFromSouth %v", canWalkFromSouth)
	}
	fromWest, isValid := g.east(startingCell)

	if isValid {
		canWalkFromWest := g.checkCanWalk(fromWest, West)

		if canWalkFromWest {
			g.walkImpl(fromWest, West, 1, &distanceFromStartingPostion)
		}

		log.Printf("CanWalkFromWest %v", canWalkFromWest)
	}

	fromEast, isValid := g.west(startingCell)

	if isValid {
		canWalkFromEast := g.checkCanWalk(fromEast, East)

		if canWalkFromEast {
			g.walkImpl(fromEast, East, 1, &distanceFromStartingPostion)

		}

		log.Printf("CanWalkFromEast %v", canWalkFromEast)
	}

	maxDistance := 0

	for _, dis := range distanceFromStartingPostion {
		maxDistance = max(dis, maxDistance)
	}

	fmt.Println(maxDistance)
}

func (g *Grid) walkImpl(cell Cell, d direction, steps int, stepsCollection *[]int) bool {
	currentPipe := g.Get(cell)

	if currentPipe == Ground {
		return false
	} else if currentPipe == S {
		return true
	}

	pos := g.toPos(cell)

	if (*stepsCollection)[pos] == 0 {
		(*stepsCollection)[pos] = steps
	} else {
		(*stepsCollection)[pos] = min((*stepsCollection)[pos], steps)
	}

	var nextCell Cell
	var nextD direction

	if currentPipe == NorthSouth {
		if d == North {
			nextCell, _ = g.south(cell)
			nextD = North
		} else {
			nextCell, _ = g.north(cell)
			nextD = South
		}
	} else if currentPipe == EastWest {
		if d == East {
			nextCell, _ = g.west(cell)
			nextD = East
		} else {
			nextCell, _ = g.east(cell)
			nextD = West
		}
	} else if currentPipe == NorthEast {
		if d == North {
			nextCell, _ = g.east(cell)
			nextD = West
		} else {
			nextCell, _ = g.north(cell)
			nextD = South
		}
	} else if currentPipe == NorthWest {
		if d == North {
			nextCell, _ = g.west(cell)
			nextD = East
		} else {
			nextCell, _ = g.north(cell)
			nextD = South
		}
	} else if currentPipe == SouthWest {
		if d == South {
			nextCell, _ = g.west(cell)
			nextD = East
		} else {
			nextCell, _ = g.south(cell)
			nextD = North
		}
	} else if currentPipe == SouthEast {
		if d == South {
			nextCell, _ = g.east(cell)
			nextD = West
		} else {
			nextCell, _ = g.south(cell)
			nextD = North
		}
	}

	return g.walkImpl(nextCell, nextD, steps+1, stepsCollection)
}

type direction = int

const (
	North direction = iota
	South
	East
	West
)

func (g *Grid) checkCanWalk(cell Cell, d direction) bool {
	currentPipe := g.Get(cell)

	if d == North {
		canTravel := currentPipe == NorthSouth || currentPipe == NorthEast || currentPipe == NorthWest
		return canTravel
	} else if d == South {
		canTravel := currentPipe == SouthEast || currentPipe == SouthWest || currentPipe == NorthSouth
		return canTravel
	} else if d == East {
		canTravel := currentPipe == EastWest || currentPipe == NorthEast || currentPipe == SouthEast
		return canTravel
	} else if d == West {
		canTravel := currentPipe == EastWest || currentPipe == NorthWest || currentPipe == SouthWest
		return canTravel
	}

	log.Fatal("Unreachable")
	return false
}

func (g *Grid) south(cell Cell) (Cell, bool) {
	maxRow := g.maxRow()

	if cell.row == maxRow {
		return Cell{}, false
	}

	return Cell{col: cell.col, row: cell.row + 1}, true
}

func (g *Grid) north(cell Cell) (Cell, bool) {

	if cell.row == 0 {
		return Cell{}, false
	}

	return Cell{col: cell.col, row: cell.row - 1}, true
}

func (g *Grid) east(cell Cell) (Cell, bool) {
	maxCol := g.maxCol()

	if cell.col == maxCol {
		return Cell{}, false
	}

	return Cell{col: cell.col + 1, row: cell.row}, true
}

func (g *Grid) west(cell Cell) (Cell, bool) {

	if cell.col == 0 {
		return Cell{}, false
	}

	return Cell{col: cell.col - 1, row: cell.row}, true
}

func (g *Grid) toCell(pos int) Cell {
	row := pos / g.colLength
	col := pos % g.colLength

	return Cell{row: row, col: col}
}

func (g *Grid) toPos(cell Cell) int {
	return (g.colLength * cell.row) + cell.col
}

func (g *Grid) maxCol() int {
	return g.colLength - 1
}

func (g *Grid) maxRow() int {
	return (len(g.grid) / g.colLength) - 1
}
