package a5

import (
	"dekamik/aoc_2024/internal/assert"
	"dekamik/aoc_2024/internal/io"
	"dekamik/aoc_2024/internal/structure"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

type day5 struct {
}

type rule struct {
	before int
	after  int
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
			for j := i + 1; j < len(queue); j++ {
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
		after:  after,
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

func getIncorrectlyOrderedUpdates(input string) ([][]int, []rule, error) {
	invalidUpdates := [][]int{}

	rules, queues, err := parseFile(input)
	if err != nil {
		return nil, nil, err
	}

	for _, queue := range queues {
		for _, rule := range rules {
			if !rule.isValid(queue) {
				invalidUpdates = append(invalidUpdates, queue)
				break
			}
		}
	}

	return invalidUpdates, rules, nil
}

func findInvalidRules(queue []int, rules []rule) []rule {
	results := []rule{}

	for _, rule := range rules {
		if rule.isValid(queue) {
			results = append(results, rule)
		}
	}

	return results
}

func repairFaultyUpdates(queues [][]int, rules []rule) [][]int {
	for i := range queues {
		timeoutSecs := 5
		timeout := time.Now().Add(time.Second * time.Duration(timeoutSecs))

		for {
			assert.Assert(time.Now().Before(timeout), "repair cannot take longer than %d seconds", timeoutSecs)

			rules := findInvalidRules(queues[i], rules)
			if len(rules) == 0 {
				break
			}

			// This doesn't work, should probably go for a tree
			for _, rule := range rules {
				var beforeIndex int = -1
				var afterIndex int = -1

				for j, num := range queues[i] {
					if num == rule.before {
						beforeIndex = j
					} else if num == rule.after {
						afterIndex = j
					}

					if beforeIndex != -1 && afterIndex != -1 {
						break
					}
				}

				if beforeIndex == -1 || afterIndex == -1 {
					continue
				}

				num := queues[i][beforeIndex]
				queues[i] = slices.Delete(queues[i], beforeIndex, beforeIndex+1)
				queues[i] = slices.Insert(queues[i], afterIndex, num)
			}
		}
	}

	return queues
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
	input, err := io.ReadStr("internal/advents/a5/input.txt")
	if err != nil {
		panic(err)
	}

	invalidQueues, rules, err := getIncorrectlyOrderedUpdates(input)
	if err != nil {
		panic(err)
	}

	repairedQueues := repairFaultyUpdates(invalidQueues, rules)

	sum := sumMiddleNumber(repairedQueues)
	fmt.Println(sum)
}

var _ structure.Challenge = day5{}

func New() structure.Challenge {
	return day5{}
}
