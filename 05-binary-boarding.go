package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func bSearch(lower int, upper int, isHighArray []bool) int {
	if len(isHighArray) == 1 {
		if isHighArray[0] {
			return upper
		}
		return lower
	}

	mid := lower + (upper-lower)/2
	nextSequence := isHighArray[1:]
	if isHighArray[0] {
		return bSearch(mid+1, upper, nextSequence)
	} else {
		return bSearch(lower, mid, nextSequence)
	}
}

func getRow(rowString string) int {
	isHighArray := []bool{}
	for _, rune := range rowString {
		if string(rune) == "F" {
			isHighArray = append(isHighArray, false)
		} else {
			isHighArray = append(isHighArray, true)
		}
	}
	return bSearch(0, 127, isHighArray)
}

func getColumn(columnString string) int {
	isHighArray := []bool{}
	for _, rune := range columnString {
		if string(rune) == "L" {
			isHighArray = append(isHighArray, false)
		} else {
			isHighArray = append(isHighArray, true)
		}
	}
	return bSearch(0, 7, isHighArray)
}

func getSeatIDFromCoord(row int, column int) int {
	return 8*row + column
}

func getSeatID(boardingPass string) int {
	rowString := boardingPass[:7]
	columnString := boardingPass[7:10]
	seatID := getSeatIDFromCoord(getRow(rowString), getColumn(columnString))

	return seatID
}

func main() {
	data, _ := ioutil.ReadFile("./inputs/05-input.txt")
	lines := strings.Split(string(data), "\n")
	highestSeatID := 0
	mySeatID := 0

	seatMap := make(map[int]bool)
	for _, line := range lines {
		if line == "" {
			break
		}
		seatID := getSeatID(line)
		seatMap[seatID] = true
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	for i := 1; i < highestSeatID; i++ {
		if !seatMap[i] && seatMap[i-1] != false && seatMap[i+1] != false {
			mySeatID = i
		}
	}
	fmt.Println("📝 Highest seat id is", highestSeatID)
	fmt.Println("📝 My seat id is", mySeatID)
}
