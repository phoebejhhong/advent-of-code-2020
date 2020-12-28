package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const FLOOR = "."
const EMPTY = "L"
const OCCUPIED = "#"

func printRoom(room [][]string) {
	for i := 0; i < len(room); i++ {
		fmt.Printf("%v\n", room[i])
	}
}

func copyRoom(room [][]string) [][]string {
	newRoom := make([][]string, len(room))
	for rowIndex, row := range room {
		newRoom[rowIndex] = make([]string, len(row))
		copy(newRoom[rowIndex], row)
	}
	return newRoom
}

func isSameRoom(room1 [][]string, room2 [][]string) bool {
	for rowIndex, row := range room1 {
		for seatIndex, seat := range row {
			if room2[rowIndex][seatIndex] != seat {
				return false
			}
		}
	}
	return true
}

func countOccupiedSeats(room [][]string) int {
	count := 0
	for _, row := range room {
		for _, seat := range row {
			if seat == OCCUPIED {
				count++
			}
		}
	}
	return count
}

func getAdjacentSeats1(room [][]string, i int, j int) []string {
	seats := []string{}
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if !(x < 0 || x > len(room)-1 || y < 0 || y > len(room[0])-1 || (x == i && y == j)) {
				seat := room[x][y]
				if seat != FLOOR {
					seats = append(seats, seat)
				}
			}
		}
	}
	return seats
}

func getAdjacentSeats2(room [][]string, i int, j int) []string {
	seats := []string{}

	// travel up
	if i > 0 {
		for x := i - 1; x >= 0; x-- {
			seat := room[x][j]
			if seat != FLOOR {
				seats = append(seats, seat)
				break
			}
		}
	}
	// travel up & left
	if !(i == 0 || j == 0) {
		for x := i - 1; x >= 0; x-- {
			y := j - i + x
			if y < 0 {
				break
			}
			seat := room[x][y]
			if seat != FLOOR {
				seats = append(seats, seat)
				break
			}
		}
	}
	// travel left
	if j > 0 {
		for y := j - 1; y >= 0; y-- {
			seat := room[i][y]
			if seat != FLOOR {
				seats = append(seats, seat)
				break
			}
		}
	}
	// travel bottom & left
	if !(i == len(room)-1 || j == 0) {
		for x := i + 1; x < len(room); x++ {
			y := j - x + i
			if y < 0 {
				break
			}
			seat := room[x][y]

			if seat != FLOOR {
				seats = append(seats, seat)
				break
			}
		}
	}
	// travel bottom
	if i != len(room)-1 {
		for x := i + 1; x < len(room); x++ {
			seat := room[x][j]
			if seat != FLOOR {
				seats = append(seats, seat)
				break
			}
		}
	}
	// travel bottom & right
	if !(i == len(room)-1 || j == len(room[0])-1) {
		for x := i + 1; x < len(room); x++ {
			y := j + x - i
			if y == len(room[0]) {
				break
			}
			seat := room[x][y]

			if seat != FLOOR {
				seats = append(seats, seat)
				break
			}
		}
	}
	// travel right
	if j != len(room[0]) {
		for y := j + 1; y < len(room[0]); y++ {
			seat := room[i][y]
			if seat != FLOOR {
				seats = append(seats, seat)
				break
			}
		}
	}
	// travel up & right
	if !(i == 0 || j == len(room[0])-1) {
		for x := i - 1; x >= 0; x-- {
			y := j + i - x
			if y == len(room[0]) {
				break
			}
			seat := room[x][y]
			if seat != FLOOR {
				seats = append(seats, seat)
				break
			}
		}
	}
	return seats
}

func shouldBeOccupied(room [][]string, i int, j int, part int) bool {
	seats := []string{}
	if part == 1 {
		seats = getAdjacentSeats1(room, i, j)
	} else {
		seats = getAdjacentSeats2(room, i, j)
	}
	for _, seat := range seats {
		if seat == OCCUPIED {
			return false
		}
	}
	return true
}

func shouldBeEmptied(room [][]string, i int, j int, part int) bool {
	seats := []string{}
	maxSeats := 0
	if part == 1 {
		seats = getAdjacentSeats1(room, i, j)
		maxSeats = 4
	} else {
		seats = getAdjacentSeats2(room, i, j)
		maxSeats = 5
	}
	occupiedAdjecentCount := 0
	for _, seat := range seats {
		if seat == OCCUPIED {
			occupiedAdjecentCount++
			if occupiedAdjecentCount == maxSeats {
				return true
			}
		}
	}
	return false
}

func runRules(room [][]string, part int) [][]string {
	newRoom := copyRoom(room)
	for i := 0; i < len(room); i++ {
		row := room[i]
		for j := 0; j < len(row); j++ {
			switch row[j] {
			case FLOOR:
				continue
			case EMPTY:
				if shouldBeOccupied(room, i, j, part) {
					newRoom[i][j] = OCCUPIED
				}
			case OCCUPIED:
				if shouldBeEmptied(room, i, j, part) {
					newRoom[i][j] = EMPTY
				}
			}
		}
	}
	return newRoom
}

func stablizeRoom(room [][]string, part int) [][]string {
	for true {
		newRoom := runRules(room, part)
		if isSameRoom(room, newRoom) {
			return newRoom
		}
		room = newRoom
	}
	return [][]string{}
}

func main() {
	data, _ := ioutil.ReadFile("./inputs/11-input.txt")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	room := make([][]string, len(lines))
	for lineIndex, line := range lines {
		row := make([]string, len(line))
		for charIndex, rune := range line {
			row[charIndex] = string(rune)
		}
		room[lineIndex] = row
	}

	stabilzedRoom := stablizeRoom(room, 1)
	count := countOccupiedSeats(stabilzedRoom)
	fmt.Printf("📝 %v seats end up occupied in part 1.\n", count)

	stabilzedRoom = stablizeRoom(room, 2)
	count = countOccupiedSeats(stabilzedRoom)
	fmt.Printf("📝 %v seats end up occupied in part 2.\n", count)
}
