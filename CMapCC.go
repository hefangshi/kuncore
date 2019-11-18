package kuncore

import (
	"github/hefangshi/kuncore/cmap"
)

type CMapCC struct {
	cm *cmap.CMap
}

func NewCMapCC() *CMapCC {
	return &CMapCC{
		cm: cmap.New(),
	}
}

func (cc *CMapCC) Increase(key string) {
	cc.cm.Update(key, func (i int64) int64{
		return i+1;
	})
}

func (cc *CMapCC) Decrease(key string) {
	cc.cm.Update(key, func (i int64) int64{
		if i <= 0 {
			return 0
		}
		return i-1;
	})
}

func (cc *CMapCC) Count(key string) int64 {
	return cc.cm.Get(key)
}
