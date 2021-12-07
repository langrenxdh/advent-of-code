package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lanternfish []uint64

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	preProcessData()
	calcAfterDays(80)
	preProcessData()
	calcAfterDays(256)
}

func calcAfterDays(days int) {
	for i := 0; i < days; i++ {
		new6or8 := lanternfish[0]
		for j := 0; j < len(lanternfish)-1; j++ {
			lanternfish[j] = lanternfish[j+1]
		}
		lanternfish[len(lanternfish)-1] = new6or8
		lanternfish[6] += new6or8
	}
	fmt.Println(sum(lanternfish))
}

func sum(fish []uint64) (sum uint64) {
	for _, f := range fish {
		sum += f
	}
	return
}

func preProcessData() {
	lanternfish = make([]uint64, 9)

	file, err := os.Open("../puzzle.dat")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			strs := strings.Split(text, ",")
			for _, s := range strs {
				n, _ := strconv.Atoi(s)
				lanternfish[n] += 1
			}
		}
	}

	file.Close()
}
