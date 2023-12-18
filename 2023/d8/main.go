package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction struct {
	left  string
	right string
}

func readInput() (string, map[string]Direction) {
	fileName := "d8.data"

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var direction string

	networks := map[string]Direction{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Empty line
			continue
		}

		isDirection := len(line) != 16

		if isDirection {
			// Direction line
			direction = line
			continue
		}

		splittedLine := strings.Split(line, " = ")

		from := splittedLine[0]

		toSplit := strings.Split(splittedLine[1], ", ")

		toLeft := toSplit[0][1:]
		toRight := toSplit[1][:3]

		networks[from] = Direction{left: toLeft, right: toRight}
	}

	return direction, networks
}

func part1() {
	direction, networks := readInput()

	curLocation := "AAA"
	destLocation := "ZZZ"

	stepsTaken := 0

	for curLocation != destLocation {

		for _, d := range direction {
			isLeft := d == 'L'

			if isLeft {
				curLocation = networks[curLocation].left
			} else {
				curLocation = networks[curLocation].right
			}

			stepsTaken++

			if curLocation == destLocation {
				break
			}
		}
	}

	fmt.Println(stepsTaken)
}

func main() {
	part1()
}
