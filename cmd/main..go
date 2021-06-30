package main

import (
	"os"
	"sync"

	"github.com/brunocalza/logical-clocks/examples"
	"github.com/brunocalza/logical-clocks/lamport"
)

func main() {
	exampleName := os.Args[1]
	processes := examples.List()[exampleName]()
	channels := generateChannels(processes)
	events := make(chan lamport.Event)

	// Execute all processes concurrently. Each process has its own clock
	wg1 := &sync.WaitGroup{}
	for i := 0; i < len(processes); i++ {
		wg1.Add(1)
		go func(i int) {
			clock := &lamport.Clock{lamport.Id(i), channels, events, 0}
			processes[i](clock)
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
func generateChannels(processes []func(*lamport.Clock)) map[lamport.ChannelKey](chan lamport.Timestamp) {
	channels := make(map[lamport.ChannelKey](chan lamport.Timestamp))
	for _, combination := range combinations(len(processes)) {
		source := lamport.Id(combination[0])
		destination := lamport.Id(combination[1])

		channels[lamport.NewChannelKey(source, destination)] = make(chan lamport.Timestamp, 100)
		channels[lamport.NewChannelKey(destination, source)] = make(chan lamport.Timestamp, 100)
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
