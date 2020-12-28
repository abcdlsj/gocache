package bloomfilter

import (
	"testing"

	"github.com/tj/assert"
)

func TestAddContains(t *testing.T) {
	a := "aaaa"
	b := "bbbb"
	c := "cccc"
	d := "dddd"
	bf := InitBloomFilter(100, 0.0001)
	bf.Add(a)
	bf.Add(b)
	bf.Add(c)

	assert.Equal(t, true, bf.Contains(a))
	assert.Equal(t, true, bf.Contains(b))
	assert.Equal(t, true, bf.Contains(c))
	assert.Equal(t, false, bf.Contains(d))
}
