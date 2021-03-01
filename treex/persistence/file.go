package persistence

import (
	"io/ioutil"
	"krolus/treex/models"
	"os"
	"sync"

	"github.com/vmihailenco/msgpack"
)

// File ...
type File struct {
	sync.RWMutex
	path string
}

// NewFile ...
func NewFile(path string) (Persister, error) {

	file := &File{
		path: path,
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return file, nil
	}

	return file, nil
}

// Save saves data into file
func (f *File) Save(root models.Node) error {
	f.Lock()
	defer f.Unlock()

	b, err := msgpack.Marshal(root)
	if err != nil {
		return err
	}

	// write to file
	return ioutil.WriteFile(f.path, b, 0600)
}

// Load loads file to memory
func (f *File) Load(root *models.Node) error {
	f.RLock()
	defer f.RUnlock()

	n, err := ioutil.ReadFile(f.path)
	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(n, root)
	if err != nil {
		return err
	}

	return nil
}
