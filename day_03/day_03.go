package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInput() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	commands := make([]string, 0)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return commands
}

func CountOnes(numbers []string) []int {
	bitCount := len(numbers[0])
	counts := make([]int, bitCount)
	for i := 0; i < bitCount; i++ {
		counts[i] = CountOnesAt(numbers, i)
	}
	return counts
}

func CountOnesAt(numbers []string, index int) int {
	var count int
	for _, v := range numbers {
		bit := v[index]
		if bit == '1' {
			count++
		}
	}
	return count
}

type BitCriteria func(number string) bool
type BitCriteriaGenerator func(remainingNumbers []string, index int) BitCriteria

func OxygenGeneratorBitCriteria(remainingNumbers []string, index int) BitCriteria {
	isOneMostCommon := float64(CountOnesAt(remainingNumbers, index)) >= float64(len(remainingNumbers))/2
	return func(number string) bool {
		return (isOneMostCommon && number[index] == '1') || (!isOneMostCommon && number[index] == '0')
	}
}

func CO2ScrubberBitCriteria(remainingNumbers []string, index int) BitCriteria {
	isZeroMostCommon := float64(len(remainingNumbers)-CountOnesAt(remainingNumbers, index)) > float64(len(remainingNumbers))/2.0
	return func(number string) bool {
		return (!isZeroMostCommon && number[index] == '0') || (isZeroMostCommon && number[index] == '1')
	}
}

func FilterValues(totalValues []string, bitCriteriaGenerator BitCriteriaGenerator) string {
	candidates := totalValues
	for i := 0; len(candidates) > 1; i++ {
		remainingInputs := make([]string, 0, len(candidates))
		bitCriteria := bitCriteriaGenerator(candidates, i)
		for _, v := range candidates {
			if bitCriteria(v) {
				remainingInputs = append(remainingInputs, v)
			}
		}
		candidates = remainingInputs
	}
	return candidates[0]
}

func main() {
	input := ReadInput()
	counts := CountOnes(input)
	var gamma, epsilon int
	for _, v := range counts {
		if v > len(input)/2 {
			gamma = gamma<<1 + 1
			epsilon <<= 1
		} else {
			gamma <<= 1
			epsilon = epsilon<<1 + 1
		}
	}
	log.Printf("First star: %v\n", gamma*epsilon)

	oxygenRatingString := FilterValues(input, OxygenGeneratorBitCriteria)
	co2RatingString := FilterValues(input, CO2ScrubberBitCriteria)

	oxygenRating, _ := strconv.ParseInt(oxygenRatingString, 2, 0)
	co2Rating, _ := strconv.ParseInt(co2RatingString, 2, 0)

	log.Printf("Second star: %v\n", oxygenRating*co2Rating)
}
