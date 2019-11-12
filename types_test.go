package kuncore

import (
	"sync"
	"testing"
)

func incr(keys []string, cc ConcurrencyCounter, b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.Increase(keys[i%len(keys)])
	}
}

func decr(keys []string, cc ConcurrencyCounter, b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.Decrease(keys[i%len(keys)])
	}
}

func parallelIncr(keys []string, cc ConcurrencyCounter, b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			cc.Increase(keys[i%len(keys)])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func parallelDecr(keys []string, cc ConcurrencyCounter, b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			cc.Decrease(keys[i%len(keys)])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

var benchmarks = []struct {
	name string
	cc   ConcurrencyCounter
}{
	{"SyncMapCC", NewSyncMapCC()},
	{"LockfreeMapCC", NewLockfreeMapCC()},
	{"MutexMapCC", NewMutexMapCC()},
	{"ChannelCC", NewChannelCC()},
}

var benchmarksWithNoSafeMapCC = append(benchmarks, struct {
	name string
	cc   ConcurrencyCounter
}{"NoSafeCC", NewNoSafeCC()})

var keys = []string{"abc1", "abc2", "abc3", "abc4", "abc5", "abc6"}

func BenchmarkIncrease(b *testing.B) {
	for _, bm := range benchmarksWithNoSafeMapCC {
		b.Run(bm.name, func(b *testing.B) {
			incr(keys, bm.cc, b)
		})
	}
}

func BenchmarkDecrease(b *testing.B) {
	for _, bm := range benchmarksWithNoSafeMapCC {
		b.Run(bm.name, func(b *testing.B) {
			decr(keys, bm.cc, b)
		})
	}
}

func BenchmarkParallelIncrease(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			parallelIncr(keys, bm.cc, b)
		})
	}
}

func BenchmarkParallelDecrease(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			parallelDecr(keys, bm.cc, b)
		})
	}
}
