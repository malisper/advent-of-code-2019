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

func solve(insts []int) int {
	insts[1] = 12
	insts[2] = 2

	pc := 0

	for insts[pc] != 99 {
		opcode := insts[pc]
		arg1Loc := insts[pc+1]
		arg2Loc := insts[pc+2]
		outputLoc := insts[pc+3]

		if opcode == 1 {
			insts[outputLoc] = insts[arg1Loc] + insts[arg2Loc]
		} else if opcode == 2 {
			insts[outputLoc] = insts[arg1Loc] * insts[arg2Loc]
		}

		pc += 4
	}

	return insts[0]
}
