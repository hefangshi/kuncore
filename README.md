kuncore
===========

```bash
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github/hefangshi/kuncore -bench ^(BenchmarkIncrease|BenchmarkDecrease|BenchmarkParallelIncrease|BenchmarkParallelDecrease|BenchmarkCount|BenchmarkParallelCount)$

Don't use NoSafeCC in production
goos: linux
goarch: amd64
pkg: github/hefangshi/kuncore
BenchmarkIncrease/SyncMapCC-8 	11026066	       101 ns/op	      24 B/op	       2 allocs/op
BenchmarkIncrease/LockfreeMapCC-8         	12416960	        99.7 ns/op	      40 B/op	       3 allocs/op
BenchmarkIncrease/MutexMapCC-8            	16021575	        75.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkIncrease/ChannelCC-8             	 7898113	       165 ns/op	       0 B/op	       0 allocs/op
BenchmarkIncrease/CMapCC-8                	23711197	        51.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIncrease/NoSafeCC-8              	33897346	        35.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecrease/SyncMapCC-8             	12968906	        95.4 ns/op	      24 B/op	       2 allocs/op
BenchmarkDecrease/LockfreeMapCC-8         	12227803	        99.1 ns/op	      40 B/op	       3 allocs/op
BenchmarkDecrease/MutexMapCC-8            	13861774	        87.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecrease/ChannelCC-8             	 8006564	       169 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecrease/CMapCC-8                	23812358	        51.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecrease/NoSafeCC-8              	33597446	        35.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkParallelIncrease/SyncMapCC-8     	 5745777	       208 ns/op	      24 B/op	       2 allocs/op
BenchmarkParallelIncrease/LockfreeMapCC-8 	 5711701	       913 ns/op	     168 B/op	       3 allocs/op
BenchmarkParallelIncrease/MutexMapCC-8    	 2059954	       573 ns/op	      16 B/op	       0 allocs/op
BenchmarkParallelIncrease/ChannelCC-8     	 1949754	       615 ns/op	      35 B/op	       0 allocs/op
BenchmarkParallelIncrease/CMapCC-8        	 5793406	       206 ns/op	       0 B/op	       0 allocs/op
BenchmarkParallelDecrease/SyncMapCC-8     	 5765322	       215 ns/op	      24 B/op	       2 allocs/op
BenchmarkParallelDecrease/LockfreeMapCC-8 	 5742038	       211 ns/op	      40 B/op	       3 allocs/op
BenchmarkParallelDecrease/MutexMapCC-8    	 1891052	       633 ns/op	      24 B/op	       0 allocs/op
BenchmarkParallelDecrease/ChannelCC-8     	 1829156	       643 ns/op	      41 B/op	       0 allocs/op
BenchmarkParallelDecrease/CMapCC-8        	 5829006	       208 ns/op	       0 B/op	       0 allocs/op
BenchmarkCount/SyncMapCC-8                	 7595704	       158 ns/op	      48 B/op	       4 allocs/op
BenchmarkCount/LockfreeMapCC-8            	 6145872	       196 ns/op	      80 B/op	       6 allocs/op
BenchmarkCount/MutexMapCC-8               	 8137027	       149 ns/op	       0 B/op	       0 allocs/op
BenchmarkCount/ChannelCC-8                	 3728235	       350 ns/op	       0 B/op	       0 allocs/op
BenchmarkCount/CMapCC-8                   	12748054	        96.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkCount/NoSafeCC-8                 	24499795	        49.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkParallelCount/SyncMapCC-8        	 5934482	       201 ns/op	      48 B/op	       4 allocs/op
BenchmarkParallelCount/LockfreeMapCC-8    	 5957728	       201 ns/op	      80 B/op	       6 allocs/op
BenchmarkParallelCount/MutexMapCC-8       	 1000000	      1034 ns/op	      46 B/op	       0 allocs/op
BenchmarkParallelCount/ChannelCC-8        	 1289914	      1109 ns/op	      48 B/op	       0 allocs/op
BenchmarkParallelCount/CMapCC-8           	 5803352	       207 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github/hefangshi/kuncore	67.756s
Success: Benchmarks passed.
```