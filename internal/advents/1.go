package advents

import "dekamik/aoc_2024/internal/command"

type day1 struct {
}

// Execute implements internal.Command.
func (d day1) Execute() {
	panic("unimplemented")
}

var _ internal.Command = day1{}
