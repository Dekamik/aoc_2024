package a5

import (
	"reflect"
	"testing"
)

func TestGetCorrectlyOrderedUpdates(t *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	expected := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
	}
	actual, err := getCorrectlyOrderedUpdates(input)
    if err != nil {
        panic(err)
    }

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected %v but got %v", expected, actual)
    }
}

func TestSumMiddleNumber(t *testing.T) {
    input := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
	}

    expected := 143
    actual := sumMiddleNumber(input)

    if actual != expected {
        t.Errorf("Expected %v but got %v", expected, actual)
    }
}
