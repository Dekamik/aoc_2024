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
}

type guard struct {
    pos pos
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

func parseMap(input string) ([][]tile, guard) {
	var guard guard
	patrolMap := [][]tile{}
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		patrolMap = append(patrolMap, []tile{})

		for j, r := range line {
			tType := parseTile(r)
            isObstructed := false

            switch tType {
            case GUARD:
				guard.pos.x = j
				guard.pos.y = i
				guard.dir = parseDir(r)
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

	return patrolMap, guard
}

func isWithinBounds(patrolMap [][]tile, position pos) bool {
	yBound := len(patrolMap)
	xBound := len(patrolMap[yBound-1])
	return position.x >= 0 && position.y >= 0 && position.x < xBound && position.y < yBound
}

func traverse(patrolMap [][]tile, guard guard) int {
    var distinctTiles map[pos]struct{} = map[pos]struct{}{}

	timeLimitSec := 5
	abortTime := time.Now().Add(time.Second * time.Duration(timeLimitSec))
	for {
		assert.Assert(time.Now().Before(abortTime), "traversal must take less than %v seconds to complete", timeLimitSec)

        slog.Debug("saving pos", "x", guard.pos.x, "y", guard.pos.y)
        distinctTiles[guard.pos] = struct{}{}

		deltaX, deltaY := guard.dir.Delta()
        nextPos := pos{
            x: guard.pos.x + deltaX,
            y: guard.pos.y + deltaY,
        }
        if !isWithinBounds(patrolMap, nextPos) {
            break
        }

        nextTile := patrolMap[nextPos.y][nextPos.x]
		if nextTile.IsObstructed {
			switch guard.dir {
			case DIR_N:
				guard.dir = DIR_E
			case DIR_E:
				guard.dir = DIR_S
			case DIR_S:
				guard.dir = DIR_W
			case DIR_W:
				guard.dir = DIR_N
            default:
                panic("Unknown position " + strconv.QuoteRune(rune(guard.dir)))
			}
			continue
		}

        // Calculate delta with new position
        deltaX, deltaY = guard.dir.Delta()
        guard.pos.x = guard.pos.x+deltaX
        guard.pos.y = guard.pos.y+deltaY
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
