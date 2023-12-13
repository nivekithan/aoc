package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
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

func par2() {

	content := read_input()

	max_pos := Pos{row: len(*content) - 1, col: len((*content)[0]) - 1}

	board := make([]string, (max_pos.row+1)*(max_pos.row+1))
	id_to_num := map[string]int{}

	engine_iterator := NewEngineIterator(content)

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
			rand_id := uuid.NewString()
			id_to_num[rand_id] = num
			for _, col := range *num_cols {
				board[(row*(max_pos.row+1))+col] = rand_id
			}
		}
	}

	engine_iterator = NewEngineIterator(content)

	sum := 0
	for ; engine_iterator.findCanContine(); engine_iterator.next() {
		cur_rune := engine_iterator.cur()

		if cur_rune != '*' {
			continue
		}

		all_adjs := get_adjacent(Pos{row: engine_iterator.cur_pos.row, col: engine_iterator.cur_pos.col}, max_pos)

		all_parts := map[string]bool{}
		for _, pos := range all_adjs {
			part_id := board[(pos.row*(max_pos.row+1) + (pos.col))]

			if part_id != "" {
				all_parts[part_id] = true
			}
		}

		all_part_num := []int{}
		for k := range all_parts {

			part_num := id_to_num[k]

			if part_num != 0 {
				all_part_num = append(all_part_num, part_num)
			}
		}

		if len(all_part_num) == 2 {
			sum += (all_part_num[0] * all_part_num[1])
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
	par2()
}
