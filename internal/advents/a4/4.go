package a4

import (
	"dekamik/aoc_2024/internal/assert"
	"dekamik/aoc_2024/internal/structure"
	"dekamik/aoc_2024/internal/io"
	"fmt"
	"log/slog"
	"strings"
)

type day4 struct {
}

type direction int

const (
	E   direction = 0
	SE  direction = 1
	S   direction = 2
	SW  direction = 3
	W   direction = 4
	NW  direction = 5
	N   direction = 6
	NE  direction = 7
	ALL direction = 8
)

func (d direction) String() string {
	switch d {

	case E:
		return "E"
	case SE:
		return "SE"
	case S:
		return "S"
	case SW:
		return "SW"
	case W:
		return "W"
	case NW:
		return "NW"
	case N:
		return "N"
	case NE:
		return "NE"

	default:
		panic("Direction must be between 0 and 7")
	}
}

func getXYDelta(d direction) (int, int) {
	switch d {

	case E:
		return 1, 0
	case SE:
		return 1, 1
	case S:
		return 0, 1
	case SW:
		return -1, 1
	case W:
		return -1, 0
	case NW:
		return -1, -1
	case N:
		return 0, -1
	case NE:
		return 1, -1

	default:
		panic("Direction must be between 0 and 7")
	}
}

func countWordsOnCoordinate(lines []string, word string, x int, y int) int {
	var matches int = 0

	for i := 0; i < int(ALL); i++ {
		deltaX, deltaY := getXYDelta(direction(i))

		xBounds := len(lines[y])
		yBounds := len(lines)
		maxDeltaX := x + (deltaX * len(word))
		maxDeltaY := y + (deltaY * len(word))

		// I don't know why the lower bounds has to be this, but it works anyway
		if (maxDeltaX < -1 || maxDeltaX > xBounds) || (maxDeltaY < -1 || maxDeltaY > yBounds) {
			continue
		}

		for i, letter := range word {
			cx := x + (i * deltaX)
			cy := y + (i * deltaY)
			if len(lines[cy]) == 0 || rune(lines[cy][cx]) != letter {
				goto SKIP
			}
		}

		matches++
		slog.Debug("match found", "x", x, "y", y, "dir", direction(i))

	SKIP:
	}

	return matches
}

func countWord(input string, word string) int {
	assert.Assert(len(word) > 1, "Word HAS to be longer than one letter")

	var count int = 0
	var lines []string = strings.Split(input, "\n")

	for y, line := range lines {
		for x, letter := range line {
			if letter == rune(word[0]) {
				count += countWordsOnCoordinate(lines, word, x, y)
			}
		}
	}

	return count
}

// ExecutePart1 implements internal.Challenge.
func (d day4) ExecutePart1() {
	str, err := io.ReadStr("internal/advents/a4/input.txt")
	if err != nil {
		panic(err)
	}

	count := countWord(str, "XMAS")

	fmt.Println(count)
}

func isXMasOnCoordinate(lines []string, x int, y int) bool {
	// X-MAS is cut off when A is on the edge
	if x == 0 || x == len(lines[y])-1 || y == 0 || y == len(lines)-1 {
		return false
	}

	if len(lines[y+1]) == 0 {
		return false
	}

	se := lines[y+1][x+1]
	sw := lines[y+1][x-1]
	nw := lines[y-1][x-1]
	ne := lines[y-1][x+1]

	match1 := (nw == 'M' && se == 'S') || (nw == 'S' && se == 'M')
	match2 := (sw == 'M' && ne == 'S') || (sw == 'S' && ne == 'M')

	return match1 && match2
}

func countXMas(input string) int {
	var count int = 0
	var lines []string = strings.Split(input, "\n")

	for y, line := range lines {
		for x, letter := range line {
			if letter == 'A' && isXMasOnCoordinate(lines, x, y) {
				count++
				slog.Debug("match found", "x", x, "y", y)
			}
		}
	}

	return count
}

// ExecutePart2 implements internal.Challenge.
func (d day4) ExecutePart2() {
	str, err := io.ReadStr("internal/advents/a4/input.txt")
	if err != nil {
		panic(err)
	}

	count := countXMas(str)

	fmt.Println(count)
}

var _ structure.Challenge = day4{}

func New() structure.Challenge {
	return day4{}
}
