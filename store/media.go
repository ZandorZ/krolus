package store

import (
	"krolus/data"
	"krolus/feed/providers"
	"krolus/media"

	"github.com/wailsapp/wails"
)

// MediaStore ...
type MediaStore struct {
	manager *data.Manager
	dl      *media.Downloader
}

// NewMediaStore ...
func NewMediaStore(manager *data.Manager, downloader *media.Downloader) *MediaStore {
	return &MediaStore{
		manager: manager,
		dl:      downloader,
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

	//TODO: decouple proxy
	providers.NewProxy().Download(item)

	return m.dl.Download(item)
}
