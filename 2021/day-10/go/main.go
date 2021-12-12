package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var syntax map[rune]rune

func main() {
	syntax = make(map[rune]rune)
	syntax['('] = ')'
	syntax['['] = ']'
	syntax['{'] = '}'
	syntax['<'] = '>'

	game1()
	game2()
}

func game2() {
	weight := make(map[rune]int)
	weight[')'] = 1
	weight[']'] = 2
	weight['}'] = 3
	weight['>'] = 4

	file, err := os.Open("../puzzle.dat")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var scores []int
	for scanner.Scan() {
		text := scanner.Text()
		var left []rune
		invalid := false
		for _, c := range text {
			if _, ok := syntax[c]; ok {
				left = append(left, c)
			} else {
				if len(left) <= 0 || syntax[left[len(left)-1]] != c {
					invalid = true
					break
				} else {
					left = left[:len(left)-1]
				}
			}
		}
		if !invalid && len(left) > 0 {
			score := 0
			for i := len(left) - 1; i >= 0; i-- {
				score = score*5 + weight[syntax[left[i]]]
			}
			scores = append(scores, score)
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	fmt.Println("Errors:", scores[len(scores)/2])
}

func game1() {
	weight := make(map[rune]int)
	weight[')'] = 3
	weight[']'] = 57
	weight['}'] = 1197
	weight['>'] = 25137

	file, err := os.Open("../puzzle.dat")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var invalidSyntax []rune
	for scanner.Scan() {
		text := scanner.Text()
		var left []rune
		for _, c := range text {
			if _, ok := syntax[c]; ok {
				left = append(left, c)
			} else {
				if len(left) <= 0 || syntax[left[len(left)-1]] != c {
					invalidSyntax = append(invalidSyntax, c)
					break
				} else {
					left = left[:len(left)-1]
				}
			}
		}
	}
	errors := 0
	for _, c := range invalidSyntax {
		errors += weight[c]
	}
	fmt.Println("Errors:", errors)
}
