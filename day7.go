package main

import (
	"log"
	"math"
	"sync"
)

type DistMatrix struct {
	mu     sync.Mutex
	Matrix *[][]float64
}

type ShareableArray struct {
	mu    sync.Mutex
	Array *[]float64
}

func day7(filepath string) {
	crabs, err := consumeIntFile(filepath, ",")
	if err != nil {
		log.Println("Error encountered consuming file", err)
		return
	}
	crabsThatFloat := make([]float64, len(crabs))
	distMatrix := make([][]float64, len(crabs))
	for i, val := range crabs {
		distMatrix[i] = make([]float64, len(crabs))
		crabsThatFloat[i] = float64(val)
	}

	log.Println(crabsThatFloat[0])
	crabArray := ShareableArray{Array: &crabsThatFloat}
	crabDistances := DistMatrix{Matrix: &distMatrix}

	for routine := 0; routine < len(crabs); routine++ {
		wg.Add(1)
		go distMatrixRow(&crabDistances, &crabArray, routine, len(crabs))
	}

	wg.Wait()

	min := math.Inf(1)
	minIndex := -1

	for index, row := range *crabDistances.Matrix {
		rowSum := floatSum(row...)
		if rowSum < min {
			min = rowSum
			minIndex = index
		}
	}

	log.Println("Minimum Distance of", min, "found at index", minIndex)
}

func floatSum(input ...float64) float64 {
	var sumValue float64 = 0
	for _, i := range input {
		sumValue += i
	}
	return sumValue
}

func distMatrixRow(matrix *DistMatrix, array *ShareableArray, rowToCompute int, arrayLength int) {
	defer wg.Done()
	array.mu.Lock()
	fromValue := (*array.Array)[rowToCompute]
	arrayCopy := (*array.Array)[:]
	array.mu.Unlock()

	matrixRow := make([]float64, arrayLength)

	for i := 0; i < arrayLength; i++ {
		matrixRow[i] = dist(fromValue, arrayCopy[i])
	}

	matrix.mu.Lock()
	(*matrix.Matrix)[rowToCompute] = matrixRow
	matrix.mu.Unlock()
	return
}

func dist(a float64, b float64) float64 {
	// return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	// return math.Abs(a - b) // This is the distance for Part 1
	return 0.5 * (math.Abs(a - b)) * (math.Abs(a-b) + 1) // This is the distance for Part 2
}
