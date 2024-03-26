package main

import (
	"log"

	"cfagudelo/advent-code-2023/day3"
)

func main() {
	s, err := day3.SumRelevantParts()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Default().Printf("%d", s)
}
