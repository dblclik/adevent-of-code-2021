package main

import (
	"log"
)

/* Day 1 asks us to solve the complement sum problem
   - Part 1: Find two values that sum to 2020, return their product
*/
func day1(filepath string) {
	log.Println("Ingesting day 1 file...")
	day1Data, err := consumeIntFile(filepath)
	if err != nil {
		log.Println("Error consuming file!")
		return
	}
	log.Println("Beginning Day 1, Part 1...")
	day1part1 := greaterThanPrev(day1Data)
	log.Println("The answer to Day 1, Part 1 is: ", day1part1)

	slidingWindow := slidingWindowSum(day1Data, 3)
	log.Println("Beginning Day 1, Part 2...")
	day1part2 := greaterThanPrev(slidingWindow)
	log.Println("The answer to Day 1, Part 2 is: ", day1part2)
}

func greaterThanPrev(inputArray []int) int {
	result := 0
	if len(inputArray) <= 1 {
		return result
	}

	for ind := 1; ind < len(inputArray); ind++ {
		if inputArray[ind] > inputArray[ind-1] {
			result++
		}
	}
	return result
}

func slidingWindowSum(inputArray []int, window int) []int {
	windowedArray := make([]int, len(inputArray)-window+1)
	for ind := 0; ind < len(windowedArray); ind++ {
		windowedArray[ind] = sum(inputArray[ind : ind+window]...)
	}
	return windowedArray
}

func sum(input ...int) int {
	sumValue := 0
	for _, i := range input {
		sumValue += i
	}
	return sumValue
}
