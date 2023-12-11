package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func d1p1() {
	file_name := "d2.data"

	file, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maximum_cubes := Cubes{Red: 12, Green: 13, Blue: 14}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Print(line)

		game_no, cubes_info_str := find_game_no(line)

		log.Printf("Game no: %v", game_no)
		log.Printf("Cubes info str: %v", cubes_info_str)

		cubes := find_cubes_in_each_round(cubes_info_str)

		is_game_impossible := false
		for _, cube := range cubes {
			log.Printf("Cube.Red %v, Cube.Green %v, Cube.Blue %v", cube.Red, cube.Green, cube.Blue)

			is_round_impossible := maximum_cubes.Blue < cube.Blue || maximum_cubes.Green < cube.Green || maximum_cubes.Red < cube.Red

			if !is_round_impossible {
				log.Printf("Round Possible")
			} else {
				log.Printf("Round impossible")
				is_game_impossible = true
				break
			}
		}

		if !is_game_impossible {
			log.Printf("Game Possible: %v", game_no)
			sum += game_no
		} else {
			log.Printf("Game impossible: %v", game_no)
		}
	}

	fmt.Println(sum)
}

func find_game_no(str string) (int, string) {
	splitted_line := strings.Split(str, ":")

	game_str := splitted_line[0]

	splitted_game_str := strings.Split(game_str, " ")

	game_no_str := splitted_game_str[1]

	game_no, err := strconv.Atoi(game_no_str)

	if err != nil {
		log.Fatal(err)
	}

	return game_no, splitted_line[1]
}

func find_cubes_in_each_round(str string) []Cubes {
	each_round_cubes := strings.Split(str, ";")

	cubes := make([]Cubes, len(each_round_cubes))

	for i, v := range each_round_cubes {
		cube := NewCubesFromString(v)
		cubes[i] = cube
	}

	return cubes
}

func main() {
	d1p1()
}
