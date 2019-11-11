kuncore
===========

```bash
go test -bench=.
```

```bash
goos: darwin
goarch: amd64
pkg: github/hefangshi/kuncore
BenchmarkSyncMapCC_Increase-4       	20000000	       112 ns/op
BenchmarkSyncMapCC_Decrease-4       	10000000	       130 ns/op
BenchmarkLockfreeMapCC_Increase-4   	10000000	       167 ns/op
BenchmarkLockfreeMapCC_Decrease-4   	10000000	       150 ns/op
BenchmarkMutextMapCC_Increase-4     	10000000	       117 ns/op
BenchmarkMutextMapCC_Decrease-4     	20000000	        98.2 ns/op
BenchmarkNoSafeCC_Increase-4        	50000000	        33.2 ns/op
BenchmarkNoSafeCC_Decrease-4        	50000000	        32.8 ns/op
PASS
ok  	github/hefangshi/kuncore	14.938s
```