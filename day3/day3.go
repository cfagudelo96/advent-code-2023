package day3

import (
	"fmt"
	"strconv"

	"cfagudelo/advent-code-2023/data"
)

const (
	irrelevant = iota
	symbol
	number
	gear
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

func SumRelevantParts() (int, error) {
	m, rm, ns, err := parseData()
	if err != nil {
		return 0, fmt.Errorf("parsing data: %w", err)
	}
	return sumRelevantNumbers(m, rm, ns), nil
}

func sumRelevantNumbers(m [][]int, rm [][]rune, ns []inputNumber) int {
	var sum int
	for _, n := range ns {
		prevSymbol := n.start > 0 && m[n.row][n.start-1] == symbol
		nextSymbol := n.end < len(m[n.row])-1 && m[n.row][n.end+1] == symbol
		prevRowSymbol := n.row > 0 && checkSymbolInRange(n, n.row-1, m)
		nextRowSymbol := n.row < len(m)-1 && checkSymbolInRange(n, n.row+1, m)
		if prevSymbol || nextSymbol || prevRowSymbol || nextRowSymbol {
			sum += n.value
			// printDebug(n, m, rm)
		} else {
			// printDebug(n, m, rm)
		}
	}
	return sum
}

func printDebug(n inputNumber, m [][]int, rm [][]rune) {
	fmt.Printf("number: %#v\n", n)
	start := n.start
	end := n.end
	if n.start > 0 {
		start--
	}
	if n.end < len(m[n.row]) {
		end++
	}
	if n.row > 0 {
		printRunes(rm[n.row-1][start : end+1])
	}
	printRunes(rm[n.row][start : end+1])
	if n.row < len(m)-1 {
		printRunes(rm[n.row+1][start : end+1])
	}
}

func printRunes(rs []rune) {
	strs := make([]string, len(rs))
	for i, r := range rs {
		strs[i] = string(r)
	}
	fmt.Println(strs)
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
