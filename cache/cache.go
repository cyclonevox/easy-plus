package cache

type Cache[K comparable, V comparable] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Del(key K)
	Flush()
}
