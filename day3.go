package main

import (
	"log"
	"strconv"
)

func day3(filepath string) {
	day3input, err := consumeFile(filepath)
	if err != nil {
		log.Println("Error consuming file!")
		return
	}
	readingLength := len(day3input[0])
	bitArray := make([][2]int, readingLength)
	log.Println(bitArray)
	for _, reading := range day3input {
		if len(reading) > 0 {
			readingRunes := []rune(reading)
			for ind, val := range readingRunes {
				bitValue, err := strconv.Atoi(string(val))
				if err != nil {
					log.Fatalln("Could not convert stringified rune to int")
				}
				bitArray[ind][bitValue]++
			}
		}
	}
	log.Println(bitArray)
	gamma := ""
	epsilon := ""
	for _, bitCount := range bitArray {
		if bitCount[0] > bitCount[1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Println("Error converting gamma to int")
	}
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Println("Error converting epsilon to int")
	}
	log.Println("Gamma reading is: ", gamma, "(int value of ", gammaInt)
	log.Println("Epsilon reading is: ", epsilon, "(int value of ", epsilonInt)

	log.Println("")
	log.Println("Beginning Day 3, part 2...")
	log.Println("Day 3, part 2 -- O2 Filtering...")
	inputCopy := make([]string, len(day3input))
	copy(inputCopy, day3input)
	var o2readings []string
	o2Key := ""
	for i := 0; i < readingLength; i++ {
		bitCountAtIndex := bitCount(inputCopy, i)
		if bitCountAtIndex[0] > bitCountAtIndex[1] {
			o2Key += "0"
		} else {
			o2Key += "1"
		}
		inputCopy = listFilter(inputCopy, o2Key)
		if len(inputCopy) == 1 {
			o2readings = inputCopy
			break
		}
	}
	log.Println(o2readings)

	log.Println("Day 3, part 2 -- CO2 Filtering...")
	inputCopy = make([]string, len(day3input))
	copy(inputCopy, day3input)
	log.Println(len(inputCopy))
	var co2readings []string
	co2Key := ""
	for i := 0; i < readingLength; i++ {
		bitCountAtIndex := bitCount(inputCopy, i)
		log.Println(bitCountAtIndex)
		if bitCountAtIndex[0] <= bitCountAtIndex[1] {
			co2Key += "0"
		} else {
			co2Key += "1"
		}
		inputCopy = listFilter(inputCopy, co2Key)
		if len(inputCopy) == 1 {
			co2readings = inputCopy
			break
		}
	}

	log.Println(co2readings)

	o2Int, err := strconv.ParseInt(o2readings[0], 2, 64)
	if err != nil {
		log.Println("Error converting o2reading to int")
	}
	co2Int, err := strconv.ParseInt(co2readings[0], 2, 64)
	if err != nil {
		log.Println("Error converting epsilon to int")
	}
	log.Println("O2 reading is: ", o2readings[0], "(int value of ", o2Int)
	log.Println("C02 reading is: ", co2readings[0], "(int value of ", co2Int)

}

func bitCount(inputArray []string, index int) []int {
	bits := make([]int, 2)
	for _, reading := range inputArray {
		if len(reading) > 0 {
			readingRunes := []rune(reading)
			bitValue, err := strconv.Atoi(string(readingRunes[index]))
			if err != nil {
				log.Println("Cannt convert bit to int", readingRunes[index])
			}
			bits[bitValue]++
		}
	}
	return bits
}

// This is simple (see: naive) and brute force... :(
func listFilter(inputArray []string, key string) []string {
	keyRune := []rune(key)
	// log.Println(keyRune)
	var mutableArray []string
	for ind, currentRune := range keyRune {
		mutableArray = []string{}
		for _, reading := range inputArray {
			if len(reading) > 0 {
				readingRune := []rune(reading)
				if readingRune[ind] == currentRune {
					mutableArray = append(mutableArray, reading)
				}
			}
		}
		// log.Println(len(mutableArray))
		if (len(mutableArray) == 1) || (ind == len(keyRune)-1) {
			break
		} else {
			inputArray = mutableArray
		}
	}

	return mutableArray
}
