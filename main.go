package main

import (
	"fmt"
	"sync"
)

type EventType int

const (
	Local EventType = iota
	Sent
	Received
)

type Timestamp uint32
type Id uint32

type Event struct {
	kind        EventType
	timestamp   Timestamp
	owner       Id
	source      *Id
	destination *Id
}

func (event *Event) Log() {
	switch event.kind {
	case Local:
		fmt.Printf("local %d NaN NaN %d\n", event.owner, event.timestamp)
	case Sent:
		fmt.Printf("sent %d %d %d %d\n", event.owner, *event.source, *event.destination, event.timestamp)
	case Received:
		fmt.Printf("received %d %d %d %d\n", event.owner, *event.source, *event.destination, event.timestamp)
	}
}

type ChannelKey string

func NewChannelKey(source Id, destination Id) ChannelKey {
	return ChannelKey(fmt.Sprint(source) + "_" + fmt.Sprint(destination))
}

type Clock struct {
	id        Id
	channels  map[ChannelKey](chan Timestamp)
	events    chan Event
	timestamp Timestamp
}

func (clock *Clock) Local() {
	clock.timestamp++
	clock.events <- Event{Local, clock.timestamp, clock.id, nil, nil}
}

func (clock *Clock) Recv(source Id) {
	receivedTimestamp := <-clock.channels[NewChannelKey(source, clock.id)]
	clock.timestamp = max(clock.timestamp, receivedTimestamp) + 1
	clock.events <- Event{Received, clock.timestamp, clock.id, &source, &clock.id}
}

func (clock *Clock) Send(destination Id) {
	clock.timestamp++
	clock.channels[NewChannelKey(clock.id, destination)] <- clock.timestamp
	clock.events <- Event{Sent, clock.timestamp, clock.id, &clock.id, &destination}
}

func main() {
	p0 := func(clock *Clock) {
		clock.Send(1) // 0_1
		clock.Recv(1)
		clock.Local()
		clock.Recv(1)
	}

	p1 := func(clock *Clock) {
		clock.Send(0) // 1_0
		clock.Send(2) // 1_2
		clock.Recv(0)
		clock.Local()
		clock.Send(2)
		clock.Send(0)
		clock.Local()
		clock.Recv(2)
	}

	p2 := func(clock *Clock) {
		clock.Local()
		clock.Send(1) // 2_1
		clock.Recv(1)
		clock.Recv(1)
	}

	processes := []func(*Clock){p0, p1, p2}

	channels := make(map[ChannelKey](chan Timestamp))
	for _, combination := range combinations([]int{0, 1, 2}, 2) {
		source := Id(combination[0])
		destination := Id(combination[1])

		channels[NewChannelKey(source, destination)] = make(chan Timestamp, 10)
		channels[NewChannelKey(destination, source)] = make(chan Timestamp, 10)
	}
	events := make(chan Event)

	wg := &sync.WaitGroup{}
	for i := 0; i < len(processes); i++ {
		wg.Add(1)
		go func(i int) {
			clock := &Clock{Id(i), channels, events, 0}
			processes[i](clock)
			wg.Done()
		}(i)
	}

	go func() {
		for event := range events {
			event.Log()
		}
	}()
	wg.Wait()

	for _, channel := range channels {
		close(channel)
	}
	close(events)
}
