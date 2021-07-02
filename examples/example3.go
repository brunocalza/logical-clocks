package examples

import "github.com/brunocalza/logical-clocks/lc"

// Example3 ...
func Example3() Example {
	A := func(clock lc.Clock) {
		clock.Send(1)
		clock.Recv(1)
		clock.Local()
		clock.Recv(1)
	}

	B := func(clock lc.Clock) {
		clock.Send(0)
		clock.Send(2)
		clock.Recv(0)
		clock.Local()
		clock.Send(2)
		clock.Send(0)
		clock.Local()
		clock.Recv(2)
	}

	C := func(clock lc.Clock) {
		clock.Local()
		clock.Send(1)
		clock.Recv(1)
		clock.Recv(1)
	}

	return Example{lc.Lamport, []func(lc.Clock){A, B, C}}
}
