kuncore
===========

```bash
go test -bench=.
```

```bash
Don't use NoSafeCC in production
goos: darwin
goarch: amd64
pkg: github/hefangshi/kuncore
BenchmarkIncrease/SyncMapCC-4 	10000000	       117 ns/op
BenchmarkIncrease/LockfreeMapCC-4         	10000000	       144 ns/op
BenchmarkIncrease/MutexMapCC-4            	20000000	       114 ns/op
BenchmarkIncrease/ChannelCC-4             	 5000000	       323 ns/op
BenchmarkIncrease/NoSafeCC-4              	50000000	        33.6 ns/op
BenchmarkDecrease/SyncMapCC-4             	20000000	       111 ns/op
BenchmarkDecrease/LockfreeMapCC-4         	10000000	       144 ns/op
BenchmarkDecrease/MutexMapCC-4            	20000000	        98.0 ns/op
BenchmarkDecrease/ChannelCC-4             	 5000000	       335 ns/op
BenchmarkDecrease/NoSafeCC-4              	50000000	        33.5 ns/op
BenchmarkParallelIncrease/SyncMapCC-4     	 1000000	      6316 ns/op
BenchmarkParallelIncrease/LockfreeMapCC-4 	 5000000	       429 ns/op
BenchmarkParallelIncrease/MutexMapCC-4    	 3000000	       476 ns/op
BenchmarkParallelIncrease/ChannelCC-4     	 2000000	       724 ns/op
BenchmarkParallelDncrease/SyncMapCC-4     	 5000000	       354 ns/op
BenchmarkParallelDncrease/LockfreeMapCC-4 	 5000000	       347 ns/op
BenchmarkParallelDncrease/MutexMapCC-4    	 3000000	       505 ns/op
BenchmarkParallelDncrease/ChannelCC-4     	 3000000	       831 ns/op
PASS
ok  	github/hefangshi/kuncore	49.215s
```