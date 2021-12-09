package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	processData()
}

func processData() {
	file, err := os.Open("../puzzle.dat")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	game1 := 0
	game2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			strs := strings.Split(text, "|")
			inputs := strings.Split(strings.TrimSpace(strs[0]), " ")
			outputs := strings.Split(strings.TrimSpace(strs[1]), " ")

			for _, o := range outputs {
				if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
					game1 += 1
				}
			}

			sort.Slice(inputs, func(i, j int) bool {
				return len(inputs[i]) < len(inputs[j])
			})

			numbers := make([]string, 10)
			numbers[1] = inputs[0]
			numbers[4] = inputs[2]
			numbers[7] = inputs[1]
			numbers[8] = inputs[9]

			inputs = remove(inputs, 9)
			inputs = remove(inputs, 2)
			inputs = remove(inputs, 1)
			inputs = remove(inputs, 0)

			for pos, s := range inputs {
				if len(s) == 5 && contains(s, numbers[1]) {
					numbers[3] = s
					inputs = remove(inputs, pos)
					break
				}
			}

			four_one := substract(numbers[4], numbers[1])
			for pos, s := range inputs {
				if len(s) == 5 && contains(s, four_one) {
					numbers[5] = s
					inputs = remove(inputs, pos)
					break
				}
			}

			for pos, s := range inputs {
				if len(s) == 5 {
					numbers[2] = s
					inputs = remove(inputs, pos)
					break
				}
			}
			for pos, s := range inputs {
				if contains(s, numbers[5]) {
					continue
				} else {
					numbers[0] = s
					inputs = remove(inputs, pos)
					break
				}
			}
			zero_one := substract(numbers[0], numbers[1])
			for pos, s := range inputs {
				if contains(s, zero_one) {
					numbers[6] = s
					inputs = remove(inputs, pos)
					break
				}
			}
			numbers[9] = inputs[0]

			var nums []int
			for pos, o := range outputs {
				nums = append(nums, index(numbers, o))
				game2 += index(numbers, o) * int(math.Pow10(len(outputs)-pos-1))
			}
			fmt.Println(numbers, nums)
		}
	}

	fmt.Println(game1, game2)

	file.Close()
}

func contains(str string, sub string) bool {
	if len(sub) == 0 {
		return false
	}
	for _, c := range sub {
		if !strings.Contains(str, string(c)) {
			return false
		}
	}
	return true
}

func equals(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, c := range str1 {
		if !strings.Contains(str2, string(c)) {
			return false
		}
	}
	return true
}

func substract(str string, sub string) string {
	for _, c := range sub {
		str = strings.Replace(str, string(c), "", -1)
	}
	return str
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func index(s []string, str string) int {
	for pos, v := range s {
		if equals(str, v) {
			return pos
		}
	}

	return -1
}
