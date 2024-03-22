package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"cfagudelo/advent-code-2023/data"
)

const (
	day2                   = 2
	expectedSubmatchLength = 3
)

const (
	nRed   = 12
	nGreen = 13
	nBlue  = 14
)

var (
	gameRegexp  = regexp.MustCompile(`^Game (\d+): (.*)$`)
	valueRegexp = regexp.MustCompile(`^(\d+) (red|green|blue)$`)
)

type game struct {
	id   int
	sets []set
}

type set struct {
	red   int
	green int
	blue  int
}

func (g game) isPossible(maxRed, maxGreen, maxBlue int) bool {
	for _, s := range g.sets {
		if s.red > maxRed || s.green > maxGreen || s.blue > maxBlue {
			return false
		}
	}
	return true
}

func (g game) fewestSet() set {
	var maxRed, maxGreen, maxBlue int
	for _, s := range g.sets {
		if s.red > maxRed {
			maxRed = s.red
		}
		if s.green > maxGreen {
			maxGreen = s.green
		}
		if s.blue > maxBlue {
			maxBlue = s.blue
		}
	}
	return set{
		red:   maxRed,
		green: maxGreen,
		blue:  maxBlue,
	}
}

func (s set) power() int {
	return s.red * s.green * s.blue
}

func SumOfPossibleGamesIDs() (int, error) {
	scanner, err := data.GetScanner(day2)
	if err != nil {
		return 0, fmt.Errorf("getting scanner: %w", err)
	}
	var sum int
	for scanner.Scan() {
		var g game
		g, err = parseGame(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("parsing game: %w", err)
		}
		if g.isPossible(nRed, nGreen, nBlue) {
			sum += g.id
		}
	}
	return sum, nil
}

func SumOfFewestSet() (int, error) {
	scanner, err := data.GetScanner(day2)
	if err != nil {
		return 0, fmt.Errorf("getting scanner: %w", err)
	}
	var sum int
	for scanner.Scan() {
		var g game
		g, err = parseGame(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("parsing game: %w", err)
		}
		sum += g.fewestSet().power()
	}
	return sum, nil
}

func parseGame(l string) (game, error) {
	m := gameRegexp.FindStringSubmatch(l)
	if len(m) != expectedSubmatchLength {
		return game{}, errors.New("invalid regex match")
	}
	id, err := strconv.Atoi(m[1])
	if err != nil {
		return game{}, fmt.Errorf("parsing id: %w", err)
	}
	setsStr := strings.Split(m[2], "; ")
	sets := make([]set, len(setsStr))
	for i, v := range setsStr {
		var s set
		s, err = parseSet(v)
		if err != nil {
			return game{}, fmt.Errorf("parsing set %s: %w", v, err)
		}
		sets[i] = s
	}
	return game{
		id:   id,
		sets: sets,
	}, nil
}

func parseSet(l string) (set, error) {
	values := strings.Split(l, ", ")
	var s set
	for _, v := range values {
		m := valueRegexp.FindStringSubmatch(v)
		if len(m) != expectedSubmatchLength {
			return s, errors.New("invalid regex match")
		}
		n, err := strconv.Atoi(m[1])
		if err != nil {
			return s, fmt.Errorf("parsing value %s: %w", v, err)
		}
		switch m[2] {
		case "red":
			s.red = n
		case "green":
			s.green = n
		case "blue":
			s.blue = n
		}
	}
	return s, nil
}
