package rcu_cache

import (
	"sync"
)

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
	cache map[string]int
}

func (c *nlcache[K, V]) Get(key string) int {
	return c.cache[key]
}

func (c *nlcache[K, V]) Set(key string, value int) {
	c.cache[key] = value
}

func (c *nlcache[K, V]) Del(key K) {
}

func (c *nlcache[K, V]) Flush() {
}
