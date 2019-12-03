package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Printf("got error: %+v", err)
		return
	}

	fmt.Println(solve(input))
}

func readInput() ([]int, error) {
	fileContents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	trimmedContents := strings.Trim(string(fileContents), "\n")

	result := []int{}
	for _, line := range strings.Split(trimmedContents, ",") {
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		result = append(result, val)
	}

	return result, nil
}

func runProgram(insts []int, noun int, verb int) int {
	insts[1] = noun
	insts[2] = verb

	pc := 0

	for insts[pc] != 99 {
		if insts[pc] == 1 {
			loc1 := insts[pc+1]
			loc2 := insts[pc+2]
			outputLoc := insts[pc+3]

			insts[outputLoc] = insts[loc1] + insts[loc2]
		} else if insts[pc] == 2 {
			loc1 := insts[pc+1]
			loc2 := insts[pc+2]
			outputLoc := insts[pc+3]

			insts[outputLoc] = insts[loc1] * insts[loc2]
		}

		pc += 4
	}

	return insts[0]
}

func solve(insts []int) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			instsCopy := append([]int(nil), insts...)

			if runProgram(instsCopy, noun, verb) == 19690720 {
				return 100*noun + verb
			}
		}
	}

	return -1
}
