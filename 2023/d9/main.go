package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Reading = []int

func readInput() []Reading {
	fileName := "d9.data"

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	allReadings := []Reading{}
	for scanner.Scan() {
		line := scanner.Text()
		readingInStr := strings.Split(line, " ")

		readings := Reading{}

		for _, readingInStr := range readingInStr {
			reading, err := strconv.Atoi(readingInStr)

			if err != nil {
				log.Fatal(err)
			}

			readings = append(readings, reading)
		}

		allReadings = append(allReadings, readings)
	}

	return allReadings
}

func part1() {
	allReadings := readInput()

	var waitGroup sync.WaitGroup

	c := make(chan int, len(allReadings))
	for _, readings := range allReadings {

		waitGroup.Add(1)
		go func(readings []int) {
			nextValue := predictNextValue(readings)
			c <- nextValue
			waitGroup.Done()
		}(readings)
	}

	go func() {
		waitGroup.Wait()
		close(c)
	}()

	sum := 0

	for value := range c {
		sum += value
	}
	fmt.Println(sum)
}

func predictNextValue(reading Reading) int {

	readingLastValues := []int{reading[len(reading)-1]}
	lastReading := reading
	for {
		difference, isSame := findDifference(lastReading)

		readingLastValues = append(readingLastValues, difference[len(difference)-1])

		if isSame {
			break
		}

		lastReading = difference
	}

	sum := 0

	for _, lastValue := range readingLastValues {
		sum += lastValue
	}

	return sum
}

func findDifference(reading Reading) (Reading, bool) {
	differenceReading := Reading{}

	isSame := true

	for i, readingValue := range reading {
		if i == len(reading)-1 {
			continue
		}

		difference := reading[i+1] - readingValue

		differenceReading = append(differenceReading, difference)

		if len(differenceReading) == 1 || !isSame {
			continue
		}

		lastValue := differenceReading[len(differenceReading)-1]
		lastBeforeValue := differenceReading[len(differenceReading)-2]

		if lastBeforeValue != lastValue {
			isSame = false
			continue
		}
	}

	return differenceReading, isSame
}

func main() {
	part1()
}
