package main

// import (
// 	"log"
// 	"sync"
// )

// type PlayerResponse struct {
// 	ID         int
// 	LastPlayed int
// 	LastIndex  int
// 	BoardSum   int
// }

// type PlayerCard struct {
// 	ID              int
// 	Enabled         bool
// 	Channel         chan []int
// 	Board           [][]int
// 	ResponseChannel chan PlayerResponse
// }

// func day4(filepath string) {
// 	calls, boards, err := readBingoFile(filepath)
// 	if err != nil {
// 		log.Println("Could not ingest bingo file")
// 	}
// 	log.Println(len(calls), len(boards))
// 	log.Println(boards[0])
// 	log.Println("")
// 	log.Println(boards[len(boards)-1])
// 	var wg sync.WaitGroup

// 	responseChannel := make(chan PlayerResponse, len(boards))
// 	defer close(responseChannel)

// 	// Create the slice of channels to send calls to the boards
// 	playerArray := make([]PlayerCard, len(boards))
// 	for i := range playerArray {
// 		playerArray[i].ID = i
// 		playerArray[i].Enabled = true
// 		playerArray[i].Channel = make(chan []int, len(calls))
// 		playerArray[i].Board = boards[i]
// 		playerArray[i].ResponseChannel = responseChannel
// 	}

// 	// load the channels up with the calls
// 	for index, call := range calls {
// 		for i := range playerArray {
// 			if playerArray[i].Enabled {
// 				playerArray[i].Channel <- []int{call, index}
// 			}
// 		}
// 	}

// 	// close the channels since all sends have happened
// 	for i := range playerArray {
// 		close(playerArray[i].Channel)
// 	}

// 	for i := range playerArray {
// 		wg.Add(1)
// 		go bingoPlayer(playerArray[i], &wg)
// 	}

// 	wg.Add(1)
// 	go monitorResponseChannel(responseChannel, &playerArray, &wg)

// 	wg.Wait()
// }

// func monitorResponseChannel(channel <-chan PlayerResponse, playerData *[]PlayerCard, waitgroup *sync.WaitGroup) {
// 	defer waitgroup.Done()
// 	for response := range channel {
// 		log.Println("Have a winner! Board:", response.ID, "Last Played:", response.LastPlayed, "Board Sum:", response.BoardSum, "Winning Index:", response.LastIndex)
// 		// (*playerData)[response.ID].Enabled = false
// 	}
// 	return
// }

// func bingoPlayer(card PlayerCard, waitgroup *sync.WaitGroup) {
// 	defer waitgroup.Done()
// 	for msg := range card.Channel {
// 		found, xy := bingoHit(card.Board, msg[0])
// 		if found {
// 			card.Board[xy[0]][xy[1]] = -1
// 		}
// 		if bingoWin(card.Board) {
// 			boardSum := bingoSum(card.Board)
// 			card.ResponseChannel <- PlayerResponse{ID: card.ID, LastPlayed: msg[0], LastIndex: msg[1], BoardSum: boardSum}
// 			break
// 		}
// 	}
// 	log.Println("Player", card.ID, "is online and ready to play!")
// 	return
// }

// func bingoHit(board [][]int, call int) (bool, []int) {
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board); j++ {
// 			if board[i][j] == call {
// 				return true, []int{i, j}
// 			}
// 		}
// 	}
// 	return false, nil
// }

// func bingoSum(board [][]int) int {
// 	sum := 0
// 	for i := 0; i < len(board); i++ {
// 		for val := range board[i] {
// 			if val >= 0 {
// 				sum += val
// 			}
// 		}
// 	}
// 	return sum
// }

// func bingoWin(board [][]int) bool {
// 	for i := 0; i < len(board); i++ {
// 		if sum(board[i]...) == (-1 * len(board)) {
// 			return true
// 		}
// 		tempArray := make([]int, len(board))
// 		for ind := 0; ind < len(board); ind++ {
// 			tempArray[ind] = board[ind][i]
// 		}
// 		if sum(tempArray...) == (-1 * len(board)) {
// 			return true
// 		}
// 	}
// 	return false
// }
