package store

import (
	"fmt"
	"krolus/data"
	"krolus/feed"
	"krolus/models"
	"krolus/treex"

	"github.com/mitchellh/mapstructure"
	"github.com/mmcdole/gofeed"
	"github.com/wailsapp/wails"
)

//FeedStore ...
type FeedStore struct {
	runtime        *wails.Runtime
	logger         *wails.CustomLogger
	manager        *data.Manager
	treeState      *treex.State
	obs            feed.Observable
	selectedLeaves []string
}

// NewFeedStore ...
func NewFeedStore(manager *data.Manager, treeState *treex.State, obs feed.Observable) IFeedStore {
	return &FeedStore{
		manager:   manager,
		treeState: treeState,
		obs:       obs,
	}
}

// WailsInit is called when the component is being initialised
func (f *FeedStore) WailsInit(runtime *wails.Runtime) error {
	f.runtime = runtime
	f.logger = runtime.Log.New("FeedStore")
	go f.listenSubNews()
	return nil
}

// listenSubNews ...
func (f *FeedStore) listenSubNews() {

	for info := range f.obs {
		f.logger.Warnf("Receiving from observable: %v", info)
		//sub is in selectedLeaves ?
		// TODO: encapsulate it
		for sub := range info {
			found := false
			for _, leafID := range f.selectedLeaves {
				if sub.ID == leafID {
					found = true
					f.runtime.Events.Emit("feed.update")
					break
				}
			}
			if found {
				break
			}
		}
	}
}

// LoadMoreItems ...
func (f *FeedStore) LoadMoreItems(request map[string]interface{}) (models.PaginatedItemCollection, error) {

	if request == nil {
		return models.PaginatedItemCollection{}, fmt.Errorf("error loading more items. Request is empty")
	}

	req := &models.PaginatedRequest{}
	if err := mapstructure.Decode(request, req); err != nil {
		return models.PaginatedItemCollection{}, err
	}

	//is directory
	if req.NodeID != "" {
		node := f.treeState.MapNodes.Get(req.NodeID)
		if node == nil {
			return models.PaginatedItemCollection{}, fmt.Errorf("node: %s not found", req.NodeID)
		}
		req.LeafIDs = node.DescendentLeaves()
	}

	f.selectedLeaves = req.LeafIDs

	return f.manager.Item.AllPaginated(*req)
}

// LoadSub Loads Subscription from xurl
func (f *FeedStore) LoadSub(XURL string) (models.SubscriptionModel, error) {

	_feed, err := gofeed.NewParser().ParseURL(XURL)
	if err != nil {
		return models.SubscriptionModel{}, err
	}

	return models.SubscriptionModel{
		Title:       _feed.Title,
		Description: _feed.Description,
		URL:         _feed.Link,
	}, nil

}

// GetSub returns a subscription
func (f *FeedStore) GetSub(ID string) (models.SubscriptionModel, error) {
	sub, err := f.manager.Subscription.Get(ID)
	return *sub, err
}
