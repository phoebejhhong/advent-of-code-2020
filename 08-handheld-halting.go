package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type command struct {
}

func runInstructions(instructions []string) (int, bool) {
	accumulator := 0
	i := 0
	visitedMap := map[int]bool{}
	running := true
	corrupted := false

	for running {
		if visitedMap[i] {
			running = false
			corrupted = true
		} else if i > len(instructions)-1 {
			running = false
		} else {
			visitedMap[i] = true
			instruction := instructions[i]
			splitted := strings.Split(string(instruction), " ")
			command := splitted[0]
			if command == "nop" {
				i++
			} else {
				secondPart := []rune(splitted[1])
				operator := string(secondPart[0])
				number, _ := strconv.Atoi(string(secondPart[1:]))

				if command == "acc" {
					if operator == "+" {
						accumulator += number
					} else {
						accumulator -= number
					}
					i++
				} else { // jmp
					if operator == "+" {
						i += number
					} else {
						i -= number
					}
				}
			}
		}
	}
	return accumulator, corrupted
}

func main() {
	data, _ := ioutil.ReadFile("./08-input.txt")
	lines := strings.Split(string(data), "\n")
	instructions := lines[:len(lines)-1]
	accumulator, _ := runInstructions(instructions)
	fmt.Println(" 📝 Immediately before any instruction is executed a second time, accumulator is", accumulator)
}
