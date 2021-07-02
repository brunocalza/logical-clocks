package examples

import "github.com/brunocalza/logical-clocks/lc"

// Example7 ...
func Example7() Example {
	A := func(clock lc.Clock) {
		clock.Send(1)
		clock.Send(2)
		clock.Recv(1)
		clock.Recv(2)
	}

	B := func(clock lc.Clock) {
		clock.Send(0)
		clock.Send(2)
		clock.Recv(0)
		clock.Recv(2)
	}

	C := func(clock lc.Clock) {
		clock.Send(0)
		clock.Send(1)
		clock.Recv(0)
		clock.Recv(1)
	}

	return Example{lc.Vector, []func(lc.Clock){A, B, C}}
}
