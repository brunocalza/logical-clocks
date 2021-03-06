# Logical Clocks Visualizer

## Description

This is an educational tool that implements Lamport Clocks and Vector Clocks and helps you visualize a flow of events defined by the user. You can learn the theory about Logical Clocks at [Getting To Know Logical Clocks By Implementing Them](https://brunocalza.me/getting-to-know-logical-clocks-by-implementing-them/). This is highly inspired in [mwhittaker/clocks](https://github.com/mwhittaker/clocks).

## How to use this tool

There are 3 kinds of events `Local`, `Send` and `Recv` that you can produce inside a process. You can define as many processes as you wish. And you can choose the clock implementation. Look at [example3](https://github.com/brunocalza/logical-clocks/blob/main/examples/example3.go).

You can run an example and visualize it by running:

```go
go run cmd/* Example8 | ./plot.py`
```

```go
// Example8 ...
func Example8() Example {
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

    return Example{lc.Vector, []func(lc.Clock){A, B, C}}
}
```

![Execution of Example 8](Example8.png)
