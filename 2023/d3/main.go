package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_input() *[][]rune {
	file, err := os.Open("d3p1.data")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	content := [][]rune{}

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		content = append(content, []rune{})
		for _, r := range line {
			content[row] = append(content[row], r)
		}

		row++
	}

	return &content
}

func par1() {

	content := read_input()

	max_pos := Pos{row: len(*content) - 1, col: len((*content)[0]) - 1}

	engine_iterator := NewEngineIterator(content)

	sum := 0
	for engine_iterator.findCanContine() {
		engine_iterator.SkipDot()

		row := engine_iterator.cur_pos.row
		num, num_cols, err := engine_iterator.ParseNum()

		if err != nil {
			engine_iterator.next()
			continue
		}

		is_part_number := false

		for _, col := range *num_cols {
			pos := Pos{row: row, col: col}

			if is_part_pos(pos, max_pos, content) {
				is_part_number = true
				break
			}
		}

		if is_part_number {
			sum += num

		}
	}

	fmt.Println(sum)
}

func is_part_pos(pos Pos, max_pos Pos, content *[][]rune) bool {
	adjs := get_adjacent(pos, max_pos)

	for _, pos := range adjs {
		rune := (*content)[pos.row][pos.col]

		is_rune_symbol := is_symbol(rune)

		if is_rune_symbol {
			return true
		}
	}

	return false
}

func is_symbol(r rune) bool {
	return r != '.' && !is_digit(r)
}

func is_digit(r rune) bool {
	return r >= '0' && r <= '9'
}

func get_adjacent(cur_pos, max_pos Pos) []Pos {

	adj := []Pos{}

	if cur_pos.col != 0 {
		adj = append(adj, Pos{col: cur_pos.col - 1, row: cur_pos.row})
	}

	if cur_pos.col != 0 && cur_pos.row != 0 {
		adj = append(adj, Pos{row: cur_pos.row - 1, col: cur_pos.col - 1})
	}

	if cur_pos.row != 0 {
		adj = append(adj, Pos{row: cur_pos.row - 1, col: cur_pos.col})
	}

	if cur_pos.row != 0 && cur_pos.col != max_pos.col {
		adj = append(adj, Pos{row: cur_pos.row - 1, col: cur_pos.col + 1})
	}

	if cur_pos.col != max_pos.col {
		adj = append(adj, Pos{row: cur_pos.row, col: cur_pos.col + 1})
	}

	if cur_pos.col != max_pos.col && cur_pos.row != max_pos.row {
		adj = append(adj, Pos{row: cur_pos.row + 1, col: cur_pos.col + 1})
	}

	if cur_pos.row != max_pos.row {
		adj = append(adj, Pos{row: cur_pos.row + 1, col: cur_pos.col})
	}

	if cur_pos.row != max_pos.row && cur_pos.col != 0 {
		adj = append(adj, Pos{row: cur_pos.row + 1, col: cur_pos.col - 1})
	}

	return adj
}

func main() {
	par1()
}
