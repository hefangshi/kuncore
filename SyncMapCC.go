package kuncore

import (
	"sync"
	"sync/atomic"
)

type SyncMapCC struct {
	sm sync.Map
}

func (cc *SyncMapCC) Increase(key string) {
	var k int64
	c, _ := cc.sm.LoadOrStore(key, &k)
	i := c.(*int64)
	atomic.AddInt64(i, 1)
}

func (cc *SyncMapCC) Decrease(key string) {
	var k int64
	c, _ := cc.sm.LoadOrStore(key, &k)
	i := c.(*int64)
	if atomic.LoadInt64(i) < 0 {
		atomic.StoreInt64(i, 0)
	} else {
		atomic.AddInt64(i, -1)
	}
}

func (cc *SyncMapCC) Count(key string) int64 {
	c, ok := cc.sm.Load(key)
	if ok {
		i := c.(*int64)
		return atomic.LoadInt64(i)
	}
	return 0
}
