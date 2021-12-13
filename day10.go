package main

import (
	"log"
	"sort"
)

func day10(filepath string) {
	navigationSubsystem, err := consumeFile(filepath)
	if err != nil {
		log.Println("Error consuming file!")
		return
	}

	syntaxScore := 0
	scoringMap := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	completeScoreMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	pairMap := map[int]int{
		40:  41,  // (: )
		41:  40,  // ): (
		60:  62,  // <: >
		62:  60,  // >: <
		91:  93,  // [: ]
		93:  91,  // ]: [
		123: 125, // {: }
		125: 123, // }: {
	}

	autoCompleteArray := []int{}

	for _, navLine := range navigationSubsystem {
		if len(navLine) > 0 {
			var navStack Stack
			for i, char := range navLine {
				// Check to see if we have a closing character
				if score, ok := scoringMap[string(char)]; ok {
					lastOpen, found := navStack.Pop()
					if found {
						lastOpenRune := []rune(lastOpen)[0]
						if pairMap[int(lastOpenRune)] != int(char) {
							log.Println("Expected", string(pairMap[int(lastOpenRune)]), "but found", string(char), "instead.")
							syntaxScore += score
							break
						}
					}
				} else { // we have an opening character
					navStack.Push(string(char))
				}
				// if we reach the end of the line and haven't break, de-stack and score
				if i == (len(navLine) - 1) {
					autocompleteScore := 0
					for {
						lastOpen, found := navStack.Pop()
						if !found {
							break
						}
						lastOpenRune := []rune(lastOpen)[0]
						autocompleteScore *= 5
						autocompleteScore += completeScoreMap[string(pairMap[int(lastOpenRune)])]
					}
					autoCompleteArray = append(autoCompleteArray, autocompleteScore)
				}
			}
		}
	}
	sort.Ints(autoCompleteArray)

	log.Println(syntaxScore)
	log.Println(autoCompleteArray[int(len(autoCompleteArray)/2)])

}
