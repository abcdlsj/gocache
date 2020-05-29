package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           // Hash 算法
	replicas int            // 虚拟节点倍数
	keys     []int          // 哈希环
	HashMap  map[int]string // 虚拟节点与真实节点的映射，key 是虚拟节点的哈希值，value 是真实节点的名称
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		HashMap:  make(map[int]string),
	}

	// default crc32.ChecksumIEEE
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

// 传入真实节点名称，然后对应的添加 replicas 个虚拟节点，然后虚拟节点值映射向真实节点名称
func (m *Map) Add(keysName ...string) {
	for _, key := range keysName {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.HashMap[hash] = key
		}
	}

	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	// 计算节点 hash 值
	hash := int(m.hash([]byte(key)))

	// 顺时针寻找
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	// 环状，通过取余数实现
	return m.HashMap[m.keys[idx%len(m.keys)]]
}
