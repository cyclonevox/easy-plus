package rcu_cache

import (
	cache2 "github.com/lilith44/easy/cache"
	"strconv"
	"sync"
	"testing"
)

func TestRCUCache(t *testing.T) {
	c := New[string, string]()
	t.Run("should get value from cache", func(t *testing.T) {
		c.Set("hello", "world")
		v, ok := c.Get("hello")
		if ok != true {
			t.Errorf("ok should be true")
		}
		if v != "world" {
			t.Errorf("value should be 'world'")
		}
	})

	t.Run("should not get value from cache", func(t *testing.T) {
		_, ok := c.Get("hello111")
		if ok != false {
			t.Errorf("ok should be false")
		}
	})

	t.Run("should get different value from cache", func(t *testing.T) {
		c.Set("hello", "world11111")
		v, ok := c.Get("hello")
		if ok != true {
			t.Errorf("ok should be true")
		}
		if v != "world11111" {
			t.Errorf("value should be 'world11111'")
		}
	})

	t.Run("batch set in cache", func(t *testing.T) {
		c.(*rcuCache[string, string]).SetBatch([]KVPack[string, string]{
			{"nihao", "shijie"},
			{"nihao1", "shijie1"},
			{"nihao2", "shijie2"},
		})
		v, ok := c.Get("nihao")
		if ok != true {
			t.Errorf("ok should be true")
		}
		if v != "shijie" {
			t.Errorf("value should be 'shijie'")
		}
		v, ok = c.Get("nihao1")
		if ok != true {
			t.Errorf("ok should be true")
		}
		if v != "shijie1" {
			t.Errorf("value should be 'shijie1'")
		}
		v, ok = c.Get("nihao2")
		if ok != true {
			t.Errorf("ok should be true")
		}
		if v != "shijie2" {
			t.Errorf("value should be 'shijie2'")
		}
	})

	t.Run("delete value from cache", func(t *testing.T) {
		c.Set("delete", "test")
		c.Del("delete")
		_, v := c.Get("delete")
		if v != false {
			t.Errorf("get should be false")
		}
	})

	t.Run("flush cache", func(t *testing.T) {
		c.Set("flush", "test")
		c.Set("flush1", "test")
		c.Flush()

		_, v := c.Get("flush")
		if v != false {
			t.Errorf("get should be false")
		}

		_, v = c.Get("flush1")
		if v != false {
			t.Errorf("get should be false")
		}
	})
}

// RCUCache 针对的是极少情况下会更新的缓存的场景。基本可以认为是完全静态缓存。
// 性能测试以及对sync.map的对比都不会包含对写入的检测和大容量需求的缓存进行检测

func BenchmarkRCUCache_Get(b *testing.B) {
	rcu := New[string, int]()

	for i := 0; i < 50; i++ {
		rcu.Set(strconv.Itoa(i), i)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rcu.Get(strconv.Itoa(25))
		}
	})
}

func BenchmarkSyncMap_Get(b *testing.B) {
	var syncmap cache2.Cache[string, int] = &sync_cache[string, int]{cache: sync.Map{}}
	for i := 0; i < 50; i++ {
		syncmap.Set(strconv.Itoa(i), i)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			syncmap.Get(strconv.Itoa(25))
		}
	})
}

func BenchmarkRWMLockMap_Get(b *testing.B) {
	var rwm cache2.Cache[string, int] = &rwm_cache[string, int]{lock: sync.RWMutex{}, cache: make(map[string]int)}

	for i := 0; i < 50; i++ {
		rwm.Set(strconv.Itoa(i), i)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rwm.Get(strconv.Itoa(25))
		}
	})
}

func BenchmarkNoLockMap_Get(b *testing.B) {
	var c = nlcache[string, int]{cache: make(map[string]int)}

	for i := 0; i < 50; i++ {
		c.Set(strconv.Itoa(i), i)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get(strconv.Itoa(25))
		}
	})
}
