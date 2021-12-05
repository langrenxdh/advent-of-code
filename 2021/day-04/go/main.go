package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var draws []int
var boards [][][]int
var boardWidth int
var boardWidthSet bool
var winners []int

func main() {
	prepareGame()
	play()
}

func play() {
	for _, draw := range draws {
		for board_idx, board := range boards {
			if contains(winners, board_idx) {
				continue
			}
			for r_idx, row := range board {
				for c_idx, cell := range row {
					if cell == draw {
						boards[board_idx][r_idx][c_idx] = -1
					}
				}
			}
		}
		for idx, board := range boards {
			if contains(winners, idx) {
				continue
			}
			for _, row := range board {
				winner := false
				for _, cell := range row {
					if cell == -1 {
						winner = true
					} else {
						winner = false
						break
					}
				}
				if winner {
					announceWinner(idx, draw)
				}
			}
		}
		for idx, board := range boards {
			if contains(winners, idx) {
				continue
			}

			for col := 0; col < boardWidth; col++ {
				winner := false
				for row := 0; row < len(board); row++ {
					cell := board[row][col]
					if cell == -1 {
						winner = true
					} else {
						winner = false
						break
					}
				}
				if winner {
					announceWinner(idx, draw)
				}
			}
		}
	}
}

func announceWinner(idx int, draw int) {
	winners = append(winners, idx)
	board := boards[idx]
	sum := 0
	for _, row := range board {
		for _, cell := range row {
			if cell != -1 {
				sum += cell
			}
		}
	}

	fmt.Println("Win:", "last draw:", draw, "board num:", idx+1, "remaining sum:", sum, "answer:", draw*sum)
}

func prepareGame() {
	file, err := os.Open("../puzzle.dat")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	spaceTrimmer := regexp.MustCompile(`\s+`)

	firstline := true

	var aBoard [][]int

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			if firstline {
				firstline = false
				for _, s := range strings.Split(text, ",") {
					n, _ := strconv.Atoi(s)
					draws = append(draws, n)
				}
			} else {
				var arr []int
				for _, s := range strings.Split(spaceTrimmer.ReplaceAllString(strings.TrimSpace(text), " "), " ") {
					n, _ := strconv.Atoi(s)
					arr = append(arr, n)
				}
				if !boardWidthSet {
					boardWidthSet = true
					boardWidth = len(arr)
				}
				aBoard = append(aBoard, arr)
			}
		} else if !firstline {
			if len(aBoard) > 0 {
				boards = append(boards, aBoard)
			}
			aBoard = nil
		}
	}

	file.Close()
}

func contains(n []int, t int) bool {
	for _, v := range n {
		if v == t {
			return true
		}
	}

	return false
}
