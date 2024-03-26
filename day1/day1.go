package day1

import (
	"fmt"
	"strings"

	"cfagudelo/advent-code-2023/data"
)

const (
	zeroRune = 48
	nineRune = 57
)

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

func SumCalibrationValues() (int, error) {
	scanner, err := data.GetScanner(1)
	if err != nil {
		return 0, fmt.Errorf("getting scanner: %w", err)
	}
	var sum int
	for scanner.Scan() {
		n := extractCalibrationNumbersFromLine(scanner.Text())
		sum += n
	}
	return sum, nil
}

func SumCalibrationValuesV2() (int, error) {
	scanner, err := data.GetScanner(1)
	if err != nil {
		return 0, fmt.Errorf("getting scanner: %w", err)
	}
	var sum int
	for scanner.Scan() {
		n := extractCalibrationNumbersFromLineV2(scanner.Text())
		sum += n
	}
	return sum, nil
}

func extractCalibrationNumbersFromLine(line string) int {
	var first, last int
	var lastAssigned bool
	for _, c := range line {
		if c >= zeroRune && c <= nineRune {
			digit := int(c - zeroRune)
			if first == 0 {
				first = digit
				continue
			}
			last = digit
			lastAssigned = true
		}
	}
	if !lastAssigned {
		return first*10 + first
	}
	return first*10 + last
}

func extractCalibrationNumbersFromLineV2(line string) int {
	var first, last int
	for i, c := range line {
		if c >= zeroRune && c <= nineRune {
			digit := int(c - zeroRune)
			if first == 0 {
				first = digit
			}
			last = digit
		}
		digit, ok := prefixToDigit(line[i:])
		if ok {
			if first == 0 {
				first = digit
			}
			last = digit
		}
	}
	return first*10 + last
}

//nolint:cyclop,gomnd // unnecessary to fix linter
func prefixToDigit(s string) (int, bool) {
	switch {
	case strings.HasPrefix(s, one):
		return 1, true
	case strings.HasPrefix(s, two):
		return 2, true
	case strings.HasPrefix(s, three):
		return 3, true
	case strings.HasPrefix(s, four):
		return 4, true
	case strings.HasPrefix(s, five):
		return 5, true
	case strings.HasPrefix(s, six):
		return 6, true
	case strings.HasPrefix(s, seven):
		return 7, true
	case strings.HasPrefix(s, eight):
		return 8, true
	case strings.HasPrefix(s, nine):
		return 9, true
	default:
		return 0, false
	}
}
