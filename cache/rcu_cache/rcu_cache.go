package rcu_cache

import (
	"sync"
	"sync/atomic"

	cacheinterface "github.com/lilith44/easy/cache"
)

type cache[K comparable, V any] struct {
	// todo 暂时以golang map作为存储实现，sonic中使用的是线性探测和rehash来实现的开放寻址哈希
	data map[K]V
}

func newCache[K comparable, V any]() *cache[K, V] {
	return &cache[K, V]{make(map[K]V)}
}

// 深拷贝原副本内容
func (c *cache[K, V]) copy() (fork *cache[K, V]) {
	fork = &cache[K, V]{make(map[K]V)}

	for k, v := range c.data {
		fork.data[k] = v
	}

	return
}

func (c *cache[K, V]) get(key K) (value V, ok bool) {
	value, ok = c.data[key]

	return
}

type rcuCache[K comparable, V comparable] struct {
	// 用于更新时保证操作数据不冲突
	lock sync.Mutex
	// 存储 *cache
	store atomic.Value
}

func New[K comparable, V comparable]() cacheinterface.Cache[K, V] {
	c := &rcuCache[K, V]{
		lock:  sync.Mutex{},
		store: atomic.Value{},
	}
	c.store.Store(newCache[K, V]())

	return c
}

func (c *rcuCache[K, V]) Get(key K) (value V, ok bool) {
	value, ok = c.store.Load().(*cache[K, V]).get(key)

	return
}

func (c *rcuCache[K, V]) Set(key K, value V) {
	c.lock.Lock()
	defer c.lock.Unlock()

	store := c.store.Load().(*cache[K, V])
	v, ok := store.get(key)
	if ok && v == value {
		return
	}

	fork := store.copy()
	fork.data[key] = value

	c.store.Store(fork)
}

func (c *rcuCache[K, V]) Del(key K) {
	c.lock.Lock()
	defer c.lock.Unlock()

	store := c.store.Load().(*cache[K, V])
	_, ok := store.get(key)
	if ok {
		fork := store.copy()
		delete(fork.data, key)
		c.store.Store(fork)
	}

	return
}

func (c *rcuCache[K, V]) Flush() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.store.Store(newCache[K, V]())

	return
}

type KVPack[K comparable, V comparable] struct {
	Key   K
	Value V
}

func (c *rcuCache[K, V]) SetBatch(batch []KVPack[K, V]) {
	c.lock.Lock()
	defer c.lock.Unlock()

	store := c.store.Load().(*cache[K, V])
	var fork *cache[K, V]

	for i := range batch {
		v, ok := store.get(batch[i].Key)
		if ok && v == batch[i].Value {
			continue
		}

		if fork == nil {
			fork = store.copy()
		}

		fork.data[batch[i].Key] = batch[i].Value
	}

	c.store.Store(fork)
}
