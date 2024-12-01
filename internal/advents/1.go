package advents

import (
	"dekamik/aoc_2024/internal/command"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type day1part1 struct {
}

// Execute implements internal.Command.
func (d day1part1) Execute() {
	data, err := os.ReadFile("inputs/1-1.txt")
	if err != nil {
		panic(err)
	}


	lhList := []int64{}
	rhList := []int64{}

	re := regexp.MustCompile(`^(\d*)[[:blank:]]*(\d*)$`)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		m1 := matches[1]
		m2 := matches[2]
		
		if (m1 == "" || m2 == "") {
			break
		}

		lh, err := strconv.ParseInt(m1, 10, 64)
		if err != nil {
			panic(err)
		}

		rh, err := strconv.ParseInt(m2, 10, 64)
		if err != nil {
			panic(err)
		}

		lhList = append(lhList, lh)
		rhList = append(rhList, rh)
	}

	sort.Slice(lhList, func(i, j int) bool {
		return lhList[i] < lhList[j]
	})

	sort.Slice(rhList, func(i, j int) bool {
		return rhList[i] < rhList[j]
	})

	var totalDistance int64 = 0

	for i := range lhList {
		totalDistance += int64(math.Abs(float64(lhList[i] - rhList[i])))
	}

	fmt.Println(totalDistance)
}

var _ internal.Command = day1part1{}

func NewDay1Part1() internal.Command {
	return day1part1{}
}
