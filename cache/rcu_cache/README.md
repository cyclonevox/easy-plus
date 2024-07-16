# 该包源于bytedance/sonic文档中的描述所做的实现
```
但是对于准静态（读远多于写），元素较少（通常不足几十个）的场景，sync.map性能
并不理想，所以使用开放寻址哈希和 RCU 技术重新实现了一个高性能且并发安全的缓存。
```

### rcu 介绍相关:  
https://martong.github.io/high-level-cpp-rcu_informatics_2017.pdf  
https://zhuanlan.zhihu.com/p/582271718

### 目前的实现对sonic中的做法做了简化  
用atomic.Value 代替了atomic.Pointer  
用golang map 代替了开放寻址哈希

### RcuMap 对比 sync.Map，读写锁Map, 非并发安全map的简易benchmark

```
goos: linux
goarch: amd64
pkg: github.com/lilith44/easy/cache/rcu_cache
cpu: 11th Gen Intel(R) Core(TM) i5-11300H @ 3.10GHz
BenchmarkRCUCache_Get
BenchmarkRCUCache_Get-8     	399712183	         2.518 ns/op
BenchmarkSyncMap_Get
BenchmarkSyncMap_Get-8      	155390730	         8.351 ns/op
BenchmarkRWMLockMap_Get
BenchmarkRWMLockMap_Get-8   	28262229	        41.56 ns/op
BenchmarkNoLockMap_Get
BenchmarkNoLockMap_Get-8    	536146963	         2.314 ns/op
PASS
```
#### 结论：   
在几乎纯读的场景下，当前实现的并发安全的RcuMap几乎接近非并发安全map  
而略微领先于sync.Map， 远远强于读写锁实现


