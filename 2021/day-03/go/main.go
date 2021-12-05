package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var bitLenSet bool
var bitLen int
var numbers []uint16

func main() {
	preProcessData()

	gammaRate := dataProcessorPart1(numbers)
	epsilonRate := ^gammaRate & ((^uint16(0) << (16 - bitLen)) >> (16 - bitLen))

	oxygenRate := dataProcessorPart2(numbers, true, bitLen-1)
	co2Rate := dataProcessorPart2(numbers, false, bitLen-1)

	fmt.Printf("gamma Rate: %d\nepsilon Rate: %d\noxygen Rate: %d\nco2 Rate: %d\nPart 1: %d\nPart 2: %d\n", gammaRate, epsilonRate, oxygenRate, co2Rate, uint64(epsilonRate)*uint64(gammaRate), uint64(oxygenRate)*uint64(co2Rate))
}

func dataProcessorPart1(data []uint16) (result uint16) {
	dataSize := uint16(len(numbers))
	refNum := dataSize >> 1

	for i := bitLen - 1; i >= 0; i-- {
		sum := uint16(0)
		for _, n := range data {
			sum += (n & (1 << i)) >> i
		}
		if sum >= refNum {
			result += 1 << i
		}
	}
	return
}

func dataProcessorPart2(data []uint16, lookForCommon bool, bitPos int) (result uint16) {
	sum := 0
	for _, n := range data {
		sum += int((n & (1 << bitPos)) >> bitPos)
	}

	ref := 0
	if (sum << 1) >= len(data) {
		if lookForCommon {
			ref = 1
		}
	} else {
		if !lookForCommon {
			ref = 1
		}
	}

	var data2 []uint16
	for _, n := range data {
		if (n&(1<<bitPos))>>bitPos == uint16(ref) {
			data2 = append(data2, n)
		}
	}

	if len(data2) == 1 {
		return data2[0]
	} else {
		if bitPos <= 0 {
			log.Fatalln("bitPos go low")
		}

		result = dataProcessorPart2(data2, lookForCommon, bitPos-1)
	}

	return
}

// converts string to number
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
			if !bitLenSet {
				bitLenSet = !bitLenSet
				bitLen = len(text)
			}
			n, err := strconv.ParseUint(text, 2, 64)
			if err != nil {
				log.Fatalln("Invalid data", text)
			}
			numbers = append(numbers, uint16(n))
		}
	}

	file.Close()
}
