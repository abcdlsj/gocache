package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes int64
	nBytes   int64
	ll       *list.List
	cache    map[string]*list.Element
}
