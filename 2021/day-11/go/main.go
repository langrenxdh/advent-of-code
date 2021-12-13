package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var octopuses [][]int
var X, Y int
var flashCount int

func main() {
	loadMap()
	allFlashed := false
	for i := 0; i < 100 || !allFlashed; i++ {
		if allFlashed = takeTurn(); allFlashed {
			fmt.Println("all flashed at turn:", i+1)
		}
		if i == 99 {
			fmt.Println("total flashes:", flashCount)
		}
	}
}

func takeTurn() bool {
	for x, r := range octopuses {
		for y := range r {
			octopuses[x][y]++
			checkEngergy(x, y)
		}
	}
	cnt := 0
	for x, r := range octopuses {
		for y, c := range r {
			if c < 0 {
				octopuses[x][y] = 0
				cnt++
			}
		}
	}
	return cnt == X*Y
}

func checkEngergy(x int, y int) {
	if octopuses[x][y] == 10 {
		octopuses[x][y] = -1000
		flashCount++
		if x > 0 {
			octopuses[x-1][y]++
			checkEngergy(x-1, y)
			if y > 0 {
				octopuses[x-1][y-1]++
				checkEngergy(x-1, y-1)
			}
		}
		if y > 0 {
			octopuses[x][y-1]++
			checkEngergy(x, y-1)
			if x < (X - 1) {
				octopuses[x+1][y-1]++
				checkEngergy(x+1, y-1)
			}
		}
		if x < (X - 1) {
			octopuses[x+1][y]++
			checkEngergy(x+1, y)
			if y < (Y - 1) {
				octopuses[x+1][y+1]++
				checkEngergy(x+1, y+1)
			}
		}
		if y < (Y - 1) {
			octopuses[x][y+1]++
			checkEngergy(x, y+1)
			if x > 0 {
				octopuses[x-1][y+1]++
				checkEngergy(x-1, y+1)
			}
		}
	}
}

func loadMap() {
	file, err := os.Open("../puzzle.dat")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if Y = len(text); Y > 0 {
			arr := make([]int, Y)
			for pos, c := range text {
				arr[pos] = int(c - '0')
			}
			octopuses = append(octopuses, arr)
		}
		X = len(octopuses)
	}

	file.Close()
}
