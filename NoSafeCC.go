package kuncore

import "fmt"

type NoSafeCC struct {
	m map[string]int64
}

func NewNoSafeCC() NoSafeCC {
	fmt.Println("Don't use in production")
	return NoSafeCC{
		m: make(map[string]int64),
	}
}

func (cc *NoSafeCC) Increase(key string) {
	cc.m[key]++
}

func (cc *NoSafeCC) Decrease(key string) {
	cc.m[key]--
}

func (cc *NoSafeCC) Count(key string) int64 {
	return cc.m[key]
}
