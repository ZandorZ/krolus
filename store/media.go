package store

import (
	"krolus/data"
	"krolus/media"

	"github.com/wailsapp/wails"
)

// MediaStore ...
type MediaStore struct {
	manager *data.Manager
	dl      *media.Downloader
}

// NewMediaStore ...
func NewMediaStore(manager *data.Manager) *MediaStore {
	return &MediaStore{
		manager: manager,
		dl:      new(media.Downloader),
	}
}

// WailsInit is called when the component is being initialised
func (m *MediaStore) WailsInit(runtime *wails.Runtime) error {
	return nil
}

// Download ...
func (m *MediaStore) Download(ID string) (string, error) {
	// TODO: check cache
	item, err := m.manager.Item.Get(ID)
	if err != nil {
		return "", err
	}
	return m.dl.Download(item.Link)
}
