package main

import (
	"log"
)

func main() {
	s, err := SumCalibrationValuesV2()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Default().Printf("%d", s)
}
