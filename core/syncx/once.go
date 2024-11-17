package syncx

import "sync"

// Once returns a func that guarantees fn can only called once.
// @tips: 这个方法就是对 fn 的一个封装，返回了一个新的函数，保证 fn 只会被调用一次
func Once(fn func()) func() {
	once := new(sync.Once)
	return func() {
		once.Do(fn)
	}
}
