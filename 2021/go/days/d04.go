package days

import (
	"fmt"
	"strconv"
	"strings"

	h "github.com/moymat/aoc2021/helpers"
)

type Cell struct {
	value  string
	marked bool
}

type Row = []Cell

type Board = []Row

func createBoard(input string) Board {
	inputRows := strings.Split(input, "\n")
	var board Board
	for _, inputRow := range inputRows {
		cells := strings.Fields(inputRow)
		var row Row
		for _, cell := range cells {
			row = append(row, Cell{value: strings.TrimSpace(cell), marked: false})
		}
		board = append(board, row)
	}
	return board
}

func hasWon(board *Board) bool {
	for _, row := range *board {
		unmarked := len(row)
		for _, cell := range row {
			if cell.marked {
				unmarked--
			}
		}
		if unmarked == 0 {
			return true
		}
	}
	for cell := range *board {
		unmarked := len((*board)[0])
		for col := range (*board)[0] {
			if (*board)[col][cell].marked {
				unmarked--
			}
		}
		if unmarked == 0 {
			return true
		}
	}
	return false
}

func computeTotalPoints(board *Board) int {
	total := 0
	for _, row := range *board {
		for _, cell := range row {
			if !cell.marked {
				val, err := strconv.Atoi(cell.value)
				h.CheckError(err)
				total += val
			}
		}
	}
	return total
}

func markNumber(draw string, board *Board) {
	for i, row := range *board {
		for y, cell := range row {
			if cell.value == draw && !cell.marked {
				(*board)[i][y].marked = true
			}
		}
	}
}

func playRound(draw string, boards *[]Board) int {
	for _, board := range *boards {
		markNumber(draw, &board)
		res := hasWon(&board)
		if res {
			return computeTotalPoints(&board)
		}
	}
	return -1
}

func playRoundToLose(draw string, boards *[]Board, winnerBoardsIdx *[]int) (int, int) {
	for i, board := range *boards {
		if !h.IntSliceContains(winnerBoardsIdx, i) {
			markNumber(draw, &board)
			res := hasWon(&board)
			if res {
				*winnerBoardsIdx = append(*winnerBoardsIdx, i)
				if len(*winnerBoardsIdx) == len(*boards) {
					drawNum, err := strconv.Atoi(draw)
					h.CheckError(err)
					return computeTotalPoints(&board), drawNum
				}
			}
		}
	}
	return -1, -1
}

func RunD04(file string) {
	inputs := strings.Split(h.GetInput("d04", file), "\n\n")

	var draws []string
	for _, number := range strings.Split(inputs[0], ",") {
		draws = append(draws, strings.TrimSpace(number))
	}

	var boards []Board
	for _, board := range inputs[1:] {
		boards = append(boards, createBoard(board))
	}

	var totalPoints int
	var winnerDraw int
	for _, draw := range draws {
		total := playRound(draw, &boards)
		if total >= 0 {
			totalPoints = total
			drawNum, err := strconv.Atoi(draw)
			h.CheckError(err)
			winnerDraw = drawNum
			break
		}
	}

	var lastTotalPoints int
	var lastWinnerDraw int
	var winnerBoardsIdx []int
	for _, draw := range draws {
		lastTotalPoints, lastWinnerDraw = playRoundToLose(draw, &boards, &winnerBoardsIdx)
		if lastTotalPoints > 0 {
			break
		}
	}

	fmt.Println("first to win :", totalPoints*winnerDraw)
	fmt.Println("last to win :", lastTotalPoints*lastWinnerDraw)
}
