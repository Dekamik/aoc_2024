package a6

import (
	"dekamik/aoc_2024/internal/io"
	"dekamik/aoc_2024/internal/structure"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

var infiniteLoopError error = errors.New("detected an infinite loop")
var timeoutError error = errors.New("timeout")

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
	x int
	y int
}

func (p pos) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
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
				Type:         tType,
				IsObstructed: isObstructed,
			}
			patrolMap[i] = append(patrolMap[i], t)
		}
	}

	return patrolMap, guard
}

func getBounds(patrolMap [][]tile) (int, int, int, int) {
	yBound := len(patrolMap)
	xBound := len(patrolMap[yBound-1])
	return -1, -1, xBound, yBound
}

func isWithinBounds(patrolMap [][]tile, position pos) bool {
	xMin, yMin, xMax, yMax := getBounds(patrolMap)
	return position.x > xMin && position.y > yMin && position.x < xMax && position.y < yMax
}

func traverse(patrolMap [][]tile, guard guard) (int, error) {
	var currentCollisions int = 0
	var distinctTiles map[string]struct{} = map[string]struct{}{}

	xMin, yMin, xMax, yMax := getBounds(patrolMap)
	slog.Debug("traversing map", "xMin", xMin, "yMin", yMin, "xMax", xMax, "yMax", yMax)

	timeLimitSec := 5
	abortTime := time.Now().Add(time.Second * time.Duration(timeLimitSec))
	for {
		if time.Now().After(abortTime) {
			return -1, timeoutError
		}

		// A better way to detect infinite loops is to count turns - if the last
		// 4 turns have the same coordinates, we are in an infinite loop.
		// This works "okay" for our purposes, although it's very bruteforce and
		// unelegant.
		if _, exists := distinctTiles[guard.pos.String()]; exists {
			currentCollisions++

			// High number for accuracy
			if currentCollisions > 1000 {
				return -1, infiniteLoopError
			}
		} else {
			currentCollisions = 0
		}

		slog.Debug("saving pos", "x", guard.pos.x, "y", guard.pos.y)
		distinctTiles[guard.pos.String()] = struct{}{}

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

		// Calculate delta with final direction
		deltaX, deltaY = guard.dir.Delta()
		guard.pos.x = guard.pos.x + deltaX
		guard.pos.y = guard.pos.y + deltaY
	}

	return len(distinctTiles), nil
}

func calculateDistinctTiles(input string) (int, error) {
	patrolMap, guard := parseMap(input)
	tiles, err := traverse(patrolMap, guard)
	if err != nil {
		return -1, err
	}

	return tiles, nil
}

func calculateInfiniteLoops(input string) (int, error) {
	var possibleObstructions int = 0

	patrolMap, guard := parseMap(input)
	startPosition := guard.pos

	for i, row := range patrolMap {
		for j, tile := range row {
			if tile.IsObstructed || (j == startPosition.x && i == startPosition.y) {
				continue
			}

			patrolMap[i][j].IsObstructed = true

			_, err := traverse(patrolMap, guard)
			if err != nil {
				if err == infiniteLoopError {
					slog.Info("found obstruction", "x", j, "y", i)
					possibleObstructions++
				} else {
					return -1, err
				}
			}

			patrolMap[i][j].IsObstructed = false
		}
	}

	return possibleObstructions, nil
}

// ExecutePart1 implements structure.Challenge.
func (d day6) ExecutePart1() {
	input, err := io.ReadStr("internal/advents/a6/input.txt")
	if err != nil {
		panic(err)
	}

	// Remove last newline
	input = input[:len(input)-1]
	tiles, err := calculateDistinctTiles(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(tiles)
}

// ExecutePart2 implements structure.Challenge.
func (d day6) ExecutePart2() {
	input, err := io.ReadStr("internal/advents/a6/input.txt")
	if err != nil {
		panic(err)
	}

	// Remove last newline
	input = input[:len(input)-1]
	obstructions, err := calculateInfiniteLoops(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(obstructions)
}

var _ structure.Challenge = day6{}

func New() structure.Challenge {
	return day6{}
}
