package lc

type LamportClock struct {
	ID        Identifier
	Channels  map[ChannelKey](chan Timestamp)
	Events    chan Event
	Timestamp Timestamp
}

func (clock *LamportClock) Local() {
	clock.Timestamp++
	clock.Events <- Event{Local, clock.Timestamp, clock.ID, nil, nil}
}

func (clock *LamportClock) Recv(source Identifier) {
	receivedTimestamp := <-clock.Channels[NewChannelKey(source, clock.ID)]
	clock.Timestamp = max(clock.Timestamp, receivedTimestamp) + 1
	clock.Events <- Event{Received, clock.Timestamp, clock.ID, &source, &clock.ID}
}

func (clock *LamportClock) Send(destination Identifier) {
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
