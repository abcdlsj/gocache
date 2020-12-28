package bloomfilter

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)

type BloomFilter struct {
	m uint // number of bits in filter
	k uint // no. of hash functions to use
	b *bitset.BitSet
}

func (bf *BloomFilter) Add(key string) {
	data := []byte(key)

	for i := uint(0); i < bf.k; i++ {
		v, _ := murmur3.Sum128WithSeed(data, uint32(i))
		idx := v % uint64(bf.m)
		bf.b.Set(uint(idx))
	}
}

func (bf *BloomFilter) Contains(key string) bool {
	data := []byte(key)

	for i := uint(0); i < bf.k; i++ {
		v, _ := murmur3.Sum128WithSeed(data, uint32(i))
		idx := v % uint64(bf.m)

		if !bf.b.Test(uint(idx)) {
			return false
		}
	}

	return true
}

func calcM(n uint, p float64) uint {
	val := -((float64(n) * (math.Log(p))) / float64(math.Pow(math.Log(2), 2)))
	return uint(math.Ceil(val))
}

func calcK(p float64) uint {
	return uint(math.Ceil(-1 * math.Log2(p)))
}

func InitBloomFilter(n uint, p float64) *BloomFilter {
	fmt.Printf("InitBloomFilter with %d elements, %2f error rate\n", n, p)
	m := calcM(n, p)
	k := calcK(p)

	return &BloomFilter{m, k, bitset.New(m)}
}
