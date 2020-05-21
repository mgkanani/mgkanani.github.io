# Changes
In quest to make goroutines event-driven, I explored `runtime` package.
To run tests/benchmarks, add below lines in 
[runtime/stubs.go](https://github.com/golang/go/blob/release-branch.go1.13/src/runtime/stubs.go#L19)

```go
func getG() *g {
	return getg()
}
```

We can solve the problem by using channels. A shared channel in case of batching can lead to performance bottlenecks.
E.g. Producer-Consumer like single-many/many-many/many-single
Since the same channel is shared across these go-routines, `lock` field in `hchan` struct will make other routines to wait 