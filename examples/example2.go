package examples

import "github.com/brunocalza/logical-clocks/lamport"

// Example2 ...
func Example2() []func(*lamport.Clock) {
	A := func(clock *lamport.Clock) {
		clock.Send(1)
		clock.Send(2)
		clock.Recv(1)
		clock.Recv(2)
	}

	B := func(clock *lamport.Clock) {
		clock.Send(0)
		clock.Send(2)
		clock.Recv(0)
		clock.Recv(2)
	}

	C := func(clock *lamport.Clock) {
		clock.Send(0)
		clock.Send(1)
		clock.Recv(0)
		clock.Recv(1)
	}

	return []func(*lamport.Clock){A, B, C}
}
