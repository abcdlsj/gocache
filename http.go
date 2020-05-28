package gocache

import (
	"fmt"
	"log"
)

const DefaultBasePath = "/_gocache/"

type HTTPPool struct {
	self     string
	basepath string
}

// Log info with server name
func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basepath: DefaultBasePath,
	}
}
