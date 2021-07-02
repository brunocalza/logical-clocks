package lc

import "fmt"

type EventType int

const (
	Local EventType = iota
	Sent
	Received
)

type Timestamp uint32
type Identifier uint32

type Event struct {
	Kind        EventType
	Timestamp   Timestamp
	Owner       Identifier
	Source      *Identifier
	Destination *Identifier
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

func NewChannelKey(source Identifier, destination Identifier) ChannelKey {
	return ChannelKey(fmt.Sprint(source) + "_" + fmt.Sprint(destination))
}
