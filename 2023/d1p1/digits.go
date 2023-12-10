package main

import (
	"fmt"
	"strings"
)

var valid_digit = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

type Digit struct {
	Value int
}

func ToDigit(str string) (Digit, error) {
	switch str {
	case "1", "one":
		return Digit{Value: 1}, nil
	case "2", "two":
		return Digit{Value: 2}, nil
	case "3", "three":
		return Digit{Value: 3}, nil
	case "4", "four":
		return Digit{Value: 4}, nil
	case "5", "five":
		return Digit{Value: 5}, nil
	case "6", "six":
		return Digit{Value: 6}, nil
	case "7", "seven":
		return Digit{Value: 7}, nil
	case "8", "eight":
		return Digit{Value: 8}, nil
	case "9", "nine":
		return Digit{Value: 9}, nil
	}

	return Digit{}, fmt.Errorf("Cannot convert: %v to Digit", str)
}

func RevToDigit(str string) (Digit, error) {
	switch str {
	case "1", reverse_str("one"):
		return Digit{Value: 1}, nil
	case "2", reverse_str("two"):
		return Digit{Value: 2}, nil
	case "3", reverse_str("three"):
		return Digit{Value: 3}, nil
	case "4", reverse_str("four"):
		return Digit{Value: 4}, nil
	case "5", reverse_str("five"):
		return Digit{Value: 5}, nil
	case "6", reverse_str("six"):
		return Digit{Value: 6}, nil
	case "7", reverse_str("seven"):
		return Digit{Value: 7}, nil
	case "8", reverse_str("eight"):
		return Digit{Value: 8}, nil
	case "9", reverse_str("nine"):
		return Digit{Value: 9}, nil
	}

	return Digit{Value: -1}, fmt.Errorf("Cannot convert: %v to Digit", str)
}

func IsParitalDigit(str string) (bool, Digit) {
	digit, err := ToDigit(str)

	if err == nil {
		return false, digit
	}

	for _, v := range valid_digit {
		correct_partial_str := strings.HasPrefix(v, str)

		if correct_partial_str {
			return true, Digit{Value: -1}
		}
	}

	return false, Digit{Value: -1}
}

var rev_valid_digit = []string{}

func RevIsParitalDigit(str string) (bool, Digit) {

	if len(rev_valid_digit) == 0 {
		for _, v := range valid_digit {
			rev_valid_digit = append(rev_valid_digit, reverse_str(v))
		}
	}

	digit, err := RevToDigit(str)

	if err == nil {
		return false, digit
	}

	for _, v := range rev_valid_digit {
		correct_partial_str := strings.HasPrefix(v, str)

		if correct_partial_str {
			return true, Digit{Value: -1}
		}
	}

	return false, Digit{Value: -1}
}
