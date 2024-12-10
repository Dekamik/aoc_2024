package a6

import "testing"

func TestCalculateDistinctTiles(t *testing.T) {
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
