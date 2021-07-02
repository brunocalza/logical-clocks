package lc

type ClockType int

const (
	Lamport ClockType = iota
	Vector
)

type Clock interface {
	Local()
	Send(Identifier)
	Recv(Identifier)
}
