package main

import (
	"testing"
	"time"
	"TurkishDraughts/Board"
)

func benchmarkDepths(t *testing.T, b *board.BoardState, testDepth int32) {
	startTime := time.Now()
	board.Hits = 0
	board.Searches = 0
	
	board.MaxDepth = testDepth
	board.MaximumHashDepth = testDepth - 2

	value := b.MinMax(0, -board.AlphaBetaMax, board.AlphaBetaMax, board.NewTable())

	duration := time.Since(startTime).Seconds()
	t.Log("---", testDepth)
	t.Log("Time:", float32(duration), "s")
	t.Log("Hits:", board.Hits/1000, "k  ", int(float64(board.Hits)/duration/1000.0),"k/s")
	t.Log("Searches:", board.Searches/1000, "k  ",int(float64(board.Searches)/duration/1000.0),"k/s")
	t.Log("Standing:", value)
}

func TestBenchStartBoard(t *testing.T){
	//Config
	board.AdvanceWeight = 0.1
	board.TableDepthAllowedInaccuracy = 0

	var i int32 = 5
	for i<=12 {
		b := board.CreateStartingBoard()
		benchmarkDepths(t, &b, i)
		i++
	}	
}

func TestBenchNoAdvanced(t *testing.T){
	//Config
	board.AdvanceWeight = 0.0
	board.TableDepthAllowedInaccuracy = 0

	var i int32 = 5
	for i<=12 {
		b := board.CreateStartingBoard()
		benchmarkDepths(t, &b, i)
		i++
	}
}

/*

	startTime := time.Now()
	board.Searches = 0
	board.Hits = 0

	options := b.MaxTakeBoards()
	if len(options) == 0 {
		options = b.AllMoveBoards()
	}

	var bestValue float32 
	var bestBoard board.BoardState
	output := make(chan PossibleMove)

	for _, branch := range options{
		go analyzeBranch(branch, board.NewTable(), output, depth)
	}

	for i := range options {
		check := <- output
		checkValue := check.value
		checkBoard := check.board
		if i == 0 || (b.Turn == board.White && checkValue > bestValue) || (b.Turn == board.Black && checkValue < bestValue) {
			bestValue = checkValue
			bestBoard = checkBoard
		}

		fmt.Println(i+1, "/", len(options), "=", checkValue)
	}

	
	return &bestBoard


*/