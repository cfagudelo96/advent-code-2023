package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"cfagudelo/advent-code-2023/data"
)

const (
	day = 4
)

var cardRegexp = regexp.MustCompile(`^Card[ ]*(\d+): .*$`)

type card struct {
	number  int
	results []int
	winMap  map[int]bool
}

func (c card) numberMatches() int {
	var sum int
	for _, n := range c.results {
		if c.winMap[n] {
			sum++
		}
	}
	return sum
}

func (c card) points() int {
	var sum int
	for _, n := range c.results {
		if c.winMap[n] {
			if sum == 0 {
				sum = 1
			} else {
				sum *= 2
			}
		}
	}
	return sum
}

func TotalScratchCards() (int, error) {
	cards, err := parseData()
	if err != nil {
		return 0, err
	}
	acc := make([]int, len(cards))
	for i, c := range cards {
		acc[i]++
		m := c.numberMatches() + 1
		if i == len(cards)-1 || m == 1 {
			continue
		}
		if i+m >= len(cards) {
			m = len(cards) - i
		}
		copied := cards[i+1 : i+m]
		for _, cc := range copied {
			acc[cc.number-1] += 1 * acc[i]
		}
	}
	var sum int
	for _, n := range acc {
		sum += n
	}
	return sum, nil
}

func TotalPoints() (int, error) {
	cards, err := parseData()
	if err != nil {
		return 0, err
	}
	var sum int
	for _, c := range cards {
		sum += c.points()
	}
	return sum, nil
}

func parseData() ([]card, error) {
	scanner, err := data.GetScanner(day)
	if err != nil {
		return nil, fmt.Errorf("getting scanner: %w", err)
	}
	var cards []card
	for scanner.Scan() {
		l := scanner.Text()
		var c card
		c, err = parseLine(l)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, nil
}

func parseLine(l string) (card, error) {
	const wantSplitLen = 2
	cardResult := cardRegexp.FindStringSubmatch(l)
	if len(cardResult) != wantSplitLen {
		return card{}, fmt.Errorf("invalid card expression %q", l)
	}
	number, err := strconv.Atoi(cardResult[1])
	if err != nil {
		return card{}, fmt.Errorf("parsing card number: %w", err)
	}
	lSplit := strings.Split(l, ":")
	if len(lSplit) != wantSplitLen {
		return card{}, fmt.Errorf("invalid card expression %q", l)
	}
	numbersSplit := strings.Split(lSplit[1], "|")
	results, err := extractNumbers(numbersSplit[0])
	if err != nil {
		return card{}, fmt.Errorf("parsing results: %w", err)
	}
	winningNumbers, err := extractNumbers(numbersSplit[1])
	if err != nil {
		return card{}, fmt.Errorf("parsing winning numbers: %w", err)
	}
	winMap := make(map[int]bool, len(winningNumbers))
	for _, wn := range winningNumbers {
		winMap[wn] = true
	}
	return card{
		number:  number,
		results: results,
		winMap:  winMap,
	}, nil
}

func extractNumbers(s string) ([]int, error) {
	split := strings.Split(s, " ")
	ns := make([]int, 0, len(split))
	for _, nstr := range split {
		if nstr == "" {
			continue
		}
		n, err := strconv.Atoi(nstr)
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}
