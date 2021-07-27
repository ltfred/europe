package sync

import (
	"sync"
)

type ShareCalls interface {
	Do(key string, fn func() (interface{}, error)) (interface{}, error)
}

type call struct {
	wg sync.WaitGroup
	val interface{}
	err error
}

type sharedGroup struct {
	calls map[string]*call
	lock  sync.Mutex
}


func NewShareCalls() ShareCalls {
	return &sharedGroup{
		calls: make(map[string]*call),
	}
}

func (s *sharedGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	s.lock.Lock()
	d, ok := s.calls[key]
	if ok {
		s.lock.Unlock()
		d.wg.Wait()
		return d.val, d.err
	}

	c := new(call)
	c.wg.Add(1)
	s.calls[key] = c
	s.lock.Lock()


	c.val, c.err = fn()

	defer func() {
		s.lock.Unlock()
		delete(s.calls, key)
		s.lock.Unlock()
		c.wg.Done()
	}()

	return c.val, c.err
}