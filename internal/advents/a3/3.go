package a3

import (
	internal "dekamik/aoc_2024/internal/command"
	"dekamik/aoc_2024/internal/io"
	"fmt"
	"regexp"
	"strconv"
)

var reMul = regexp.MustCompile(`(mul)\((\d+)\,(\d+)\)`)

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
	instructions := []mulArgs{}
	matches := reMul.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		a, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(match[3])
		if err != nil {
			panic(err)
		}

		token := mulArgs{a: a, b: b}
		instructions = append(instructions, token)
	}

	return instructions
}

func sum(args []mulArgs) int {
    var result int = 0
    for _, arg := range args {
        result += mul(arg)
    }
    return result
}

// ExecutePart1 implements internal.Challenge.
func (c challenge) ExecutePart1() {
    str, err := io.ReadStr("internal/advents/a3/input.txt")
    if err != nil {
        panic(err)
    }

    instructions := parseInput(str)
    result := sum(instructions)

    fmt.Println(result)
}

// ExecutePart2 implements internal.Challenge.
func (c challenge) ExecutePart2() {
	fmt.Println("unimplemented")
}

var _ internal.Challenge = challenge{}

func New() internal.Challenge {
	return challenge{}
}
