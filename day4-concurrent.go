package main

import (
	"log"
	"sort"
	"sync"
	"time"
)

type PlayerResponse struct {
	ID         int
	LastPlayed int
	LastIndex  int
	BoardSum   int
}

type ResponseArray struct {
	Responses []PlayerResponse
	mu        sync.Mutex
}

type PlayerCard struct {
	ID      int
	Enabled bool
	Channel chan []int
	Board   [][]int
	// ResponseChannel chan PlayerResponse
	mu sync.Mutex
}

func day4(filepath string) {
	calls, boards, err := readBingoFile(filepath)
	if err != nil {
		log.Println("Could not ingest bingo file")
	}
	log.Println(len(calls), len(boards))
	log.Println(boards[0])
	log.Println("")
	log.Println(boards[len(boards)-1])
	// var wg sync.WaitGroup

	// responseChannel := make(chan PlayerResponse, len(boards))

	// Create the slice of channels to send calls to the boards
	playerArray := make([]PlayerCard, len(boards))
	for i := range playerArray {
		playerArray[i].ID = i
		playerArray[i].Enabled = true
		playerArray[i].Channel = make(chan []int, len(calls))
		playerArray[i].Board = boards[i]
		// playerArray[i].ResponseChannel = responseChannel
		// wg.Add(1)
		// go bingoPlayer(&playerArray[i], &wg)
	}

	// load the channels up with the calls
	for index := 0; index < len(calls); index++ {
		call := calls[index]
		for i := range playerArray {
			if playerArray[i].Enabled {
				playerArray[i].Channel <- []int{call, index}
			}
		}
	}

	var responses ResponseArray
	log.Println("Number of responses:", len(responses.Responses))

	// close the channels since all sends have happened
	for i := range playerArray {
		defer close(playerArray[i].Channel)
	}

	for i := range playerArray {
		wg.Add(1)
		go bingoPlayer(&playerArray[i], &wg, &responses)
	}

	time.Sleep(5 * time.Second)

	// wg.Add(1)
	// go monitorResponseChannel(responseChannel, &playerArray, &wg, &responses)
	// defer close(responseChannel)
	wg.Wait()
	log.Println("Number of responses:", len(responses.Responses))
	printSortedRespones(responses.Responses, 0)
	printSortedRespones(responses.Responses, len(responses.Responses)-1)

}

func printSortedRespones(responseArray []PlayerResponse, indexToPrint int) {
	sort.Slice(responseArray[:], func(i, j int) bool {
		return responseArray[i].LastIndex < responseArray[j].LastIndex
	})
	if indexToPrint >= 0 {
		log.Println("After sorting, the response at index", indexToPrint, "is", responseArray[indexToPrint])
	} else {
		log.Println("Negative index given, will print all...")
		for ind, response := range responseArray {
			log.Println(ind, response)
		}
	}
	return
}

// func monitorResponseChannel(channel <-chan PlayerResponse, playerData *[]PlayerCard, waitgroup *sync.WaitGroup, responseArray *[]PlayerResponse) {
// 	defer waitgroup.Done()
// 	for response := range channel {
// 		// log.Println("Have a winner! Board:", response.ID, "Last Played:", response.LastPlayed, "Board Sum:", response.BoardSum, "Winning Index:", response.LastIndex)
// 		(*playerData)[response.ID].mu.Lock()
// 		(*playerData)[response.ID].Enabled = false
// 		(*playerData)[response.ID].mu.Unlock()
// 		*responseArray = append(*responseArray, response)
// 	}
// 	return
// }

func bingoPlayer(card *PlayerCard, waitgroup *sync.WaitGroup, responseStruct *ResponseArray) {
	defer waitgroup.Done()
	for msg := range card.Channel {
		found, xy := bingoHit(card.Board, msg[0])
		if found {
			card.mu.Lock()
			card.Board[xy[0]][xy[1]] = -1
			card.mu.Unlock()
		}
		if bingoWin(card.Board) {
			card.mu.Lock()
			card.Enabled = false
			card.mu.Unlock()
			boardSum := bingoSum(card.Board)
			responseStruct.mu.Lock()
			(*responseStruct).Responses = append((*responseStruct).Responses, PlayerResponse{ID: card.ID, LastPlayed: msg[0], LastIndex: msg[1], BoardSum: boardSum})
			responseStruct.mu.Unlock()
			break
		}
	}
	// log.Println("Player", card.ID, "is online and ready to play!")
	return
}

func bingoHit(board [][]int, call int) (bool, []int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == call {
				return true, []int{i, j}
			}
		}
	}
	return false, nil
}

func bingoSum(board [][]int) int {
	sum := 0
	for i := 0; i < len(board); i++ {
		for _, val := range board[i] {
			if val >= 0 {
				sum += val
			}
		}
	}
	return sum
}

func bingoWin(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		if sum(board[i]...) == (-1 * len(board)) {
			return true
		}
		tempArray := make([]int, len(board))
		for ind := 0; ind < len(board); ind++ {
			tempArray[ind] = board[ind][i]
		}
		if sum(tempArray...) == (-1 * len(board)) {
			return true
		}
	}
	return false
}
