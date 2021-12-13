package main

import (
	"log"
	"strings"
)

func day8(filepath string) {
	diagnosticsList, err := consumeFile(filepath)
	if err != nil {
		log.Println("Error consuming file!")
		return
	}
	uniqueDigits := 0
	seenMap := make(map[string]int)

	for _, row := range diagnosticsList {
		output := strings.Split(row, " | ")
		if len(output) > 1 {
			digits := strings.Split(output[1], " ")
			for _, command := range digits {
				if command != "|" {
					if (len(command) >= 2 && len(command) <= 4) || (len(command) == 7) {
						uniqueDigits++
						seenMap[command]++
					}
				}
			}
		}
	}
	log.Println(uniqueDigits)
}
