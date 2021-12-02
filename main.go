package main

import "os"

func main() {
	fileMap := initFileMap()

	dayToRun := os.Args[1]

	aocSwitch(dayToRun, fileMap)
}
