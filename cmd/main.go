package main

import (
	"dekamik/aoc_2024/internal/advents/a5"
	"dekamik/aoc_2024/internal/logging"
	"log/slog"
)

func main() {
	logging.NewDefault(slog.LevelInfo)

	command := a5.New()
	command.ExecutePart1()
	command.ExecutePart2()
}
