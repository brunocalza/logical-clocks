package examples

import "github.com/brunocalza/logical-clocks/lamport"

// Example3 ...
func Example3() []func(*lamport.Clock) {
	A := func(clock *lamport.Clock) {
		clock.Send(1)
		clock.Recv(1)
		clock.Local()
		clock.Recv(1)
	}

	B := func(clock *lamport.Clock) {
		clock.Send(0)
		clock.Send(2)
		clock.Recv(0)
		clock.Local()
		clock.Send(2)
		clock.Send(0)
		clock.Local()
		clock.Recv(2)
	}

	C := func(clock *lamport.Clock) {
		clock.Local()
		clock.Send(1)
		clock.Recv(1)
		clock.Recv(1)
	}

	return []func(*lamport.Clock){A, B, C}
}
