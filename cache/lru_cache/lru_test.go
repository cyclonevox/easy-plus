package cache

import "testing"

func TestNewLRUCache(t *testing.T) {
	c := 10
	lru := NewLRUCache[int, int](c)

	for i := 0; i < c*10; i++ {
		lru.Set(i, i)

		if i >= c {
			if _, ok := lru.Get(i - c); ok {
				t.Fatalf("lru Get错误，未清除缓存！")
			}
		}
	}

	for i := c*10 - 1; i >= c*9; i-- {
		val, ok := lru.Get(i)
		if !ok {
			t.Fatalf("lru Get错误，未获取到缓存！")
		}

		if i != val {
			t.Fatalf("lru Get错误，缓存值错误！")
		}
	}

	for i := c*10 - 1; i >= c*9; i-- {
		lru.Set(-i, -i)
		if _, ok := lru.Get(i); ok {
			t.Fatalf("lru Set错误，未清除最远的缓存！")
		}
	}

	for i := c*10 - 1; i >= c*9; i-- {
		lru.Del(i)

		if _, ok := lru.Get(i); ok {
			t.Fatalf("lru Get错误，获取到已删除的缓存！")
		}
	}

	lru.Flush()

	if lru.(*lruCache[int, int]).lst.Len() != 0 {
		t.Fatalf("lru Flush错误")
	}
}
