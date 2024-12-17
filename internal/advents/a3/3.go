package a3

import (
	"dekamik/aoc_2024/internal/io"
	"dekamik/aoc_2024/internal/structure"
	"fmt"
	"regexp"
	"strconv"
)

var reMul = regexp.MustCompile(`(mul)\((\d+)\,(\d+)\)`)
var reTokens = regexp.MustCompile(`do\(\)|don't\(\)|(mul)\((\d+)\,(\d+)\)`)

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

// ExecutePart1 implements structure.Challenge.
func (c challenge) ExecutePart1() {
	str, err := io.ReadStr("internal/advents/a3/input.txt")
	if err != nil {
		panic(err)
	}

	instructions := parseInput(str)
	result := sum(instructions)

	fmt.Println(result)
}

type token struct {
	function string
	arg1     int
	arg2     int
}

func tokenize(s string) ([]token, error) {
	tokens := []token{}
	matches := reTokens.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		var tkn token

		switch match[0] {

		case "do()":
			fallthrough
		case "don't()":
			trimmed := match[0][:len(match[0])-2]
			tkn = token{
				function: trimmed,
			}

		default:
			if match[1] == "mul" {
				a, err := strconv.Atoi(match[2])
				if err != nil {
					return nil, err
				}

				b, err := strconv.Atoi(match[3])
				if err != nil {
					return nil, err
				}

				tkn = token{
					function: match[1],
					arg1:     a,
					arg2:     b,
				}
			}
		}

		tokens = append(tokens, tkn)
	}

	return tokens, nil
}

func execute(program []token) int {
	var enabled bool = true
	var result int = 0

	for _, token := range program {
		switch token.function {

		case "do":
			enabled = true

		case "don't":
			enabled = false

		case "mul":
			if enabled {
				args := mulArgs{
					a: token.arg1,
					b: token.arg2,
				}
				result += mul(args)
			}
		}
	}

	return result
}

// ExecutePart2 implements structure.Challenge.
func (c challenge) ExecutePart2() {
	str, err := io.ReadStr("internal/advents/a3/input.txt")
	if err != nil {
		panic(err)
	}

	instructions, err := tokenize(str)
	if err != nil {
		panic(err)
	}

	result := execute(instructions)

	fmt.Println(result)
}

var _ structure.Challenge = challenge{}

func New() structure.Challenge {
	return challenge{}
}
