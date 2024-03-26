package main

import (
	"log"

	"cfagudelo/advent-code-2023/day3"
)

func main() {
	s, err := day3.SumGearRatios()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Default().Printf("%d", s)
}
