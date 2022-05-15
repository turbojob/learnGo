package main

import "sync"

var lock sync.Mutex

func a() {
	lock.Lock()
	lock.Unlock()
	lock.TryLock()
}
func readWriteLock() {
	var readwritelock = sync.RWMutex{}
	readwritelock.RLock()
	readwritelock.RUnlock()
	readwritelock.Lock()

}
