package lc

type ClockType int

const (
	Lamport ClockType = iota
	Vector
)

type Clock interface {
	Local()
	Send(Identifier)
	Recv(Identifier)
}

func NewClock(clock ClockType, id Identifier, channels map[ChannelKey]chan Timestamp, events chan Event) Clock {
	if clock == Lamport {
		return &LamportClock{Identifier(id), channels, events, 0}
	}

	return nil
}
