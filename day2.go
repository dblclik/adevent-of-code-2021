package main

import (
	"log"
	"strconv"
	"strings"
)

func day2(filepath string) {
	day2commands, err := consumeFile(filepath)
	if err != nil {
		log.Println("Error consuming file!")
		return
	}

	log.Println("Beginning Day 2, Part 1")
	horizontal, depth, err := vec2movement(day2commands, 0, 0, false)
	log.Println("The new position after the commands is: ", horizontal, depth)

	log.Println("Beginning Day 2, Part 2")
	horizontal, depth, err = vec2movement(day2commands, 0, 0, true)
	log.Println("The new position after the commands is: ", horizontal, depth)
}

// vec2movement allows for a line of commands to be translated into a new (x,y) position
func vec2movement(commands []string, horiz int, vert int, useAim bool) (int, int, error) {
	movementMap := map[string][2]int{
		"up":       {0, 1},
		"down":     {0, -1},
		"forward":  {1, 0},
		"backward": {-1, 0},
	}

	aim := 0

	for _, command := range commands {
		commandComponents := strings.Split(command, " ")
		if len(commandComponents) < 2 {
			log.Println("Encountered a line with less commands than expected... ", command)
		} else {
			magnitude, err := strconv.Atoi(commandComponents[1])
			if err != nil {
				log.Println("Could not parse magnitude: ", commandComponents)
			}
			movement := movementMap[commandComponents[0]]
			horiz += movement[0] * magnitude
			if !useAim {
				vert += movement[1] * magnitude
			} else {
				aim += movement[1] * magnitude
				vert += aim * movement[0] * magnitude
			}
		}
	}
	return horiz, vert, nil
}
