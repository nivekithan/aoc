package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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

func readInput2() (string, map[string]Direction, []string) {
	fileName := "d8.data"

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var direction string

	networks := map[string]Direction{}

	startingLocation := []string{}

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

		if strings.HasSuffix(from, "A") {
			startingLocation = append(startingLocation, from)
		}

		toSplit := strings.Split(splittedLine[1], ", ")

		toLeft := toSplit[0][1:]
		toRight := toSplit[1][:3]

		networks[from] = Direction{left: toLeft, right: toRight}
	}

	return direction, networks, startingLocation
}
func part2() {
	direction, networks, curLocation := readInput2()

	var waitGroup sync.WaitGroup

	c := make(chan int, len(curLocation))

	for _, location := range curLocation {
		waitGroup.Add(1)
		go func(location string) {
			stepsTaken := 0
			for !strings.HasSuffix(location, "Z") {
				for _, d := range direction {
					isLeft := d == 'L'

					if isLeft {
						location = networks[location].left
					} else {
						location = networks[location].right
					}

					stepsTaken++
					if strings.HasSuffix(location, "Z") {
						break
					}

				}
			}
			waitGroup.Done()
			c <- stepsTaken
		}(location)
	}

	go func() {
		waitGroup.Wait()
		close(c)
	}()

	stepsTaken := 0
	for value := range c {

		if stepsTaken == 0 {
			stepsTaken = value
			continue
		}

		stepsTaken = Lcm(stepsTaken, value)
	}

	fmt.Println(stepsTaken)

}

func Lcm(first int, second int) int {

	var gcd int

	for i := 1; i <= first && i <= second; i++ {
		if first%i == 0 && second%i == 0 {
			gcd = i
		}
	}

	lcm := (first * second) / gcd

	return lcm
}

func main() {
	part2()
}
