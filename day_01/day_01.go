package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInput() []int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	measurements := make([]int, 0)
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Failed to convert measurements to int: %v\n", err)
		}
		measurements = append(measurements, measurement)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return measurements
}

type MeasurementCalculator func(measurements []int, index int) int

func CountIncreases(measurements []int, upperLimit int, calculator MeasurementCalculator) int {
	prev, increaseCounter := calculator(measurements, 0), 0
	for i := range measurements[1 : upperLimit+1] {
		v := calculator(measurements, i)
		if v > prev {
			increaseCounter++
		}
		prev = v
	}
	return increaseCounter
}

func SingleMeasurementCalculator(measurements []int, index int) int {
	return measurements[index]
}

func WindowedMeasurementCalculator(measurements []int, index int) int {
	return measurements[index] + measurements[index+1] + measurements[index+2]
}

func main() {
	measurements := ReadInput()
	log.Printf("First star: %v\n", CountIncreases(measurements, len(measurements), SingleMeasurementCalculator))
	log.Printf("Second star: %v\n", CountIncreases(measurements, len(measurements)-2, WindowedMeasurementCalculator))
}
