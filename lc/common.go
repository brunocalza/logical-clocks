package lc

import (
	"fmt"
	"strings"
)

type EventType int

const (
	Local EventType = iota
	Sent
	Received
)

func (eventType EventType) toString() string {
	switch eventType {
	case Local:
		return "local"
	case Sent:
		return "sent"
	case Received:
		return "received"
	default:
		return ""
	}
}

type Timestamp uint32
type Identifier uint32

func (id *Identifier) toString() string {
	if id == nil {
		return "NaN"
	}

	return fmt.Sprint(*id)
}

type Event struct {
	Clock       ClockType
	Kind        EventType
	Timestamp   interface{} // can be Timestamp or []Timestamp
	Owner       Identifier
	Source      *Identifier
	Destination *Identifier
}

func (event *Event) Log() {
	if event.Clock == Lamport {
		event.logForLamport()
	} else {
		event.logForVector()
	}
}

func (event *Event) logForLamport() {
	fmt.Printf("%s %d %s %s %d\n", event.Kind.toString(), event.Owner, event.Source.toString(), event.Destination.toString(), event.Timestamp.(Timestamp))
}

func (event *Event) logForVector() {
	timestamps := event.Timestamp.([]Timestamp)
	timestampsStr := make([]string, len(timestamps))
	for i, t := range timestamps {
		timestampsStr[i] = fmt.Sprint(t)
	}

	fmt.Printf("%s %d %s %s %s\n", event.Kind.toString(), event.Owner, event.Source.toString(), event.Destination.toString(), strings.Join(timestampsStr, " "))
}

type ChannelKey string

func NewChannelKey(source Identifier, destination Identifier) ChannelKey {
	return ChannelKey(fmt.Sprint(source) + "_" + fmt.Sprint(destination))
}
