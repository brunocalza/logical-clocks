package main

import (
	"os"
	"sync"

	"github.com/brunocalza/logical-clocks/examples"
	"github.com/brunocalza/logical-clocks/lc"
)

func main() {
	exampleName := os.Args[1]
	example, ok := examples.List()[exampleName]
	if !ok {
		panic("you have to add your example to examples list in examples.go")
	}
	exampleRun := newExampleRun(example())

	exampleRun.execute()
	exampleRun.startLogging()
	exampleRun.processesWG.Wait()

	exampleRun.closeChannels()
	exampleRun.loggingWG.Wait()
}

type exampleRun struct {
	example         examples.Example
	lamportChannels map[lc.ChannelKey](chan lc.Timestamp)
	vectorChannels  map[lc.ChannelKey](chan []lc.Timestamp)
	events          chan lc.Event
	processesWG     *sync.WaitGroup
	loggingWG       *sync.WaitGroup
}

func newExampleRun(example examples.Example) *exampleRun {
	run := &exampleRun{}
	run.example = example
	run.events = make(chan lc.Event)
	if example.Clock == lc.Lamport {
		run.lamportChannels = run.generateChannelsForLamport(run.example.Processes)
	} else {
		run.vectorChannels = run.generateChannelsForVector(run.example.Processes)
	}
	return run
}

// execute all processes concurrently. Each process has its own clock
func (run *exampleRun) execute() {
	run.processesWG = &sync.WaitGroup{}
	for i := 0; i < len(run.example.Processes); i++ {
		run.processesWG.Add(1)
		go func(i int) {
			clock := run.newClock(lc.Identifier(i))
			run.example.Processes[i](clock)
			run.processesWG.Done()
		}(i)
	}
}

// log all events
func (run *exampleRun) startLogging() {
	run.loggingWG = &sync.WaitGroup{}
	run.loggingWG.Add(1)
	go func() {
		for event := range run.events {
			event.Log()
		}
		run.loggingWG.Done()
	}()
}

func (run *exampleRun) closeChannels() {
	if run.example.Clock == lc.Lamport {
		for _, channel := range run.lamportChannels {
			close(channel)
		}
	} else {
		for _, channel := range run.vectorChannels {
			close(channel)
		}
	}

	close(run.events)
}

// NewClock creates the clock for the example
func (run *exampleRun) newClock(id lc.Identifier) lc.Clock {
	if run.example.Clock == lc.Lamport {
		return &lc.LamportClock{lc.Identifier(id), run.lamportChannels, run.events, 0}
	}

	return &lc.VectorClock{lc.Identifier(id), run.vectorChannels, run.events, make([]lc.Timestamp, len(run.example.Processes))}
}

func (example *exampleRun) generateChannelsForLamport(processes []func(lc.Clock)) map[lc.ChannelKey](chan lc.Timestamp) {
	channels := make(map[lc.ChannelKey](chan lc.Timestamp))
	for _, combination := range combinations(len(processes)) {
		source := lc.Identifier(combination[0])
		destination := lc.Identifier(combination[1])

		channels[lc.NewChannelKey(source, destination)] = make(chan lc.Timestamp, 100)
		channels[lc.NewChannelKey(destination, source)] = make(chan lc.Timestamp, 100)
	}
	return channels
}

func (example *exampleRun) generateChannelsForVector(processes []func(lc.Clock)) map[lc.ChannelKey](chan []lc.Timestamp) {
	channels := make(map[lc.ChannelKey](chan []lc.Timestamp))
	for _, combination := range combinations(len(processes)) {
		source := lc.Identifier(combination[0])
		destination := lc.Identifier(combination[1])

		channels[lc.NewChannelKey(source, destination)] = make(chan []lc.Timestamp, 100)
		channels[lc.NewChannelKey(destination, source)] = make(chan []lc.Timestamp, 100)
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
