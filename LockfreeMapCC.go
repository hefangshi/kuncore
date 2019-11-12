package kuncore

import (
	"github.com/cornelk/hashmap"
	"sync/atomic"
)

type LockfreeMapCC struct {
	hm *hashmap.HashMap
}

func NewLockfreeMapCC() *LockfreeMapCC {
	return &LockfreeMapCC{
		hm: &hashmap.HashMap{},
	}
}

func (cc *LockfreeMapCC) Increase(key string) {
	var k int64
	c, _ := cc.hm.GetOrInsert(key, &k)
	i := c.(*int64)
	atomic.AddInt64(i, 1)
}

func (cc *LockfreeMapCC) Decrease(key string) {
	var k int64
	c, _ := cc.hm.GetOrInsert(key, &k)
	i := c.(*int64)
	if atomic.LoadInt64(i) < 0 {
		atomic.StoreInt64(i, 0)
	} else {
		atomic.AddInt64(i, -1)
	}
}

func (cc *LockfreeMapCC) Count(key string) int64 {
	c, ok := cc.hm.Get(key)
	if ok {
		i := c.(*int64)
		return atomic.LoadInt64(i)
	}
	return 0
}
