package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var X, Y int
var maps [][]int
var lowPointsMap [][]bool
var curBasinSize int
var basins []int

func main() {
	loadMap()
	findLowPoints()
	findBasins()
}

func findBasins() {
	for _, _ = range maps {
		lowPointsMap = append(lowPointsMap, make([]bool, Y))
	}
	for x, r := range maps {
		for y, _ := range r {
			if maps[x][y] >= 9 {
				lowPointsMap[x][y] = true
			} else {
				lowPointsMap[x][y] = false
			}
		}
	}

	for x, r := range lowPointsMap {
		for y, c := range r {
			if c {
				continue
			}
			curBasinSize = 0
			flood_fill(x, y)
			basins = append(basins, curBasinSize)
		}
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	fmt.Println("Part 2:", basins[0]*basins[1]*basins[2])
}

func flood_fill(x int, y int) {
	curBasinSize++
	lowPointsMap[x][y] = true
	if x > 0 && !lowPointsMap[x-1][y] {
		flood_fill(x-1, y)
	}
	if y > 0 && !lowPointsMap[x][y-1] {
		flood_fill(x, y-1)
	}
	if x < (X-1) && !lowPointsMap[x+1][y] {
		flood_fill(x+1, y)
	}
	if y < (Y-1) && !lowPointsMap[x][y+1] {
		flood_fill(x, y+1)
	}
}

func findLowPoints() {
	points := 0
	for x, row := range maps {
		for y, cell := range row {
			if isLowPoint(x, y) {
				points += cell + 1
			}
		}
	}

	fmt.Println("Part1:", points)
}

func isLowPoint(x int, y int) bool {
	if (x-1 >= 0 && maps[x][y] >= maps[x-1][y]) || (x+1 < X && maps[x][y] >= maps[x+1][y]) || (y-1 >= 0 && maps[x][y] >= maps[x][y-1]) || (y+1 < Y && maps[x][y] >= maps[x][y+1]) {
		return false
	}

	return true
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
			maps = append(maps, arr)
		}
		X = len(maps)
	}

	file.Close()
}
