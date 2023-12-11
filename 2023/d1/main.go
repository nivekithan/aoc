package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {

	file, err := os.Open("d1.data")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		first, err := first_digit(line)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%v", line)
		log.Printf("First: %v", first)

		last, err := last_digit(line)

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("last : %v", last)

		sum += first*10 + last
	}

	fmt.Println(sum)
}

func first_digit(line string) (int, error) {
	left_window := 0
	right_window := 0

	for left_window <= right_window && right_window < len(line) {
		isDigit := unicode.IsDigit(rune(line[right_window]))

		if isDigit {
			return int(rune(line[right_window]) - '0'), nil
		}

		slice := line[left_window : right_window+1]
		is_partially_correct, digit := IsParitalDigit(slice)

		if !is_partially_correct && digit.Value != -1 {
			return digit.Value, nil
		} else if is_partially_correct {
			right_window += 1
		} else if !is_partially_correct && digit.Value == -1 {
			left_window += 1
			right_window = max(left_window, right_window)
		}
	}

	return 0, fmt.Errorf("Unable to find any digit on: %v", line)

}

func last_digit(line string) (int, error) {
	left_window := len(line) - 1
	right_window := len(line) - 1

	for right_window >= left_window && left_window >= 0 {
		isDigit := unicode.IsDigit(rune(line[left_window]))

		if isDigit {
			return int(rune(line[left_window]) - '0'), nil
		}

		slice := reverse_str(line[left_window : right_window+1])
		is_partially_correct, digit := RevIsParitalDigit(slice)

		if !is_partially_correct && digit.Value != -1 {
			return digit.Value, nil
		} else if is_partially_correct {
			left_window -= 1
		} else if !is_partially_correct && digit.Value == -1 {
			right_window -= 1
			left_window = min(left_window, right_window)
		}
	}

	return 0, fmt.Errorf("Unable to find any digit on: %v", line)
}

func reverse_str(str string) string {
	reversed := make([]byte, len(str))
	i := 0

	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		str = str[:len(str)-size]
		i += utf8.EncodeRune(reversed[i:], r)
	}

	return string(reversed)

}
