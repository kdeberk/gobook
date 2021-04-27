package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f      Func
	mtx    sync.Mutex // guards cache
	cache  map[string]result
	cache4 map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result), cache4: make(map[string]*entry)}
}

func (memo *Memo) Get1(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func (memo *Memo) Get2(key string) (interface{}, error) {
	memo.mtx.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mtx.Unlock()
	return res.value, res.err
}

func (memo *Memo) Get3(key string) (interface{}, error) {
	memo.mtx.Lock()
	res, ok := memo.cache[key]
	memo.mtx.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical section, several goroutines
		// may race to compute f(key) and update the map.
		memo.mtx.Lock()
		memo.cache[key] = res
		memo.mtx.Unlock()
	}
	return res.value, res.err
}

func (memo *Memo) Get4(key string) (interface{}, error) {
	memo.mtx.Lock()
	e := memo.cache4[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache4[key] = e
		memo.mtx.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // broadcast ready condition.
	} else {
		// This is a repeat request for this key.
		memo.mtx.Unlock()
		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
