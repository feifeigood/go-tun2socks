package core

import (
	"sync"

	"github.com/Dreamacro/clash/common/pool"
)

var pl *sync.Pool

const BufSize = 2 * 1024

func SetBufferPool(p *sync.Pool) {
	pl = p
}

func NewBytes(size int) []byte {
	if size <= BufSize {
		return pool.Get(BufSize)
	} else {
		return pool.Get(size)
	}
}

func FreeBytes(b []byte) {
	pool.Put(b)
}

// func init() {
// 	SetBufferPool(&sync.Pool{
// 		New: func() interface{} {
// 			return make([]byte, BufSize)
// 		},
// 	})
// }
