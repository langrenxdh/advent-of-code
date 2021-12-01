package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var cnt int = 0
var totalProcessedNums = 0
var step = 1
var pastRead []int

/*
	The two puzzles can be thought of as follows:
	1. Sliding window, where size 1 denotes the size of the window, size 3 denotes the size of the window, and so on.
	2. Comparing the sum of all components within the sliding window is as simple as comparing the last element of the previous window to the last element of the current window.
	As a result, we can omit the sum calculation.
*/
func main() {
	flag.IntVar(&step, "s", 1, "Size of sliding window, minimum is 1.")
	flag.Parse()

	if step < 1 {
		log.Fatalln("Incorrect size of sliding window.")
	}

	readInput()
	fmt.Println("Sliding window:", step, " ^_^ Total increase:", cnt)
}

func needsCompare() bool {
	return totalProcessedNums >= step
}

func countIncrease(numsStr string) {
	nums := strings.Split(numsStr, "\n")
	for _, n := range nums {
		if len(n) > 0 {
			num, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalln(err)
			}

			if needsCompare() {
				refNum := pastRead[0]
				if num > refNum {
					cnt++
				}
				pastRead = pastRead[1:]
			}

			pastRead = append(pastRead, num)

			totalProcessedNums += 1
		}
	}
}

// read data from stdin pipe
func readInput() {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {
			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		countIncrease(string(buf))

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
}
