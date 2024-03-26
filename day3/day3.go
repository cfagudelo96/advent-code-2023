package day3

import (
	"fmt"
	"strconv"

	"cfagudelo/advent-code-2023/data"
)

const (
	irrelevant = iota
	symbol
)

const (
	day     = 3
	decimal = 10
)

type inputNumber struct {
	value int
	row   int
	start int
	end   int
}

type gear struct {
	row      int
	position int
}

func (g gear) ratio(maps []map[int]*inputNumber) int {
	var n1, n2 *inputNumber
	start := g.position - 1
	end := g.position + 1
	for i := g.row - 1; i <= g.row+1; i++ {
		if g.row < 0 || g.row == len(maps) {
			continue
		}
		for j := start; j <= end; j++ {
			n := maps[i][j]
			if n == nil {
				continue
			}
			switch {
			case n1 == nil:
				n1 = n
			case n2 == nil && n != n1:
				n2 = n
			case n != n1 && n2 != nil && n != n2:
				return 0
			}
		}
	}
	if n1 != nil && n2 != nil {
		return n1.value * n2.value
	}
	return 0
}

func SumRelevantParts() (int, error) {
	m, _, ns, err := parseData()
	if err != nil {
		return 0, fmt.Errorf("parsing data: %w", err)
	}
	return sumRelevantNumbers(m, ns), nil
}

func SumGearRatios() (int, error) {
	gears, numbersMap, err := parseDataGears()
	if err != nil {
		return 0, fmt.Errorf("parsing data: %w", err)
	}
	var sum int
	for _, g := range gears {
		sum += g.ratio(numbersMap)
	}
	return sum, nil
}

func sumRelevantNumbers(m [][]int, ns []inputNumber) int {
	var sum int
	for _, n := range ns {
		prevSymbol := n.start > 0 && m[n.row][n.start-1] == symbol
		nextSymbol := n.end < len(m[n.row])-1 && m[n.row][n.end+1] == symbol
		prevRowSymbol := n.row > 0 && checkSymbolInRange(n, n.row-1, m)
		nextRowSymbol := n.row < len(m)-1 && checkSymbolInRange(n, n.row+1, m)
		if prevSymbol || nextSymbol || prevRowSymbol || nextRowSymbol {
			sum += n.value
		}
	}
	return sum
}

func checkSymbolInRange(n inputNumber, row int, m [][]int) bool {
	start := n.start
	end := n.end
	if start > 0 {
		start--
	}
	if end < len(m[row])-1 {
		end++
	}
	for i := start; i <= end; i++ {
		if m[row][i] == symbol {
			return true
		}
	}
	return false
}

func parseDataGears() ([]gear, []map[int]*inputNumber, error) {
	scanner, err := data.GetScanner(day)
	if err != nil {
		return nil, nil, fmt.Errorf("getting scanner: %w", err)
	}
	var (
		gears            []gear
		inputNumbersMaps []map[int]*inputNumber
		i                int
	)
	for scanner.Scan() {
		l := scanner.Text()
		newGears, newMap, pErr := parseLineGears(l, i)
		if pErr != nil {
			return nil, nil, fmt.Errorf("parsing line: %w", pErr)
		}
		gears = append(gears, newGears...)
		inputNumbersMaps = append(inputNumbersMaps, newMap)
		i++
	}
	return gears, inputNumbersMaps, nil
}

func parseLineGears(l string, row int) ([]gear, map[int]*inputNumber, error) {
	var (
		gears []gear
		n     *inputNumber
	)
	nsMap := make(map[int]*inputNumber)
	ns := make([]int, len(l))
	for i, c := range l {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, nil, fmt.Errorf("parsing character %v: %w", c, err)
			}
			if n == nil {
				n = &inputNumber{
					value: v,
					row:   row,
					start: i,
					end:   i,
				}
			} else {
				n.value = n.value*decimal + v
				n.end = i
			}
			nsMap[i] = n
		default:
			if c == '*' {
				gears = append(gears, gear{
					row:      row,
					position: i,
				})
			} else if c != '.' {
				ns[i] = symbol
			}
			if n != nil {
				n = nil
			}
		}
	}
	return gears, nsMap, nil
}

func parseData() ([][]int, [][]rune, []inputNumber, error) {
	scanner, err := data.GetScanner(day)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("getting scanner: %w", err)
	}
	var matrix [][]int
	var runeMatrix [][]rune
	var inputNumbers []inputNumber
	var i int
	for scanner.Scan() {
		l := scanner.Text()
		var runes []rune
		for _, r := range l {
			runes = append(runes, r)
		}
		runeMatrix = append(runeMatrix, runes)
		lns, nins, pErr := parseLine(l, i)
		if pErr != nil {
			return nil, nil, nil, fmt.Errorf("parsing line: %w", pErr)
		}
		matrix = append(matrix, lns)
		inputNumbers = append(inputNumbers, nins...)
		i++
	}
	return matrix, runeMatrix, inputNumbers, nil
}

func parseLine(l string, row int) ([]int, []inputNumber, error) {
	var inputNumbers []inputNumber
	var n *inputNumber
	ns := make([]int, len(l))
	for i, c := range l {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, nil, fmt.Errorf("parsing character %v: %w", c, err)
			}
			if n == nil {
				n = &inputNumber{
					value: v,
					row:   row,
					start: i,
					end:   i,
				}
			} else {
				n.value = n.value*decimal + v
				n.end = i
			}
		default:
			if c != '.' {
				ns[i] = symbol
			}
			if n != nil {
				inputNumbers = append(inputNumbers, *n)
				n = nil
			}
		}
	}
	if n != nil {
		inputNumbers = append(inputNumbers, *n)
	}
	return ns, inputNumbers, nil
}
