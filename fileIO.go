package main

import (
	"fmt"
	"io/ioutil"
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
