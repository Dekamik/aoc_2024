package a2

import (
	"dekamik/aoc_2024/internal/io"
	"dekamik/aoc_2024/internal/structure"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
)

const (
	Unknown int = iota
	Increasing
	Decreasing
)

type challenge struct {
}

// ExecutePart1 implements structure.Challenge.
func (d challenge) ExecutePart1() {
	reports, err := io.ReadLines("internal/advents/a2/input.txt")
	if err != nil {
		panic(err)
	}

	var safeReports int = 0

	for _, report := range reports {
		rawLevels := strings.Split(report, " ")
		levels := []int{}

		if report == "" {
			continue
		}

		for _, raw := range rawLevels {
			level, err := strconv.Atoi(raw)
			if err != nil {
				panic(err)
			}
			levels = append(levels, level)
		}

		var expectedTrend int = Unknown
		var unsafeLevels int = 0

		for i, level := range levels {
			if i == 0 {
				continue
			}

			a := levels[i-1]
			b := level
			absDiff := int(math.Abs(float64(a - b)))

			if absDiff > 3 || absDiff < 1 {
				slog.Debug("Unsafe: Difference higher than threshold")
				unsafeLevels++

				if i != len(levels)-1 {
					c := levels[i+1]
					newAbsDiff := int(math.Abs(float64(a - c)))

					if newAbsDiff > 3 || newAbsDiff < 1 {
						slog.Debug("Unsafe: Difference higher than threshold despite dampening")
						unsafeLevels++
					}
				}
			}

			if i == 1 {
				if a < b {
					expectedTrend = Increasing
				} else if a > b {
					expectedTrend = Decreasing
				}
			} else {
				if expectedTrend == Increasing && a > b {
					slog.Debug("Unsafe: Fluctuation detected")
					unsafeLevels++
				} else if expectedTrend == Decreasing && a < b {
					slog.Debug("Unsafe: Fluctuation detected")
					unsafeLevels++
				}
			}
		}

		if unsafeLevels <= 1 {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

// ExecutePart2 implements structure.Challenge.
func (d challenge) ExecutePart2() {
	fmt.Println("unimplemented")
}

var _ structure.Challenge = challenge{}

func New() structure.Challenge {
	return challenge{}
}
