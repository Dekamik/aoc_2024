package advents

import (
	internal "dekamik/aoc_2024/internal/command"
	"dekamik/aoc_2024/internal/io"
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

type day2 struct {
}

// ExecutePart1 implements internal.Challenge.
func (d day2) ExecutePart1() {
    reports, err := io.ReadLines("inputs/2-1.txt")
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
        var isReportSafe bool = true

        for i, level := range levels {
            if i == 0 {
                continue
            }

            a := levels[i-1]
            b := level
            absDiff := int(math.Abs(float64(a - b)))

            if absDiff > 3 {
                slog.Debug("Unsafe: Increase higher than threshold")
                isReportSafe = false
                break
            } else if absDiff < 1 {
                slog.Debug("Unsafe: Increase lower than threshold")
                isReportSafe = false
                break
            }

            if i == 1 {
                if a < b {
                    expectedTrend = Increasing
                } else if a > b {
                    expectedTrend = Decreasing
                }
            } else {
                if expectedTrend == Increasing && a > b {
                    slog.Debug("Unsafe: Fluctiation detected")
                    isReportSafe = false
                    break
                } else if expectedTrend == Decreasing && a < b {
                    slog.Debug("Unsafe: Fluctiation detected")
                    isReportSafe = false
                    break
                }
            }
        }

        if isReportSafe {
            safeReports++
        }
    }

    fmt.Println(safeReports)
}

// ExecutePart2 implements internal.Challenge.
func (d day2) ExecutePart2() {
    fmt.Println("unimplemented")
}

var _ internal.Challenge = day2{}

func NewDay2() internal.Challenge {
	return day2{}
}
