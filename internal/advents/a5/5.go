package a5

import (
	"dekamik/aoc_2024/internal/structure"
	"strconv"
	"strings"
)

type day5 struct {
}

type rule struct {
	before int
	after int
}

func (r rule) isValid(row []int) bool {
	var isValid bool = true

	for i, current := range row {
		if current == r.before {
			for j := 0; j < i; j++ {
				if row[j] == r.after {
					isValid = false
					break
				}
			}
		} else if current == r.after {
			for j := i+1; j < len(row); j++ {
				if row[j] == r.before {
					isValid = false
					break
				}
			}
		}

		if !isValid {
			break
		}
	}

	return isValid
}

func parseRule(s string) (*rule, error) {
	raw := strings.Split(s, "|")

	before, err := strconv.Atoi(raw[0])
	if err != nil {
		return nil, err
	}

	after, err := strconv.Atoi(raw[1])
	if err != nil {
		return nil, err
	}

	return &rule{
		before: before,
		after: after,
	}, nil
}

func getCorrectlyOrderedUpdates(input string) []int {
    panic("unimplemented")
}

func sumMiddleNumber(input [][]int) int {
    panic("unimplemented")
}

// ExecutePart1 implements structure.Challenge.
func (d day5) ExecutePart1() {
	panic("unimplemented")
}

// ExecutePart2 implements structure.Challenge.
func (d day5) ExecutePart2() {
	panic("unimplemented")
}

var _ structure.Challenge = day5{}

func New() structure.Challenge {
    return day5{}
}
