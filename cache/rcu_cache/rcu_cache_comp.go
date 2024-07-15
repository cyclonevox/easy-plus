package rcu_cache

import (
	"sync"
)

// 本文件仅仅用几种不同的实现来实现接口，作为和rcu_cache之间的对比 并非是用于日常代码使用

type sync_cache[K comparable, V comparable] struct {
	cache sync.Map
}

func (c *sync_cache[K, V]) Get(key K) (V, bool) {
	value, ok := c.cache.Load(key)

	return value.(V), ok
}

func (c *sync_cache[K, V]) Set(key K, value V) {
	c.cache.Store(key, value)
}

func (c *sync_cache[K, V]) Del(key K) {
}

func (c *sync_cache[K, V]) Flush() {
}

type rwm_cache[K comparable, V comparable] struct {
	lock  sync.RWMutex
	cache map[K]V
}

func (c *rwm_cache[K, V]) Get(key K) (V, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cache[key], false
}

func (c *rwm_cache[K, V]) Set(key K, value V) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache[key] = value
}

func (c *rwm_cache[K, V]) Del(key K) {
}

func (c *rwm_cache[K, V]) Flush() {
}

type nlcache[K comparable, V comparable] struct {
	cache map[K]V
}

func (c *nlcache[K, V]) Get(key K) (V, bool) {
	v, ok := c.cache[key]
	return v, ok
}

func (c *nlcache[K, V]) Set(key K, value V) {
	c.cache[key] = value
}

func (c *nlcache[K, V]) Del(key K) {
}

func (c *nlcache[K, V]) Flush() {
}
