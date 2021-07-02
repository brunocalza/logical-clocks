package examples

import "github.com/brunocalza/logical-clocks/lc"

// Example1 ...
func Example1() Example {
	A := func(clock lc.Clock) {
		clock.Send(1)
		clock.Recv(1)
	}

	B := func(clock lc.Clock) {
		clock.Recv(0)
		clock.Send(0)
	}

	return Example{lc.Lamport, []func(lc.Clock){A, B}}
}
