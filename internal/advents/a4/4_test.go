package a4

import "testing"

func TestCountWord(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	expected := 18
	actual := countWord(input, "XMAS")

	if actual != expected {
		t.Errorf("Actual %v not equal to expected %v", actual, expected)
	}
}

func TestCountXMas(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	expected := 9
	actual := countXMas(input)

	if actual != expected {
		t.Errorf("Actual %v not equal to expected %v", actual, expected)
	}
}
