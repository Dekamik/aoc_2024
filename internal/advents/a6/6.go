package a6

import "dekamik/aoc_2024/internal/structure"

type day6 struct {
}

type tileType int

const (
	GROUND tileType = iota
	OBSTRUCTION
)

type tile struct {
    Type tileType
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
        fallthrough
    case '.':
        fallthrough
    default:
        return GROUND
    }
}

func calculateDistinctTiles(input string) int {
    panic("unimplemented")
}

// ExecutePart1 implements structure.Challenge.
func (d day6) ExecutePart1() {
	panic("unimplemented")
}

// ExecutePart2 implements structure.Challenge.
func (d day6) ExecutePart2() {
	panic("unimplemented")
}

var _ structure.Challenge = day6{}

func New() structure.Challenge {
	return day6{}
}
