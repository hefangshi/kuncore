package kuncore

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

var keys []string

var benchmarks = []struct {
	name string
	cc   ConcurrencyCounter
}{
	{"SyncMapCC", NewSyncMapCC()},
	{"LockfreeMapCC", NewLockfreeMapCC()},
	{"MutexMapCC", NewMutexMapCC()},
	{"ChannelCC", NewChannelCC()},
	{"CMapCC", NewCMapCC()},
}

var benchmarksWithNoSafeMapCC = append(benchmarks, struct {
	name string
	cc   ConcurrencyCounter
}{"NoSafeCC", NewNoSafeCC()})

func init() {
	rand.Seed(time.Now().UnixNano())
	var initKeys = []string{"abc1", "abc2", "abc3", "abc4", "abc5", "abc6"}
	keys = make([]string, 10e6)
	for i := 0; i < len(keys); i++ {
		keys[i] = initKeys[rand.Int()%len(initKeys)]
	}
}

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

func count(keys []string, cc ConcurrencyCounter, b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.Increase(keys[i%len(keys)])
		cc.Decrease(keys[i%len(keys)])
	}
}

func parallelCount(keys []string, cc ConcurrencyCounter, b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			cc.Increase(keys[i%len(keys)])
			cc.Decrease(keys[i%len(keys)])
			wg.Done()
		}(i)
	}
	wg.Wait()
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

func BenchmarkCount(b *testing.B) {
	for _, bm := range benchmarksWithNoSafeMapCC {
		b.Run(bm.name, func(b *testing.B) {
			count(keys, bm.cc, b)
		})
	}
}

func BenchmarkParallelCount(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			parallelCount(keys, bm.cc, b)
		})
	}
}
