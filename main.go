package main

import (
	"log"

	"cfagudelo/advent-code-2023/day4"
)

func main() {
	s, err := day4.TotalScratchCards()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Default().Printf("%d", s)
}
