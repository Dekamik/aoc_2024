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
    actual := calculateDistinctTiles(input)

    if actual != expected {
        t.Errorf("Expected %v, but got %v", expected, actual)
    }
}
