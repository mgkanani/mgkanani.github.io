package runtime

import (
	"testing"
	"time"
)

func BenchmarkCustomNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		now()
	}
}

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func BenchmarkChan(b *testing.B) {
	ch := make(chan int)
	N := b.N
	go func(consumer <-chan int) {
		for i := 0; i < N; i++ {
			<-consumer
		}
	}(ch)
	for i := 0; i < N; i++ {
		ch <- i
	}
	close(ch)
}

func BenchmarkCustom(b *testing.B) {
	// number of the this benchmark can be misleading.
	// Remember, in this benchmark both go-routines running one-by-one
	// in channel based communication, sender underneath calls goready
	// while receiver calls gopark
	// For comparing with chan-based approach, half the ops/ns number.
	N := b.N
	var goRtn *g
	mainRtn := getG()
	lck := &mutex{}

	go func(max int) {
		goRtn = getG()
		for i := 0; i < max; i++ {
			lock(lck)
			goready(mainRtn, 0)
			gopark(lck, waitReasonChanReceive, traceEvGoBlockRecv, 0)
		}
	}(N)

	lock(lck)
	gopark(lck, waitReasonChanReceive, traceEvGoBlockRecv, 0)
	for i := 0; i < N-1; i++ {
		lock(lck)
		goready(goRtn, 0)
		gopark(lck, waitReasonChanReceive, traceEvGoBlockRecv, 0)
	}
}
