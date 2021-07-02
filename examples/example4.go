package examples

import "github.com/brunocalza/logical-clocks/lc"

// Example4 ...
func Example4() Example {
	A := func(clock lc.Clock) {
		clock.Local()
		clock.Send(1)
		clock.Local()
	}

	B := func(clock lc.Clock) {
		clock.Recv(0)
		clock.Send(2)
	}

	C := func(clock lc.Clock) {
		clock.Local()
		clock.Recv(1)
	}

	return Example{lc.Lamport, []func(lc.Clock){A, B, C}}
}
