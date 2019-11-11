package main

import (
	"sync"
)

var globalConcurrency map[string]int
var m sync.Mutex

func main() {
	go func() {
		m.Lock()
		globalConcurrency["a"] = 10
		m.Unlock()
	}()

	go func() {
		m.Lock()
		globalConcurrency["b"] = 10
		m.Unlock()
	}()

	go func() {
		m.Lock()
		globalConcurrency["b"] = globalConcurrency["b"] + 1
		m.Unlock()
	}()

}
