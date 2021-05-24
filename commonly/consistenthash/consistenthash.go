package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

//一致性哈希实现， 主要用于分布式存储，通过哈希key选择响应的节点。哈希环为了增加节点时候，减少缓存失效的现象。

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           //哈希算法
	replicas int            //虚拟节点倍数
	keys     []int          //数组模拟节点环
	hashMap  map[int]string //虚拟节点对应真实节点
}

func New(replicas int, hash Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     hash,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	index := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[index%len(m.keys)]]
}
