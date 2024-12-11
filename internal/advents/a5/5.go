package a5

import (
	"dekamik/aoc_2024/internal/io"
	"dekamik/aoc_2024/internal/structure"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type day5 struct {
}

type rule struct {
	before int
	after int
}

func (r rule) isValid(queue []int) bool {
	var isValid bool = true

	for i, current := range queue {
		if current == r.before {
			for j := 0; j < i; j++ {
				if queue[j] == r.after {
					isValid = false
					break
				}
			}
		} else if current == r.after {
			for j := i+1; j < len(queue); j++ {
				if queue[j] == r.before {
					isValid = false
					break
				}
			}
		}

		if !isValid {
			break
		}
	}

	return isValid
}

func parseRule(s string) (*rule, error) {
	raw := strings.Split(s, "|")

	before, err := strconv.Atoi(raw[0])
	if err != nil {
		return nil, err
	}

	after, err := strconv.Atoi(raw[1])
	if err != nil {
		return nil, err
	}

	return &rule{
		before: before,
		after: after,
	}, nil
}

func parseQueue(s string) ([]int, error) {
    results := []int{}
    rawLine := strings.Split(s, ",")

    for _, raw := range rawLine {
        num, err := strconv.Atoi(raw)
        if err != nil {
            return nil, err
        }

        results = append(results, num)
    }

    return results, nil
}

func parseFile(input string) ([]rule, [][]int, error) {
    var isOnQueues bool = false
    var rules []rule = []rule{}
    var queues [][]int = [][]int{}

    lines := strings.Split(input, "\n")

    for _, line := range lines {
        if isOnQueues {
            if line == "" {
                break
            }

            queue, err := parseQueue(line)
            if err != nil {
                return nil, nil, err
            }
            queues = append(queues, queue)

        } else {
            if line == "" {
                isOnQueues = true
                continue
            }

            rule, err := parseRule(line)
            if err != nil {
                return nil, nil, err
            }
            rules = append(rules, *rule)
        }
    }

    return rules, queues, nil
}

func sumMiddleNumber(input [][]int) int {
    var sum int = 0

    for _, row := range input {
        midIndex := int(math.Floor(float64(len(row)) * 0.5))
        sum += row[midIndex]
    }

    return sum
}

func getCorrectlyOrderedUpdates(input string) ([][]int, error) {
    validUpdates := [][]int{}

    rules, queues, err := parseFile(input)
    if err != nil {
        return nil, err
    }

    for _, queue := range queues {
        var invalidRules int = 0

        for _, rule := range rules {
            if !rule.isValid(queue) {
                invalidRules++
            }
        }

        if invalidRules == 0 {
            validUpdates = append(validUpdates, queue)
        }
    }

    return validUpdates, nil
}

// ExecutePart1 implements structure.Challenge.
func (d day5) ExecutePart1() {
    input, err := io.ReadStr("internal/advents/a5/input.txt")
    if err != nil {
        panic(err)
    }

    validQueues, err := getCorrectlyOrderedUpdates(input)
    if err != nil {
        panic(err)
    }

    sum := sumMiddleNumber(validQueues)

    fmt.Println(sum)
}

// ExecutePart2 implements structure.Challenge.
func (d day5) ExecutePart2() {
	fmt.Println("unimplemented")
}

var _ structure.Challenge = day5{}

func New() structure.Challenge {
    return day5{}
}
