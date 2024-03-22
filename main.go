package main

import (
	"log"
)

func main() {
	s, err := SumOfFewestSet()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Default().Printf("%d", s)
}
