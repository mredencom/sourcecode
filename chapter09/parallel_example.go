package main

import "sync"

// TODO 包中的 Parallel 方法用于表示当前测试只会与其他带有 Parallel 方法的测试`并行`进行测试
var (
	data   = make(map[string]string)
	locker sync.Mutex
)

// 写map
func WriteToMap(k, v string) {
	locker.Lock()
	defer locker.Unlock()
	data[k] = v
}

// 读map
func ReaderToMap(k string) string {
	locker.Lock()
	defer locker.Unlock()
	return data[k]
}
