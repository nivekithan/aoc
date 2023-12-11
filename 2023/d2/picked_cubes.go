package main

import (
	"log"
	"strconv"
	"strings"
)

type Cubes struct {
	Red   int
	Green int
	Blue  int
}

func NewCubesFromString(str string) Cubes {
	colors := strings.Split(str, ",")

	log.Printf("Colors: %v", colors)

	cube := Cubes{}
	for _, v := range colors {
		v = strings.TrimSpace(v)

		num_and_color := strings.Split(v, " ")

		num, err := strconv.Atoi(num_and_color[0])

		if err != nil {
			log.Fatal(err)
		}

		color := num_and_color[1]

		switch color {
		case "red":
			cube.Red = num
		case "green":
			cube.Green = num
		case "blue":
			cube.Blue = num
		}
	}

	return cube
}
