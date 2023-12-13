package main

import (
	"fmt"
	"strconv"
)

type EngineIterator struct {
	content *[][]rune
	cur_pos Pos
}

func NewEngineIterator(content *[][]rune) EngineIterator {
	return EngineIterator{content: content, cur_pos: Pos{row: 0, col: 0}}
}

func (engine *EngineIterator) findCanContine() bool {

	max_col := len((*engine.content)[0]) - 1

	if engine.cur_pos.col >= max_col {
		max_row := len(*engine.content) - 1

		if engine.cur_pos.row >= max_row {
			return false
		}
	}
	return true
}

func (engine *EngineIterator) next() (rune, bool) {

	can_contine := engine.findCanContine()

	if !can_contine {
		return 0, false
	}

	max_col := len((*engine.content)[0]) - 1

	if engine.cur_pos.col >= max_col {
		max_row := len(*engine.content) - 1

		if engine.cur_pos.row >= max_row {
			return 0, false
		}

		engine.cur_pos.row += 1
		engine.cur_pos.col = 0
	} else {
		engine.cur_pos.col += 1
	}

	cur_value := (*engine.content)[engine.cur_pos.row][engine.cur_pos.col]
	return cur_value, true
}

func (engine *EngineIterator) cur() rune {
	return (*engine.content)[engine.cur_pos.row][engine.cur_pos.col]
}

func (engine *EngineIterator) SkipDot() {
	for engine.findCanContine() {
		cur_rune := engine.cur()

		if cur_rune != '.' {
			break
		}

		engine.next()

	}
}

func (engine *EngineIterator) ParseNum() (int, *[]int, error) {

	num_in_str := ""

	num_cols := []int{}

	cur_rune := engine.cur()

	for is_digit(cur_rune) {
		// log.Printf("cur_rune: %v, cur_rune_in_str: %v", cur_rune, string(cur_rune))
		num_in_str += string(cur_rune)
		num_cols = append(num_cols, engine.cur_pos.col)

		next_run, can_continue := engine.next()

		if !can_continue {
			break
		}

		cur_rune = next_run
	}

	if num_in_str == "" {
		return 0, &num_cols, fmt.Errorf("No number found in cur_pos")
	}

	num, err := strconv.Atoi(num_in_str)

	return num, &num_cols, err
}
