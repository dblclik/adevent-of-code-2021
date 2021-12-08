package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func day5(filepath string) {
	day5input, err := consumeFile(filepath)
	if err != nil {
		log.Println("Error consuming file!")
		return
	}
	maxX := 0
	maxY := 0
	lineArray := make([][]Point, len(day5input))

	for ind, line := range day5input {
		if len(line) > 0 {
			points := strings.Split(line, " -> ")
			point0 := strings.Split(points[0], ",")
			point1 := strings.Split(points[1], ",")
			point0x, err := strconv.Atoi(point0[0])
			if err != nil {
				log.Println("Could not parse magnitude: ", point0[0])
			}
			point0y, err := strconv.Atoi(point0[1])
			if err != nil {
				log.Println("Could not parse magnitude: ", point0[1])
			}

			point1x, err := strconv.Atoi(point1[0])
			if err != nil {
				log.Println("Could not parse magnitude: ", point1[0])
			}
			point1y, err := strconv.Atoi(point1[1])
			if err != nil {
				log.Println("Could not parse magnitude: ", point1[1])
			}

			if point0x > maxX {
				maxX = point0x
			}
			if point1x > maxX {
				maxX = point1x
			}

			if point0y > maxY {
				maxY = point0y
			}
			if point1y > maxY {
				maxY = point1y
			}

			lineArray[ind] = make([]Point, 2)
			lineArray[ind][0] = Point{x: point0x, y: point0y}
			lineArray[ind][1] = Point{x: point1x, y: point1y}
		}
	}
	log.Println("max-X:", maxX)
	log.Println("max-Y:", maxY)

	// initialize the grid
	floorGrid := make([][]int, maxY+1)
	for row := 0; row < len(floorGrid); row++ {
		floorGrid[row] = make([]int, maxX+1)
	}

	log.Println("Made the grid with size:", len(floorGrid), "x", len(floorGrid[0]))

	overlaps := 0
	for ind, points := range lineArray {
		// if x vals are equal, iterate over y values
		if len(points) == 0 {
			log.Println("Encountered 0 length lineArray at row", ind)
			break
		}
		if points[0].x == points[1].x {
			minYpoint := int(math.Min(float64(points[0].y), float64(points[1].y)))
			maxYpoint := int(math.Max(float64(points[0].y), float64(points[1].y)))
			for y := minYpoint; y <= maxYpoint; y++ {
				floorGrid[y][points[0].x]++
			}
		} else if points[0].y == points[1].y {
			minXpoint := int(math.Min(float64(points[0].x), float64(points[1].x)))
			maxXpoint := int(math.Max(float64(points[0].x), float64(points[1].x)))
			for x := minXpoint; x <= maxXpoint; x++ {
				floorGrid[points[0].y][x]++
			}
		} else {
			// diagonal lines can go up and to right, or down and to right (from ref of min-x point)
			var startPoint, endPoint Point
			var incrementer int
			if points[0].x < points[1].x {
				startPoint = points[0]
				endPoint = points[1]
			} else {
				startPoint = points[1]
				endPoint = points[0]
			}

			// check if end point is up or down rel to startPoint
			if endPoint.y > startPoint.y {
				incrementer = 1
			} else {
				incrementer = -1
			}

			nextPoint := Point{x: startPoint.x, y: startPoint.y}
			log.Println("Start Point:", startPoint, "End Point:", endPoint, "Next Point:", nextPoint)
			for {
				floorGrid[nextPoint.x][nextPoint.y]++
				if (nextPoint.x == endPoint.x) && (nextPoint.y == endPoint.y) {
					break
				}
				nextPoint.x++
				nextPoint.y += incrementer
				log.Println("Start Point:", startPoint, "End Point:", endPoint, "Next Point:", nextPoint)
			}
		}

	}

	for row := 0; row < len(floorGrid); row++ {
		for col := 0; col < len(floorGrid[row]); col++ {
			if floorGrid[row][col] > 1 {
				overlaps++
			}
		}
	}
	log.Println("The number of overlapping points is:", overlaps)
}
