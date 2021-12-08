package main

import (
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// var mainWG sync.WaitGroup
	fileMap := initFileMap()

	dayToRun := os.Args[1]
	aocSwitch(dayToRun, fileMap)
}
