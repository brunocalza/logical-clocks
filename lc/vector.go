package lc

type VectorClock struct {
	ID        Identifier
	Channels  map[ChannelKey](chan []Timestamp)
	Events    chan Event
	Timestamp []Timestamp
}

func (clock *VectorClock) Local() {
	clock.Timestamp[clock.ID]++
	timestampCopy := append(make([]Timestamp, 0, len(clock.Timestamp)), clock.Timestamp...)
	clock.Events <- Event{Vector, Local, timestampCopy, clock.ID, nil, nil}
}

func (clock *VectorClock) Recv(source Identifier) {
	receivedTimestamp := <-clock.Channels[NewChannelKey(source, clock.ID)]
	clock.Timestamp = maxVector(clock.Timestamp, receivedTimestamp)
	clock.Timestamp[clock.ID]++
	timestampCopy := append(make([]Timestamp, 0, len(clock.Timestamp)), clock.Timestamp...)
	clock.Events <- Event{Vector, Received, timestampCopy, clock.ID, &source, &clock.ID}
}

func (clock *VectorClock) Send(destination Identifier) {
	clock.Timestamp[clock.ID]++
	clock.Channels[NewChannelKey(clock.ID, destination)] <- clock.Timestamp
	timestampCopy := append(make([]Timestamp, 0, len(clock.Timestamp)), clock.Timestamp...)
	clock.Events <- Event{Vector, Sent, timestampCopy, clock.ID, &clock.ID, &destination}
}

func maxVector(a []Timestamp, b []Timestamp) []Timestamp {
	result := make([]Timestamp, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = maxTimestamp(a[i], b[i])
	}
	return result
}
