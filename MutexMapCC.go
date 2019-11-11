package main

import (
	"sync"
	"sync/atomic"
)

type MutexMapCC struct {
	c map[string]*int64
	m sync.RWMutex
}

func NewMutexMapCC() MutexMapCC {
	return MutexMapCC{
		c: make(map[string]*int64),
	}
}

func (cc *MutexMapCC) Increase(key string) {
	cc.m.Lock()
	defer cc.m.Unlock()
	if cc.c[key] == nil {
		var k int64
		cc.c[key] = &k
	}
	c := cc.c[key]
	atomic.AddInt64(c, 1)
}

func (cc *MutexMapCC) Decrease(key string) {
	cc.m.Lock()
	defer cc.m.Unlock()
	c := cc.c[key]
	if c == nil {
		return
	}
	if atomic.LoadInt64(c) < 0 {
		atomic.StoreInt64(c, 0)
	} else {
		atomic.AddInt64(c, -1)
	}
}

func (cc *MutexMapCC) Count(key string) int64 {
	cc.m.RLock()
	defer cc.m.RUnlock()
	c := cc.c[key]
	if c == nil {
		return 0
	}
	return atomic.LoadInt64(c)
}
