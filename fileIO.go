package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func consumeIntFile(filepath string) ([]int, error) {
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, fmt.Errorf("FileParseError: Could not parse file: %v", filepath)
	}

	stringData := strings.Split(string(fileData), "\n")
	returnData := make([]int, len(stringData))

	for i, val := range stringData {
		returnData[i], err = strconv.Atoi(val)
		if err != nil {
			fmt.Println("Data parsing error", err)
			if val != "" {
				return nil, fmt.Errorf("FileParseError: Could not convert value to int: %v", val)
			}
		}
	}

	return returnData, nil
}

func consumeFile(filepath string) ([]string, error) {
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, fmt.Errorf("FileParseError: Could not parse file: %v", filepath)
	}

	stringData := strings.Split(string(fileData), "\n")

	return stringData, nil
}

func readBingoFile(filepath string) ([]int, [][][]int, error) {
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, nil, fmt.Errorf("FileParseError: Could not parse file: %v", filepath)
	}
	stringData := strings.Split(string(fileData), "\n")
	gameCallouts := strings.Split(stringData[0], ",")
	gameCalloutsInt := make([]int, len(gameCallouts))
	for ind, value := range gameCallouts {
		callout, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Game Callout Conversion Error", err)
			return nil, nil, fmt.Errorf("FileParseError: Could not parse file: %v", filepath)
		}
		gameCalloutsInt[ind] = callout
	}

	log.Println(int((len(stringData) - 2)) / 6)
	gameBoards := make([][][]int, int((len(stringData)-2))/6)
	log.Println(len(gameBoards))

	board := make([][]int, 5)
	boardIndex := 0
	currRow := 0
	for currIndex := 2; currIndex < len(stringData); currIndex++ {
		if stringData[currIndex] == "" {
			gameBoards[boardIndex] = board
			boardIndex++
			board = make([][]int, 5)
			currRow = 0
		} else {
			currBoardRow := strings.Split(stringData[currIndex], " ")
			board[currRow] = make([]int, 5)
			colIndex := 0
			for _, rowValue := range currBoardRow {
				if rowValue != "" {
					boardRowColValue, err := strconv.Atoi(rowValue)
					if err != nil {
						fmt.Println("Game Board Conversion Error", err)
						return nil, nil, fmt.Errorf("FileParseError: Could not parse file: %v", filepath)
					}
					board[currRow][colIndex] = boardRowColValue
					colIndex++
				}
			}
			currRow++
		}
	}

	return gameCalloutsInt, gameBoards, nil
}
