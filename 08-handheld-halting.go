package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

func runInstructionsAndFindCorruptedCommand(instructions []string) int {
	for i, instruction := range instructions {
		command := strings.Split(string(instruction), " ")[0]
		newInstructions := make([]string, len(instructions))
		copy(newInstructions, instructions)
		switch command {
		case "nop":
			newInstruction := strings.Replace(instruction, "nop", "jmp", -1)
			newInstructions[i] = newInstruction
			accumulator, corrupted := runInstructions(newInstructions)
			if !corrupted {
				return accumulator
			}
		case "jmp":
			newInstruction := strings.Replace(instruction, "jmp", "nop", -1)
			newInstructions[i] = newInstruction
			accumulator, corrupted := runInstructions(newInstructions)
			if !corrupted {
				return accumulator
			}
		}
	}
	return -1
}

func main() {
	data, _ := ioutil.ReadFile("./08-input.txt")
	lines := strings.Split(string(data), "\n")
	instructions := lines[:len(lines)-1]
	accumulator, _ := runInstructions(instructions)
	accumulator2 := runInstructionsAndFindCorruptedCommand(instructions)
	fmt.Println(" 📝 Immediately before any instruction is executed a second time, accumulator value is", accumulator)
	fmt.Println(" 📝 When Program is correctly terminated, accumulator value is", accumulator2)
}
