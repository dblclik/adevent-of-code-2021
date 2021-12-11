package main

import (
	"log"
	"strconv"
	"strings"
)

func day6(filepath string) {
	startingSchool, err := consumeFile(filepath)
	if err != nil {
		log.Println("Error encountered consuming file", err)
		return
	}

	schoolString := strings.Split(startingSchool[0], ",")
	lanternSchool := map[int]map[int]int{
		0: {
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
			8: 0,
		},
	}

	for _, str := range schoolString {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Println("Error converting string to int:", str)
			return
		}
		lanternSchool[0][val]++
	}

	nextDayMap := map[int]int{
		0: 6,
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
		6: 5,
		7: 6,
		8: 7,
	}

	for day := 1; day <= 256; day++ {
		lanternSchool[day] = newMap()
		for popHash := 0; popHash < 9; popHash++ {
			lanternSchool[day][nextDayMap[popHash]] += lanternSchool[day-1][popHash]
			if popHash == 0 {
				lanternSchool[day][8] += lanternSchool[day-1][popHash]
			}
		}
		log.Println("After", day, "days, school size is:", mapSum(lanternSchool[day]))
	}

}

func mapSum(input map[int]int) int {
	currSum := 0
	for k, _ := range input {
		currSum += input[k]
	}

	return currSum
}

func newMap() map[int]int {
	return map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}
}
