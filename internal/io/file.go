package io

import (
	"os"
	"strings"
)

func ReadStr(path string) (string, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func ReadLines(path string) ([]string, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(buffer), "\n")

	return lines, nil
}
