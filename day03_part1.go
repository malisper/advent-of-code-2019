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

func wireToSet(w *wire) map[point]bool {
	result := map[point]bool{}

	row, col := 0, 0

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

			result[point{row, col}] = true
		}
	}

	return result
}

func solve(wire1 *wire, wire2 *wire) int {
	set1 := wireToSet(wire1)
	set2 := wireToSet(wire2)

	minDist := math.MaxInt32
	for p := range set1 {
		if _, ok := set2[p]; ok {
			dist := abs(p.row) + abs(p.col)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
