package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type wire struct {
	segments []*wireSegment
}

type wireSegment struct {
	dir    byte
	length int
}

func readInput() (*wire, *wire, error) {
	fileContents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	trimmedContents := strings.Trim(string(fileContents), "\n")

	wires := []*wire{}
	for _, line := range strings.Split(trimmedContents, "\n") {
		w, err := parseWire(line)
		if err != nil {
			return nil, nil, err
		}
		wires = append(wires, w)
	}

	return wires[0], wires[1], nil
}

func parseWire(line string) (*wire, error) {
	segments := []*wireSegment{}

	for _, part := range strings.Split(line, ",") {
		dir := part[0]

		length, err := strconv.Atoi(part[1:])
		if err != nil {
			return nil, errors.WithStack(err)
		}

		segment := &wireSegment{dir, length}
		segments = append(segments, segment)
	}

	return &wire{segments}, nil
}

func main() {
	wire1, wire2, err := readInput()
	if err != nil {
		fmt.Printf("got error: %+v", err)
		return
	}

	fmt.Println(solve(wire1, wire2))
}

type point struct {
	row int
	col int
}

func wireToMap(w *wire) map[point]int {
	result := map[point]int{}

	row, col, dist := 0, 0, 0

	for _, segment := range w.segments {
		for i := 0; i < segment.length; i++ {
			if segment.dir == 'R' {
				col += 1
			} else if segment.dir == 'L' {
				col -= 1
			} else if segment.dir == 'U' {
				row += 1
			} else if segment.dir == 'D' {
				row -= 1
			}

			dist++
			if _, ok := result[point{row, col}]; !ok {
				result[point{row, col}] = dist
			}
		}
	}

	return result
}

func solve(wire1 *wire, wire2 *wire) int {
	map1 := wireToMap(wire1)
	map2 := wireToMap(wire2)

	minDist := math.MaxInt32
	for p := range map1 {
		if _, ok := map2[p]; ok {
			dist := map1[p] + map2[p]
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}
