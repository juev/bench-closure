# bench-closure

Benchmarks for closure and functions.

run it:

```bash
go test -bench=. -benchtime=1000x -benchmem
```

results:

```bash
❯ go test -bench=. -benchtime=1000x -benchmem
goos: darwin
goarch: arm64
pkg: bench-closure
Benchmark_closure-8         1000           5871084 ns/op         3360993 B/op      90003 allocs/op
Benchmark_run-8             1000           5876669 ns/op         3360994 B/op      90003 allocs/op
PASS
ok      bench-closure   11.839s
```

Should we use closures? 
It doesn't matter, since closures have no effect on performance. 
If this makes your code more readable, use closures.