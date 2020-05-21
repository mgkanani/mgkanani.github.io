package runtime

import (
	_ "unsafe"
)

// Event types in the trace, args are given in square brackets.
// go tool trace will behave as per data-values inside square brackets.
const (
	waitReasonChanReceive = 14
	waitReasonChanSend    = 15
	traceEvGoBlockSend    = 22 // goroutine blocks on chan send [timestamp, stack]
	traceEvGoBlockRecv    = 23 // goroutine blocks on chan recv [timestamp, stack]
)

type g struct{}

type mutex struct {
	key uintptr
}

//go:linkname getG runtime.getG
func getG() *g

//go:linkname lock runtime.lock
func lock(m *mutex)

//go:linkname unlock runtime.unlock
func unlock(m *mutex)

//go:linkname gopark runtime.goparkunlock
func gopark(lock *mutex, reason uint8, traceEv int, traceSkip int)

//go:linkname goready runtime.goready
func goready(gp *g, traceskip int)

// can be used for metrics recording
//go:linkname now runtime.nanotime
func now() int64
