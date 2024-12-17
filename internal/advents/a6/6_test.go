package a6

import (
	"dekamik/aoc_2024/internal/logging"
	"log/slog"
	"testing"
)

func TestCalculateDistinctTiles(t *testing.T) {
	logging.NewDefault(slog.LevelDebug)

	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expected := 41
	actual, err := calculateDistinctTiles(input)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestCalculateInfiniteLoops(t *testing.T) {
	logging.NewDefault(slog.LevelInfo)

	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expected := 6
	actual, err := calculateInfiniteLoops(input)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
