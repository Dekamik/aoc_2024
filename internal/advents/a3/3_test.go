package a3

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := []mulArgs{
		mulArgs{a: 2, b: 4},
		mulArgs{a: 5, b: 5},
		mulArgs{a: 11, b: 8},
		mulArgs{a: 8, b: 5},
	}
	actual := parseInput(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual %v not equal to expected %v", actual, expected)
	}
}

func TestSum(t *testing.T) {
    input := []mulArgs{
		mulArgs{a: 2, b: 4},
		mulArgs{a: 5, b: 5},
		mulArgs{a: 11, b: 8},
		mulArgs{a: 8, b: 5},
	}
    expected := 161
    actual := sum(input)
    if expected != actual {
		t.Errorf("Actual %v not equal to expected %v", actual, expected)
    }
}

func TestTokenize(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := []token{
        token{function: "mul", arg1: 2, arg2: 4},
        token{function: "don't", arg1: 0, arg2: 0},
        token{function: "mul", arg1: 5, arg2: 5},
        token{function: "mul", arg1: 11, arg2: 8},
        token{function: "do", arg1: 0, arg2: 0},
        token{function: "mul", arg1: 8, arg2: 5},
	}
	actual, err := tokenize(input)
    if err != nil {
        panic(err)
    }
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual %v not equal to expected %v", actual, expected)
	}
}

func TestExecute(t *testing.T) {
	input := []token{
        token{function: "mul", arg1: 2, arg2: 4},
        token{function: "don't", arg1: 0, arg2: 0},
        token{function: "mul", arg1: 5, arg2: 5},
        token{function: "mul", arg1: 11, arg2: 8},
        token{function: "do", arg1: 0, arg2: 0},
        token{function: "mul", arg1: 8, arg2: 5},
	}
    expected := 48
    actual := execute(input)
    if expected != actual {
		t.Errorf("Actual %v not equal to expected %v", actual, expected)
    }
}
