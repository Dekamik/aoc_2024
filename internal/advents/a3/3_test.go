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
