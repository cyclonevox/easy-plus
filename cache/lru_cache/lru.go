package cache

import (
	"container/list"
	"math"
	"sync"

	"github.com/lilith44/easy/cache"
)

const defaultLRUCap = math.MaxInt16

type lruCache[K, V comparable] struct {
	cap   int
	lst   *list.List
	cache sync.Map
	mutex sync.Mutex
}

type node[K, V comparable] struct {
	key   K
	value V
}

func NewLRUCache[K, V comparable](cap ...int) cache.Cache[K, V] {
	c := defaultLRUCap
	if len(cap) != 0 {
		c = cap[0]
	}

	if c <= 0 {
		panic("a positive cap is required")
	}

	return &lruCache[K, V]{
		cap: c,
		lst: list.New(),
	}
}

func (lc *lruCache[K, V]) Set(key K, value V) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	// key存在，移到最近的节点，更新value值
	if element, ok := lc.cache.Load(key); ok {
		e := element.(*list.Element)

		e.Value.(*node[K, V]).value = value
		lc.lst.MoveToFront(e)

		lc.cache.Store(key, e)

		return
	}

	// 已达到最大数量，移除最远节点
	if lc.lst.Len() == lc.cap {
		back := lc.lst.Back()
		lc.lst.Remove(back)

		lc.cache.Delete(back.Value.(*node[K, V]).key)
	}

	// 添加最近的节点
	lc.cache.Store(key, lc.lst.PushFront(&node[K, V]{key: key, value: value}))

	return
}

func (lc *lruCache[K, V]) Get(key K) (v V, exist bool) {
	val, ok := lc.cache.Load(key)
	if !ok {
		return
	}

	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	value := val.(*list.Element)
	lc.lst.MoveToFront(value)

	return value.Value.(*node[K, V]).value, true
}

func (lc *lruCache[K, V]) Del(key K) {
	val, ok := lc.cache.Load(key)
	if !ok {
		return
	}

	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	lc.lst.Remove(val.(*list.Element))
	lc.cache.Delete(key)
}

func (lc *lruCache[K, V]) Flush() {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	lc.lst = list.New()
	lc.cache.Range(func(key, value any) bool {
		lc.cache.Delete(key)

		return true
	})
}
