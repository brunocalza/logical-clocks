package main

import (
	"os"
	"sync"

	"github.com/brunocalza/logical-clocks/examples"
	"github.com/brunocalza/logical-clocks/lc"
)

func main() {
	exampleName := os.Args[1]
	example := examples.List()[exampleName]()
	channels := generateChannels(example.Processes)
	events := make(chan lc.Event)

	// Execute all processes concurrently. Each process has its own clock
	wg1 := &sync.WaitGroup{}
	for i := 0; i < len(example.Processes); i++ {
		wg1.Add(1)
		go func(i int) {
			clock := lc.NewClock(example.Clock, lc.Identifier(i), channels, events)
			example.Processes[i](clock)
			wg1.Done()
		}(i)
	}

	// Log all events
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		for event := range events {
			event.Log()
		}
		wg2.Done()
	}()
	wg1.Wait()

	for _, channel := range channels {
		close(channel)
	}
	close(events)
	wg2.Wait()
}

// Generates all combinations of channels necessary for communication among processes
func generateChannels(processes []func(lc.Clock)) map[lc.ChannelKey](chan lc.Timestamp) {
	channels := make(map[lc.ChannelKey](chan lc.Timestamp))
	for _, combination := range combinations(len(processes)) {
		source := lc.Identifier(combination[0])
		destination := lc.Identifier(combination[1])

		channels[lc.NewChannelKey(source, destination)] = make(chan lc.Timestamp, 100)
		channels[lc.NewChannelKey(destination, source)] = make(chan lc.Timestamp, 100)
	}
	return channels
}

// Receives a number n, and generates a list of pairs that can be formed from [0, n).
// Example: n = 2, [[0, 1], [0, 2], [1, 2]]
func combinations(n int) [][]int {
	combs := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				combs = append(combs, []int{i, j})
			}
		}
	}
	return combs
}
