package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// 对于相同的 key，fn 只会调用一次
func (g *Group) DO(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()

	// 延迟初始化
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()         // 如果请求正在进行，则等待
		return c.val, c.err // 请求结束，返回结果
	}
	c := new(call)
	c.wg.Add(1)  // 发起请求前加锁，锁 + 1
	g.m[key] = c // 添加到 g.m，表面 key 已经有请求在处理了
	g.mu.Unlock()

	c.val, c.err = fn() // 发起请求
	c.wg.Done()         // 请求结束，锁 - 1

	// 并发控制
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
