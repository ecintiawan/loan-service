package lock

import (
	"sync"
)

type (
	Lock interface {
		Lock(key string)
		Unlock(key string)
	}

	lockImpl struct {
		locks   map[string]*sync.Mutex
		mapLock sync.Mutex
	}
)
