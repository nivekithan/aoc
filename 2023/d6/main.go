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

type RaceInfo struct {
	Time           int
	RecordDistance int
}

func read_input() []RaceInfo {
	fileName := "d6.data"

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	timeInfo := []int{}
	recordDistanceInfo := []int{}

	for lineScanner.Scan() {
		line := lineScanner.Text()

		if strings.HasPrefix(line, "Time:") {
			wordScanner := bufio.NewScanner(strings.NewReader(line))
			wordScanner.Split(bufio.ScanWords)

			for wordScanner.Scan() {
				word := wordScanner.Text()

				if word == "Time:" {
					continue
				}

				time, err := strconv.Atoi(word)

				if err != nil {
					log.Fatal(err)
				}

				timeInfo = append(timeInfo, time)
			}
		} else if strings.HasPrefix(line, "Distance:") {
			wordScanner := bufio.NewScanner(strings.NewReader(line))
			wordScanner.Split(bufio.ScanWords)

			for wordScanner.Scan() {
				word := wordScanner.Text()

				if word == "Distance:" {
					continue
				}

				time, err := strconv.Atoi(word)

				if err != nil {
					log.Fatal(err)
				}

				recordDistanceInfo = append(recordDistanceInfo, time)
			}
		}
	}

	raceInfo := []RaceInfo{}

	for i, time := range timeInfo {
		recordDistance := recordDistanceInfo[i]

		raceInfo = append(raceInfo, RaceInfo{Time: time, RecordDistance: recordDistance})
	}

	return raceInfo
}

func d5p1() {
	allRaceInfo := read_input()

	var waitGroup sync.WaitGroup

	result := make(chan int, len(allRaceInfo))

	for _, raceInfo := range allRaceInfo {
		waitGroup.Add(1)
		go func(raceInfo RaceInfo) {
			noWaysToWin := findNumberOfWaysToWin(&raceInfo)
			result <- noWaysToWin
			waitGroup.Done()
		}(raceInfo)
	}

	go func() {
		waitGroup.Wait()
		close(result)
	}()

	output := 1

	for value := range result {
		output *= value
	}

	fmt.Println(output)
}

func findNumberOfWaysToWin(raceInfo *RaceInfo) int {
	startingSpeed := findStartingSpeed(raceInfo)
	endingSpeed := findEndingSpeed(raceInfo)

	return endingSpeed - startingSpeed + 1
}

func findEndingSpeed(raceInfo *RaceInfo) int {
	return findEndingSpeedBs(1, raceInfo.Time-1, raceInfo)
}

func findEndingSpeedBs(minSpeed int, maxSpeed int, raceInfo *RaceInfo) int {
	speed := (minSpeed + maxSpeed) / 2

	isGreaterThanRecordDistance, isBeforeGreaterThanRecordDistance, direction := processSpeed(speed, raceInfo)

	if isGreaterThanRecordDistance && !isBeforeGreaterThanRecordDistance && direction == right {
		return speed
	}

	if !isBeforeGreaterThanRecordDistance && direction == left {
		return findEndingSpeedBs(speed, maxSpeed, raceInfo)
	} else if !isBeforeGreaterThanRecordDistance && direction == right {
		return findEndingSpeedBs(minSpeed, speed, raceInfo)
	} else if isGreaterThanRecordDistance {
		return findEndingSpeedBs(speed, maxSpeed, raceInfo)
	}

	log.Panic("Unreachable")
	return 0
}

func findStartingSpeed(raceInfo *RaceInfo) int {
	return findStartingSpeedBs(1, raceInfo.Time-1, raceInfo)

}

func findStartingSpeedBs(minSpeed int, maxSpeed int, raceInfo *RaceInfo) int {
	speed := (minSpeed + maxSpeed) / 2

	isGreaterThanRecordDistance, beforeIsGreaterThanRecordDistance, direction := processSpeed(speed, raceInfo)

	if isGreaterThanRecordDistance && !beforeIsGreaterThanRecordDistance && direction == left {
		return speed
	}

	if !isGreaterThanRecordDistance && direction == left {
		return findStartingSpeedBs(speed, maxSpeed, raceInfo)
	} else if !isGreaterThanRecordDistance && direction == right {
		return findStartingSpeedBs(minSpeed, speed, raceInfo)
	} else if isGreaterThanRecordDistance {
		return findStartingSpeedBs(minSpeed, speed, raceInfo)
	}

	log.Panic("Unreachable")
	return 0
}

type CurveDirection int

const (
	left CurveDirection = iota
	right
)

// Speed cannot be greater or equal to raceInfo.time
func processSpeed(speed int, raceInfo *RaceInfo) (bool, bool, CurveDirection) {
	time := raceInfo.Time - speed
	distance := speed * time

	isGreaterThanRecordDistance := distance > raceInfo.RecordDistance

	nextSpeed := speed + 1
	nextTime := raceInfo.Time - nextSpeed
	nextDistance := nextSpeed * nextTime

	if nextDistance > distance {
		beforeSpeed := speed - 1
		beforeTime := raceInfo.Time - beforeSpeed
		beforeDistance := beforeSpeed * beforeTime

		beforeIsGreaterThanRecordDistance := beforeDistance > raceInfo.RecordDistance
		return isGreaterThanRecordDistance, beforeIsGreaterThanRecordDistance, left
	} else {
		beforeSpeed := speed + 1
		beforeTime := raceInfo.Time - beforeSpeed
		beforeDistance := beforeSpeed * beforeTime

		beforeIsGreaterThanRecordDistance := beforeDistance > raceInfo.RecordDistance
		return isGreaterThanRecordDistance, beforeIsGreaterThanRecordDistance, right
	}
}

func main() {
	d5p1()
}
