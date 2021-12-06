package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var X, Y int
var board [][]int
var points [][]int

func main() {
	preProcessData()
	game1()
}

func game1() {
	board = make([][]int, X)
	for i := 0; i < X; i++ {
		board[i] = make([]int, Y)
	}
	for _, pp := range points {
		x1 := pp[0]
		y1 := pp[1]
		x2 := pp[2]
		y2 := pp[3]
		if x1 == x2 || y1 == y2 {
			count := Abs(x1 - x2 + y1 - y2)

			d := math.Sqrt((float64(x1)-float64(x2))*(float64(x1)-float64(x2))+(float64(y1)-float64(y2))*(float64(y1)-float64(y2))) / float64(count)
			fi := math.Atan2(float64(y2)-float64(y1), float64(x2)-float64(x1))

			for i := 0; i <= count; i++ {
				x := x1 + int(float64(i)*d*math.Cos(fi))
				y := y1 + int(float64(i)*d*math.Sin(fi))
				board[x][y] += 1
			}
		}
	}
	cnt := 0
	for _, row := range board {
		for _, cell := range row {
			if cell >= 2 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

func preProcessData() {
	file, err := os.Open("../puzzle.dat")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			pointsStr := strings.Split(text, " -> ")
			x1, y1 := parsePoint(pointsStr[0])
			x2, y2 := parsePoint(pointsStr[1])
			points = append(points, []int{x1, y1, x2, y2})
			if x1 > X {
				X = x1
			}
			if x2 > X {
				X = x2
			}
			if y1 > Y {
				Y = y1
			}
			if y2 > Y {
				Y = y2
			}
		}
	}

	X += 1
	Y += 1

	file.Close()
}

func parsePoint(p_str string) (x int, y int) {
	strs := strings.Split(p_str, ",")
	x, _ = strconv.Atoi(strs[0])
	y, _ = strconv.Atoi(strs[1])
	return
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
