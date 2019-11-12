package kuncore

import (
	"sync"
)

type MutexMapCC struct {
	c map[string]*int64
	m sync.RWMutex
}

func NewMutexMapCC() *MutexMapCC {
	return &MutexMapCC{
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
	*cc.c[key]++
}

func (cc *MutexMapCC) Decrease(key string) {
	cc.m.Lock()
	defer cc.m.Unlock()
	c := cc.c[key]
	if c == nil {
		return
	}
	if *cc.c[key] <= 0 {
		*cc.c[key] = 0
	} else {
		*cc.c[key]--
	}
}

func (cc *MutexMapCC) Count(key string) int64 {
	cc.m.RLock()
	defer cc.m.RUnlock()
	c := cc.c[key]
	if c == nil {
		return 0
	}
	return *c
}
