package examples

import "github.com/brunocalza/logical-clocks/lamport"

type Examples map[string](func() []func(*lamport.Clock))

func List() Examples {
	return Examples{
		"Example1": Example1,
		"Example2": Example2,
		"Example3": Example3,
	}
}
