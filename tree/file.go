package tree

import (
	"sync"
)

type File struct {
	Path string
	MimeType string
	Cache []byte
	sync.RWMutex
}
