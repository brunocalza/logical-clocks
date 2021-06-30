package examples

import "github.com/brunocalza/logical-clocks/lamport"

// Example1 ...
func Example1() []func(*lamport.Clock) {
	A := func(clock *lamport.Clock) {
		clock.Send(1)
		clock.Recv(1)
	}

	B := func(clock *lamport.Clock) {
		clock.Recv(0)
		clock.Send(0)
	}

	return []func(*lamport.Clock){A, B}
}
