package persistence

import (
	"krolus/treex/models"
	"sync"
)

// Mem ...
type Mem struct {
	sync.RWMutex
}

// NewMem ...
func NewMem() Persister {

	Mem := &Mem{}
	return Mem
}

// Save saves data into Mem
func (f *Mem) Save(root models.Node) error {
	f.Lock()
	defer f.Unlock()

	// write to Mem
	return nil
}

// Load loads Mem to memory
func (f *Mem) Load(root *models.Node) error {
	f.RLock()
	defer f.RUnlock()

	return nil
}
