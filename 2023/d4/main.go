package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func d4p1() {
	file_name := "d4.data"

	file, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		line_scanner := bufio.NewScanner(strings.NewReader(line))

		line_scanner.Split(bufio.ScanWords)

		winning_nums := map[int]bool{}

		is_scanning_winning_num := true

		winning_points := 0
		for line_scanner.Scan() {
			word := line_scanner.Text()

			if word == "|" {
				is_scanning_winning_num = false
				continue
			}

			num, err := strconv.Atoi(word)

			if err != nil {
				continue
			}

			if is_scanning_winning_num {
				winning_nums[num] = true
				continue
			}

			if winning_nums[num] {
				if winning_points == 0 {
					winning_points = 1
				} else {
					winning_points *= 2
				}
			}

		}

		sum += winning_points
	}

	fmt.Println(sum)
}

func main() {
	d4p1()
}
