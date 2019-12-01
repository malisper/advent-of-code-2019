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
	for _, line := range strings.Split(trimmedContents, "\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		result = append(result, val)
	}

	return result, nil
}

func solve(partMasses []int) int {
	result := 0

	for _, partMass := range partMasses {
		fuelNeeded := partMass/3 - 2

		for fuelNeeded > 0 {
			result += fuelNeeded
			fuelNeeded = fuelNeeded/3 - 2
		}

	}

	return result
}
