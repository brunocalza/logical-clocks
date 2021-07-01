package examples

import "github.com/brunocalza/logical-clocks/lamport"

// Example5 ...
func Example5() []func(*lamport.Clock) {
	A := func(clock *lamport.Clock) {
		clock.Send(2)
		clock.Send(1)
		clock.Recv(1)
	}

	B := func(clock *lamport.Clock) {
		clock.Recv(0)
		clock.Send(2)
		clock.Send(0)
	}

	C := func(clock *lamport.Clock) {
		clock.Recv(1)
		clock.Recv(0)
	}

	return []func(*lamport.Clock){A, B, C}
}
