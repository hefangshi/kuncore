package kuncore

import (
	"sync"
	"testing"
)

var keys = []string{"abc1", "abc2", "abc3", "abc4", "abc5", "abc6"}
var sm = SyncMapCC{}
var hm = NewLockfreeMapCC()
var mm = NewMutexMapCC()
var ns = NewNoSafeCC()

func incr(cc ConcurrencyCounter, b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.Increase(keys[i%len(keys)])
	}
}

func decr(cc ConcurrencyCounter, b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.Decrease(keys[i%len(keys)])
	}
}

func parallelIncr(cc ConcurrencyCounter, b *testing.B) {
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

func parallelDecr(cc ConcurrencyCounter, b *testing.B) {
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

func BenchmarkSyncMapCC_Increase(b *testing.B) {
	incr(&sm, b)
}

func BenchmarkSyncMapCC_Decrease(b *testing.B) {
	decr(&sm, b)
}

func BenchmarkLockfreeMapCC_Increase(b *testing.B) {
	incr(&hm, b)
}

func BenchmarkLockfreeMapCC_Decrease(b *testing.B) {
	decr(&hm, b)
}

func BenchmarkMutextMapCC_Increase(b *testing.B) {
	incr(&mm, b)
}

func BenchmarkMutextMapCC_Decrease(b *testing.B) {
	decr(&mm, b)
}

func BenchmarkNoSafeCC_Increase(b *testing.B) {
	incr(&ns, b)
}

func BenchmarkNoSafeCC_Decrease(b *testing.B) {
	decr(&ns, b)
}

func BenchmarkParallelSyncMapCC_Increase(b *testing.B) {
	parallelIncr(&sm, b)
}

func BenchmarkParallelSyncMapCC_Decrease(b *testing.B) {
	parallelDecr(&sm, b)
}

func BenchmarkParallelLockfreeMapCC_Increase(b *testing.B) {
	parallelIncr(&hm, b)
}

func BenchmarkParallelLockfreeMapCC_Decrease(b *testing.B) {
	parallelDecr(&hm, b)
}

func BenchmarkParallelMutextMapCC_Increase(b *testing.B) {
	parallelIncr(&mm, b)
}

func BenchmarkParallelMutextMapCC_Decrease(b *testing.B) {
	parallelDecr(&mm, b)
}

//func BenchmarkNoSafeCC_Increase(b *testing.B) {
//	incr(&ns, b)
//}
//
//func BenchmarkNoSafeCC_Decrease(b *testing.B) {
//	decr(&ns, b)
//}
