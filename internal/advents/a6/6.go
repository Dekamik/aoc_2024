package a6

import (
	"dekamik/aoc_2024/internal/assert"
	"dekamik/aoc_2024/internal/io"
	"dekamik/aoc_2024/internal/structure"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

type day6 struct {
}

type direction int

const (
	DIR_N direction = 0
	DIR_E direction = 1
	DIR_S direction = 2
	DIR_W direction = 3
)

func (d direction) Delta() (int, int) {
	switch d {
	case DIR_N:
		return 0, -1
	case DIR_E:
		return 1, 0
	case DIR_S:
		return 0, 1
	case DIR_W:
		return -1, 0
	default:
		return 0, 0
	}
}

func (d direction) String() string {
    s := strconv.Itoa(int(d))
    return s
}

type tileType int

const (
	GROUND tileType = iota
	OBSTRUCTION
	GUARD
)

type tile struct {
	Type         tileType
	IsObstructed bool
}

type pos struct {
	x   int
	y   int
	dir direction
}

func parseTile(r rune) tileType {
	switch r {
	case '#':
		return OBSTRUCTION

	case '^':
		fallthrough
	case '<':
		fallthrough
	case '>':
		fallthrough
	case 'v':
		return GUARD

	case '.':
		fallthrough
	default:
		return GROUND
	}
}

func parseDir(r rune) direction {
	switch r {
	case '^':
		return DIR_N
	case '>':
		return DIR_E
	case 'v':
		return DIR_S
	case '<':
		return DIR_W
	default:
		panic("unknown direction: '" + string(r) + "'")
	}
}

func parseMap(input string) ([][]tile, pos) {
	var guardPos pos
	patrolMap := [][]tile{}
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		patrolMap = append(patrolMap, []tile{})

		for j, r := range line {
			tType := parseTile(r)
            isObstructed := false

            switch tType {
            case GUARD:
				guardPos.x = j
				guardPos.y = i
				guardPos.dir = parseDir(r)
				tType = GROUND
            case OBSTRUCTION:
                isObstructed = true
            }

			t := tile{
				Type: tType,
                IsObstructed: isObstructed,
			}
			patrolMap[i] = append(patrolMap[i], t)
		}
	}

	return patrolMap, guardPos
}

func isWithinBounds(patrolMap [][]tile, position pos) bool {
	yBound := len(patrolMap)
	xBound := len(patrolMap[yBound-1])
	return position.x > 0 && position.y > 0 && position.x < xBound && position.y < yBound
}

func traverse(patrolMap [][]tile, startPosition pos) int {
    var distinctTiles map[pos]struct{} = map[pos]struct{}{}
	var currentPos pos = startPosition

	timeLimitSec := 5
	abortTime := time.Now().Add(time.Second * time.Duration(timeLimitSec))
	for {
        slog.Debug("pos", "x", currentPos.x, "y", currentPos.y)

		assert.Assert(time.Now().Before(abortTime), "traversal must take less than %v seconds to complete", timeLimitSec)
		if !isWithinBounds(patrolMap, currentPos) {
			break
		}

		deltaX, deltaY := currentPos.dir.Delta()
        nextTile := patrolMap[currentPos.y+deltaY][currentPos.x+deltaX]
		if nextTile.IsObstructed {
			switch currentPos.dir {
			case DIR_N:
				currentPos.dir = DIR_E
			case DIR_E:
				currentPos.dir = DIR_S
			case DIR_S:
				currentPos.dir = DIR_W
			case DIR_W:
				currentPos.dir = DIR_N
            default:
                panic("Unknown position " + strconv.QuoteRune(rune(currentPos.dir)))
			}
			continue
		}

        distinctTiles[currentPos] = struct{}{}

        deltaX, deltaY = currentPos.dir.Delta()
        currentPos.x = currentPos.x+deltaX
        currentPos.y = currentPos.y+deltaY
	}

	return len(distinctTiles)
}

func calculateDistinctTiles(input string) int {
	patrolMap, guardPosition := parseMap(input)
	tiles := traverse(patrolMap, guardPosition)

	return tiles
}

// ExecutePart1 implements structure.Challenge.
func (d day6) ExecutePart1() {
	input, err := io.ReadStr("internal/advents/a6/input.txt")
	if err != nil {
		panic(err)
	}

	tiles := calculateDistinctTiles(input)
	fmt.Println(tiles)
}

// ExecutePart2 implements structure.Challenge.
func (d day6) ExecutePart2() {
	panic("unimplemented")
}

var _ structure.Challenge = day6{}

func New() structure.Challenge {
	return day6{}
}
