package a3

import internal "dekamik/aoc_2024/internal/command"

type challenge struct {
}

type mulArgs struct {
    a int
    b int
}

func mul(args mulArgs) int {
    return args.a * args.b
}

func parseInput(s string) []mulArgs {
    return []mulArgs{}
}

// ExecutePart1 implements internal.Challenge.
func (c challenge) ExecutePart1() {
	panic("unimplemented")
}

// ExecutePart2 implements internal.Challenge.
func (c challenge) ExecutePart2() {
	panic("unimplemented")
}

var _ internal.Challenge = challenge{}

func New() internal.Challenge {
    return challenge{}
}
