package main

import (
	"bufio"
	"log"
	"os"
)

func readInput() *Grid {
	fileName := "d10.data"

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid *Grid

	for scanner.Scan() {
		line := scanner.Text()

		if grid == nil {
			grid = NewGrid(len(line))
		}

		for _, r := range line {
			pipe := RuneToPipe[r]

			grid.Append(pipe)
		}
	}

	return grid
}

func part1() {
	grid := readInput()

	grid.Walk()
}

func main() {
	part1()

}
