package examples

import "github.com/brunocalza/logical-clocks/lc"

type Example struct {
	Clock     lc.ClockType
	Processes []func(lc.Clock)
}

type Examples map[string](func() Example)

func List() Examples {
	return Examples{
		"Example1": Example1,
		"Example2": Example2,
		"Example3": Example3,
		"Example4": Example4,
		"Example5": Example5,
		"Example6": Example6,
		"Example7": Example7,
		"Example8": Example8,
	}
}
