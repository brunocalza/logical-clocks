package lamport

import "fmt"

type EventType int

const (
	Local EventType = iota
	Sent
	Received
)

type Timestamp uint32
type Id uint32

type Event struct {
	Kind        EventType
	Timestamp   Timestamp
	Owner       Id
	Source      *Id
	Destination *Id
}

func (event *Event) Log() {
	switch event.Kind {
	case Local:
		fmt.Printf("local %d NaN NaN %d\n", event.Owner, event.Timestamp)
	case Sent:
		fmt.Printf("sent %d %d %d %d\n", event.Owner, *event.Source, *event.Destination, event.Timestamp)
	case Received:
		fmt.Printf("received %d %d %d %d\n", event.Owner, *event.Source, *event.Destination, event.Timestamp)
	}
}

type ChannelKey string

func NewChannelKey(source Id, destination Id) ChannelKey {
	return ChannelKey(fmt.Sprint(source) + "_" + fmt.Sprint(destination))
}

type Clock struct {
	ID        Id
	Channels  map[ChannelKey](chan Timestamp)
	Events    chan Event
	Timestamp Timestamp
}

func (clock *Clock) Local() {
	clock.Timestamp++
	clock.Events <- Event{Local, clock.Timestamp, clock.ID, nil, nil}
}

func (clock *Clock) Recv(source Id) {
	receivedTimestamp := <-clock.Channels[NewChannelKey(source, clock.ID)]
	clock.Timestamp = max(clock.Timestamp, receivedTimestamp) + 1
	clock.Events <- Event{Received, clock.Timestamp, clock.ID, &source, &clock.ID}
}

func (clock *Clock) Send(destination Id) {
	clock.Timestamp++
	clock.Channels[NewChannelKey(clock.ID, destination)] <- clock.Timestamp
	clock.Events <- Event{Sent, clock.Timestamp, clock.ID, &clock.ID, &destination}
}

func max(a Timestamp, b Timestamp) Timestamp {
	if a > b {
		return a

	}
	return b
}
