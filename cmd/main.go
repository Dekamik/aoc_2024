package main

import (
	"dekamik/aoc_2024/internal/advents/a6"
	"dekamik/aoc_2024/internal/logging"
	"log/slog"
)

func main() {
    logging.NewDefault(slog.LevelDebug)

	command := a6.New()
	command.ExecutePart1()
	command.ExecutePart2()
}
