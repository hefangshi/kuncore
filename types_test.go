package main

import (
	"testing"
)

var keys = []string{"abc1", "abc2", "abc3", "abc4", "abc5", "abc6"}
var sm = SyncMapCC{}
var hm = NewLockfreeMapCC()
var mm = NewMutexMapCC()

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
