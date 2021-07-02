package lc

type LamportClock struct {
	ID        Identifier
	Channels  map[ChannelKey](chan Timestamp)
	Events    chan Event
	Timestamp Timestamp
}

func (clock *LamportClock) Local() {
	clock.Timestamp++
	clock.Events <- Event{Lamport, Local, clock.Timestamp, clock.ID, nil, nil}
}

func (clock *LamportClock) Recv(source Identifier) {
	receivedTimestamp := <-clock.Channels[NewChannelKey(source, clock.ID)]
	clock.Timestamp = maxTimestamp(clock.Timestamp, receivedTimestamp) + 1
	clock.Events <- Event{Lamport, Received, clock.Timestamp, clock.ID, &source, &clock.ID}
}

func (clock *LamportClock) Send(destination Identifier) {
	clock.Timestamp++
	clock.Channels[NewChannelKey(clock.ID, destination)] <- clock.Timestamp
	clock.Events <- Event{Lamport, Sent, clock.Timestamp, clock.ID, &clock.ID, &destination}
}

func maxTimestamp(a Timestamp, b Timestamp) Timestamp {
	if a > b {
		return a

	}
	return b
}
