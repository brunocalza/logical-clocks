package examples

import "github.com/brunocalza/logical-clocks/lamport"

// Example4 ...
func Example4() []func(*lamport.Clock) {
	A := func(clock *lamport.Clock) {
		clock.Local()
		clock.Send(1)
		clock.Local()
	}

	B := func(clock *lamport.Clock) {
		clock.Recv(0)
		clock.Send(2)
	}

	C := func(clock *lamport.Clock) {
		clock.Local()
		clock.Recv(1)
	}

	return []func(*lamport.Clock){A, B, C}
}
