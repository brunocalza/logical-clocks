package examples

import "github.com/brunocalza/logical-clocks/lc"

// Example5 ...
func Example5() Example {
	A := func(clock lc.Clock) {
		clock.Send(2)
		clock.Send(1)
		clock.Recv(1)
	}

	B := func(clock lc.Clock) {
		clock.Recv(0)
		clock.Send(2)
		clock.Send(0)
	}

	C := func(clock lc.Clock) {
		clock.Recv(1)
		clock.Recv(0)
	}

	return Example{lc.Lamport, []func(lc.Clock){A, B, C}}
}
