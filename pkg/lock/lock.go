package lock

import (
	"sync"
)

func NewLockImpl() Lock {
	return &lockImpl{
		locks: make(map[string]*sync.Mutex),
	}
}

func (l *lockImpl) Lock(key string) {
	l.getLockByKey(key).Lock()
}

func (l *lockImpl) Unlock(key string) {
	l.getLockByKey(key).Unlock()
}

func (l *lockImpl) getLockByKey(key string) *sync.Mutex {
	l.mapLock.Lock()
	defer l.mapLock.Unlock()

	lock, found := l.locks[key]
	if found {
		return lock
	}

	lock = &sync.Mutex{}
	l.locks[key] = lock

	return lock
}
