package io

import (
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
        return nil, err
	}

	lines := strings.Split(string(data), "\n")

    return lines, nil
}
